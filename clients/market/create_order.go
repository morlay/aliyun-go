package market

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateOrderRequest struct {
	requests.RpcRequest
	OrderSouce  string `position:"Query" name:"OrderSouce"`
	Commodity   string `position:"Query" name:"Commodity"`
	ClientToken string `position:"Query" name:"ClientToken"`
	OwnerId     string `position:"Query" name:"OwnerId"`
	PaymentType string `position:"Query" name:"PaymentType"`
	OrderType   string `position:"Query" name:"OrderType"`
}

func (req *CreateOrderRequest) Invoke(client *sdk.Client) (resp *CreateOrderResponse, err error) {
	req.InitWithApiInfo("Market", "2015-11-01", "CreateOrder", "yunmarket", "")
	resp = &CreateOrderResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateOrderResponse struct {
	responses.BaseResponse
	OrderId string
}
