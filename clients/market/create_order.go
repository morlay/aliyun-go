package market

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateOrderRequest struct {
	OrderType   string `position:"Query" name:"OrderType"`
	Commodity   string `position:"Query" name:"Commodity"`
	OrderSouce  string `position:"Query" name:"OrderSouce"`
	PaymentType string `position:"Query" name:"PaymentType"`
	ClientToken string `position:"Query" name:"ClientToken"`
}

func (r CreateOrderRequest) Invoke(client *sdk.Client) (response *CreateOrderResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateOrderRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Market", "2015-11-01", "CreateOrder", "", "")

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
	OrderId string
}
