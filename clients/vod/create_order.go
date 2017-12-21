package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateOrderRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (r CreateOrderRequest) Invoke(client *sdk.Client) (response *CreateOrderResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateOrderRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "CreateOrder", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateOrderResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateOrderResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateOrderResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
}
