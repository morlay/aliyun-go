package market

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateOrderRequest struct {
	requests.RpcRequest
	OrderType   string `position:"Query" name:"OrderType"`
	Commodity   string `position:"Query" name:"Commodity"`
	OrderSouce  string `position:"Query" name:"OrderSouce"`
	PaymentType string `position:"Query" name:"PaymentType"`
	ClientToken string `position:"Query" name:"ClientToken"`
}

func (req *CreateOrderRequest) Invoke(client *sdk.Client) (resp *CreateOrderResponse, err error) {
	req.InitWithApiInfo("Market", "2015-11-01", "CreateOrder", "", "")
	resp = &CreateOrderResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateOrderResponse struct {
	responses.BaseResponse
	OrderId string
}
