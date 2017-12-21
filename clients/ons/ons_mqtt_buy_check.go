package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMqttBuyCheckRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Data         string `position:"Query" name:"Data"`
}

func (r OnsMqttBuyCheckRequest) Invoke(client *sdk.Client) (response *OnsMqttBuyCheckResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsMqttBuyCheckRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttBuyCheck", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsMqttBuyCheckResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsMqttBuyCheckResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsMqttBuyCheckResponse struct {
	Success   bool
	RequestId string
	Code      string
	Message   string
	Data      string
}
