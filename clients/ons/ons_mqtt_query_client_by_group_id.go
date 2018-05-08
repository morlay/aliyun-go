package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsMqttQueryClientByGroupIdRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	GroupId      string `position:"Query" name:"GroupId"`
}

func (req *OnsMqttQueryClientByGroupIdRequest) Invoke(client *sdk.Client) (resp *OnsMqttQueryClientByGroupIdResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttQueryClientByGroupId", "", "")
	resp = &OnsMqttQueryClientByGroupIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsMqttQueryClientByGroupIdResponse struct {
	responses.BaseResponse
	RequestId       common.String
	HelpUrl         common.String
	MqttClientSetDo OnsMqttQueryClientByGroupIdMqttClientSetDo
}

type OnsMqttQueryClientByGroupIdMqttClientSetDo struct {
	OnlineCount  common.Long
	PersistCount common.Long
}
