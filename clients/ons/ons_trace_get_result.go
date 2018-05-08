package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	HelpUrl   common.String
	TraceData OnsTraceGetResultTraceData
}

type OnsTraceGetResultTraceData struct {
	QueryId    common.String
	UserId     common.String
	Topic      common.String
	MsgId      common.String
	MsgKey     common.String
	Status     common.String
	CreateTime common.Long
	UpdateTime common.Long
	TraceList  OnsTraceGetResultTraceMapDoList
}

type OnsTraceGetResultTraceMapDo struct {
	PubTime      common.Long
	Topic        common.String
	PubGroupName common.String
	MsgId        common.String
	Tag          common.String
	MsgKey       common.String
	BornHost     common.String
	CostTime     common.Integer
	Status       common.String
	SubList      OnsTraceGetResultSubMapDoList
}

type OnsTraceGetResultSubMapDo struct {
	SubGroupName common.String
	SuccessCount common.Integer
	FailCount    common.Integer
	ClientList   OnsTraceGetResultSubClientInfoDoList
}

type OnsTraceGetResultSubClientInfoDo struct {
	SubGroupName   common.String
	SubTime        common.Long
	ClientHost     common.String
	ReconsumeTimes common.Integer
	CostTime       common.Integer
	Status         common.String
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
