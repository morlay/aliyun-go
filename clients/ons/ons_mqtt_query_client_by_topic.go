package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsMqttQueryClientByTopicRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ParentTopic  string `position:"Query" name:"ParentTopic"`
	SubTopic     string `position:"Query" name:"SubTopic"`
}

func (req *OnsMqttQueryClientByTopicRequest) Invoke(client *sdk.Client) (resp *OnsMqttQueryClientByTopicResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttQueryClientByTopic", "", "")
	resp = &OnsMqttQueryClientByTopicResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsMqttQueryClientByTopicResponse struct {
	responses.BaseResponse
	RequestId       common.String
	HelpUrl         common.String
	MqttClientSetDo OnsMqttQueryClientByTopicMqttClientSetDo
}

type OnsMqttQueryClientByTopicMqttClientSetDo struct {
	OnlineCount  common.Long
	PersistCount common.Long
}
