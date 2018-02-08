package aegis

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetEntityListRequest struct {
	requests.RpcRequest
	GroupId     int64  `position:"Query" name:"GroupId"`
	PageSize    int    `position:"Query" name:"PageSize"`
	Remark      string `position:"Query" name:"Remark"`
	EventType   string `position:"Query" name:"EventType"`
	CurrentPage int    `position:"Query" name:"CurrentPage"`
	RegionNo    string `position:"Query" name:"RegionNo"`
}

func (req *GetEntityListRequest) Invoke(client *sdk.Client) (resp *GetEntityListResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "GetEntityList", "vipaegis", "")
	resp = &GetEntityListResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetEntityListResponse struct {
	responses.BaseResponse
	RequestId string
	Code      string
	Success   bool
	Message   string
	Data      GetEntityListData
}

type GetEntityListData struct {
	List     GetEntityListEntityList
	PageInfo GetEntityListPageInfo
}

type GetEntityListEntity struct {
	Uuid         string
	GroupId      int64
	Ip           string
	InstanceName string
	InstanceId   string
	Region       string
	Os           string
	Flag         string
	BuyVersion   string
	AegisOnline  bool
	AegisVersion string
}

type GetEntityListPageInfo struct {
	CurrentPage int
	PageSize    int
	TotalCount  int
	Count       int
}

type GetEntityListEntityList []GetEntityListEntity

func (list *GetEntityListEntityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetEntityListEntity)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
