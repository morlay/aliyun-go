package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMqttQueryHistoryOnlineRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	GroupId      string `position:"Query" name:"GroupId"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
}

func (r OnsMqttQueryHistoryOnlineRequest) Invoke(client *sdk.Client) (response *OnsMqttQueryHistoryOnlineResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsMqttQueryHistoryOnlineRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttQueryHistoryOnline", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsMqttQueryHistoryOnlineResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsMqttQueryHistoryOnlineResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsMqttQueryHistoryOnlineResponse struct {
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
