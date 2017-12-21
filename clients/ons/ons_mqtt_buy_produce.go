package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMqttBuyProduceRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Data         string `position:"Query" name:"Data"`
}

func (r OnsMqttBuyProduceRequest) Invoke(client *sdk.Client) (response *OnsMqttBuyProduceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsMqttBuyProduceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttBuyProduce", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsMqttBuyProduceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsMqttBuyProduceResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsMqttBuyProduceResponse struct {
	Success   bool
	RequestId string
	Code      string
	Message   string
	Data      string
}
