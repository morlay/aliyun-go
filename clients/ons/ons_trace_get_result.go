package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsTraceGetResultRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	QueryId      string `position:"Query" name:"QueryId"`
}

func (req *OnsTraceGetResultRequest) Invoke(client *sdk.Client) (resp *OnsTraceGetResultResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsTraceGetResult", "", "")
	resp = &OnsTraceGetResultResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsTraceGetResultResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
	TraceData OnsTraceGetResultTraceData
}

type OnsTraceGetResultTraceData struct {
	QueryId    string
	UserId     string
	Topic      string
	MsgId      string
	MsgKey     string
	Status     string
	CreateTime int64
	UpdateTime int64
	TraceList  OnsTraceGetResultTraceMapDoList
}

type OnsTraceGetResultTraceMapDo struct {
	PubTime      int64
	Topic        string
	PubGroupName string
	MsgId        string
	Tag          string
	MsgKey       string
	BornHost     string
	CostTime     int
	Status       string
	SubList      OnsTraceGetResultSubMapDoList
}

type OnsTraceGetResultSubMapDo struct {
	SubGroupName string
	SuccessCount int
	FailCount    int
	ClientList   OnsTraceGetResultSubClientInfoDoList
}

type OnsTraceGetResultSubClientInfoDo struct {
	SubGroupName   string
	SubTime        int64
	ClientHost     string
	ReconsumeTimes int
	CostTime       int
	Status         string
}

type OnsTraceGetResultTraceMapDoList []OnsTraceGetResultTraceMapDo

func (list *OnsTraceGetResultTraceMapDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTraceGetResultTraceMapDo)
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

type OnsTraceGetResultSubMapDoList []OnsTraceGetResultSubMapDo

func (list *OnsTraceGetResultSubMapDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTraceGetResultSubMapDo)
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

type OnsTraceGetResultSubClientInfoDoList []OnsTraceGetResultSubClientInfoDo

func (list *OnsTraceGetResultSubClientInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTraceGetResultSubClientInfoDo)
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
