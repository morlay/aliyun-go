package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMqttBuyRefundRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Data         string `position:"Query" name:"Data"`
}

func (r OnsMqttBuyRefundRequest) Invoke(client *sdk.Client) (response *OnsMqttBuyRefundResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsMqttBuyRefundRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttBuyRefund", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsMqttBuyRefundResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsMqttBuyRefundResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsMqttBuyRefundResponse struct {
	Success   bool
	RequestId string
	Code      string
	Message   string
	Data      string
}
