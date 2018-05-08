package aegis

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryLoginEventRequest struct {
	requests.RpcRequest
	EndTime     string `position:"Query" name:"EndTime"`
	CurrentPage int    `position:"Query" name:"CurrentPage"`
	StartTime   string `position:"Query" name:"StartTime"`
	Uuid        string `position:"Query" name:"Uuid"`
	Status      int    `position:"Query" name:"Status"`
}

func (req *QueryLoginEventRequest) Invoke(client *sdk.Client) (resp *QueryLoginEventResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "QueryLoginEvent", "vipaegis", "")
	resp = &QueryLoginEventResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryLoginEventResponse struct {
	responses.BaseResponse
	RequestId common.String
	Code      common.String
	Success   bool
	Message   common.String
	Data      QueryLoginEventData
}

type QueryLoginEventData struct {
	List     QueryLoginEventEntityList
	PageInfo QueryLoginEventPageInfo
}

type QueryLoginEventEntity struct {
	Uuid          common.String
	LoginTime     common.String
	LoginType     common.Integer
	LoginTypeName common.String
	BuyVersion    common.String
	LoginSourceIp common.String
	GroupId       common.Integer
	InstanceName  common.String
	InstanceId    common.String
	Ip            common.String
	Region        common.String
	Status        common.Integer
	StatusName    common.String
	Location      common.String
	UserName      common.String
}

type QueryLoginEventPageInfo struct {
	CurrentPage common.Integer
	PageSize    common.Integer
	TotalCount  common.Integer
	Count       common.Integer
}

type QueryLoginEventEntityList []QueryLoginEventEntity

func (list *QueryLoginEventEntityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryLoginEventEntity)
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
