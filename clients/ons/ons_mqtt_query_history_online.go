package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMqttQueryHistoryOnlineRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	GroupId      string `position:"Query" name:"GroupId"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
}

func (req *OnsMqttQueryHistoryOnlineRequest) Invoke(client *sdk.Client) (resp *OnsMqttQueryHistoryOnlineResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttQueryHistoryOnline", "", "")
	resp = &OnsMqttQueryHistoryOnlineResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsMqttQueryHistoryOnlineResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
	Data      OnsMqttQueryHistoryOnlineData
}

type OnsMqttQueryHistoryOnlineData struct {
	Title   string
	XUnit   string
	YUnit   string
	Records OnsMqttQueryHistoryOnlineStatsDataDoList
}

type OnsMqttQueryHistoryOnlineStatsDataDo struct {
	X int64
	Y float32
}

type OnsMqttQueryHistoryOnlineStatsDataDoList []OnsMqttQueryHistoryOnlineStatsDataDo

func (list *OnsMqttQueryHistoryOnlineStatsDataDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMqttQueryHistoryOnlineStatsDataDo)
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
