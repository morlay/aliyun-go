package aegis

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryCrackEventRequest struct {
	requests.RpcRequest
	EndTime     string `position:"Query" name:"EndTime"`
	CurrentPage int    `position:"Query" name:"CurrentPage"`
	StartTime   string `position:"Query" name:"StartTime"`
	Uuid        string `position:"Query" name:"Uuid"`
	Status      int    `position:"Query" name:"Status"`
}

func (req *QueryCrackEventRequest) Invoke(client *sdk.Client) (resp *QueryCrackEventResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "QueryCrackEvent", "vipaegis", "")
	resp = &QueryCrackEventResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryCrackEventResponse struct {
	responses.BaseResponse
	RequestId common.String
	Code      common.String
	Success   bool
	Message   common.String
	Data      QueryCrackEventData
}

type QueryCrackEventData struct {
	List     QueryCrackEventEntityList
	PageInfo QueryCrackEventPageInfo
}

type QueryCrackEventEntity struct {
	Uuid           common.String
	AttackTime     common.String
	AttackType     common.Integer
	AttackTypeName common.String
	BuyVersion     common.String
	CrackSourceIp  common.String
	CrackTimes     common.Integer
	GroupId        common.Integer
	InstanceName   common.String
	InstanceId     common.String
	Ip             common.String
	Region         common.String
	Status         common.Integer
	StatusName     common.String
	Location       common.String
	InWhite        common.Integer
	UserName       common.String
}

type QueryCrackEventPageInfo struct {
	CurrentPage common.Integer
	PageSize    common.Integer
	TotalCount  common.Integer
	Count       common.Integer
}

type QueryCrackEventEntityList []QueryCrackEventEntity

func (list *QueryCrackEventEntityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCrackEventEntity)
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
