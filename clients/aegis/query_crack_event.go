package aegis

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	req.InitWithApiInfo("aegis", "2016-11-11", "QueryCrackEvent", "", "")
	resp = &QueryCrackEventResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryCrackEventResponse struct {
	responses.BaseResponse
	RequestId string
	Code      string
	Success   bool
	Message   string
	Data      QueryCrackEventData
}

type QueryCrackEventData struct {
	List     QueryCrackEventEntityList
	PageInfo QueryCrackEventPageInfo
}

type QueryCrackEventEntity struct {
	Uuid           string
	AttackTime     string
	AttackType     int
	AttackTypeName string
	BuyVersion     string
	CrackSourceIp  string
	CrackTimes     int
	GroupId        int
	InstanceName   string
	InstanceId     string
	Ip             string
	Region         string
	Status         int
	StatusName     string
	Location       string
	InWhite        int
	UserName       string
}

type QueryCrackEventPageInfo struct {
	CurrentPage int
	PageSize    int
	TotalCount  int
	Count       int
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
