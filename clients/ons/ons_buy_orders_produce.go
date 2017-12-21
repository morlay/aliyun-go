package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsBuyOrdersProduceRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Data         string `position:"Query" name:"Data"`
}

func (r OnsBuyOrdersProduceRequest) Invoke(client *sdk.Client) (response *OnsBuyOrdersProduceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsBuyOrdersProduceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsBuyOrdersProduce", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsBuyOrdersProduceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsBuyOrdersProduceResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsBuyOrdersProduceResponse struct {
	Success   bool
	RequestId string
	Code      string
	Message   string
	Data      string
}
