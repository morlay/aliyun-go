package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMqttQueryMsgTransTrendRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Qos          int    `position:"Query" name:"Qos"`
	TransType    string `position:"Query" name:"TransType"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
	TpsType      string `position:"Query" name:"TpsType"`
	ParentTopic  string `position:"Query" name:"ParentTopic"`
	MsgType      string `position:"Query" name:"MsgType"`
	SubTopic     string `position:"Query" name:"SubTopic"`
}

func (r OnsMqttQueryMsgTransTrendRequest) Invoke(client *sdk.Client) (response *OnsMqttQueryMsgTransTrendResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsMqttQueryMsgTransTrendRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttQueryMsgTransTrend", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsMqttQueryMsgTransTrendResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsMqttQueryMsgTransTrendResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsMqttQueryMsgTransTrendResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsMqttQueryMsgTransTrendData
}

type OnsMqttQueryMsgTransTrendData struct {
	Title   string
	XUnit   string
	YUnit   string
	Records OnsMqttQueryMsgTransTrendStatsDataDoList
}

type OnsMqttQueryMsgTransTrendStatsDataDo struct {
	X int64
	Y float32
}

type OnsMqttQueryMsgTransTrendStatsDataDoList []OnsMqttQueryMsgTransTrendStatsDataDo

func (list *OnsMqttQueryMsgTransTrendStatsDataDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMqttQueryMsgTransTrendStatsDataDo)
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
