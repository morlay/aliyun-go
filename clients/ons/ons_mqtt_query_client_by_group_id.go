package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMqttQueryClientByGroupIdRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	GroupId      string `position:"Query" name:"GroupId"`
}

func (r OnsMqttQueryClientByGroupIdRequest) Invoke(client *sdk.Client) (response *OnsMqttQueryClientByGroupIdResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsMqttQueryClientByGroupIdRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttQueryClientByGroupId", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsMqttQueryClientByGroupIdResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsMqttQueryClientByGroupIdResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsMqttQueryClientByGroupIdResponse struct {
	RequestId       string
	HelpUrl         string
	MqttClientSetDo OnsMqttQueryClientByGroupIdMqttClientSetDo
}

type OnsMqttQueryClientByGroupIdMqttClientSetDo struct {
	OnlineCount  int64
	PersistCount int64
}
