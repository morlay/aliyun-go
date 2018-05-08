package aegis

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Code      common.String
	Success   bool
	Message   common.String
	Data      GetEntityListData
}

type GetEntityListData struct {
	List     GetEntityListEntityList
	PageInfo GetEntityListPageInfo
}

type GetEntityListEntity struct {
	Uuid         common.String
	GroupId      common.Long
	Ip           common.String
	InstanceName common.String
	InstanceId   common.String
	Region       common.String
	Os           common.String
	Flag         common.String
	BuyVersion   common.String
	AegisOnline  bool
	AegisVersion common.String
}

type GetEntityListPageInfo struct {
	CurrentPage common.Integer
	PageSize    common.Integer
	TotalCount  common.Integer
	Count       common.Integer
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
