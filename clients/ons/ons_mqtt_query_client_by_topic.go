package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMqttQueryClientByTopicRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ParentTopic  string `position:"Query" name:"ParentTopic"`
	SubTopic     string `position:"Query" name:"SubTopic"`
}

func (r OnsMqttQueryClientByTopicRequest) Invoke(client *sdk.Client) (response *OnsMqttQueryClientByTopicResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsMqttQueryClientByTopicRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttQueryClientByTopic", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsMqttQueryClientByTopicResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsMqttQueryClientByTopicResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsMqttQueryClientByTopicResponse struct {
	RequestId       string
	HelpUrl         string
	MqttClientSetDo OnsMqttQueryClientByTopicMqttClientSetDo
}

type OnsMqttQueryClientByTopicMqttClientSetDo struct {
	OnlineCount  int64
	PersistCount int64
}
