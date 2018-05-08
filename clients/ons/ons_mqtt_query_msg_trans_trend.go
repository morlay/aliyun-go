package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsMqttQueryMsgTransTrendRequest struct {
	requests.RpcRequest
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

func (req *OnsMqttQueryMsgTransTrendRequest) Invoke(client *sdk.Client) (resp *OnsMqttQueryMsgTransTrendResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttQueryMsgTransTrend", "", "")
	resp = &OnsMqttQueryMsgTransTrendResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsMqttQueryMsgTransTrendResponse struct {
	responses.BaseResponse
	RequestId common.String
	HelpUrl   common.String
	Data      OnsMqttQueryMsgTransTrendData
}

type OnsMqttQueryMsgTransTrendData struct {
	Title   common.String
	XUnit   common.String
	YUnit   common.String
	Records OnsMqttQueryMsgTransTrendStatsDataDoList
}

type OnsMqttQueryMsgTransTrendStatsDataDo struct {
	X common.Long
	Y common.Float
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
