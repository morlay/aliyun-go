package commondriver

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetOrderIdByQueryPurchaseRequest struct {
	OrderId string `position:"Query" name:"OrderId"`
}

func (r GetOrderIdByQueryPurchaseRequest) Invoke(client *sdk.Client) (response *GetOrderIdByQueryPurchaseResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetOrderIdByQueryPurchaseRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Commondriver", "2015-12-29", "GetOrderIdByQueryPurchase", "", "")

	resp := struct {
		*responses.BaseResponse
		GetOrderIdByQueryPurchaseResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetOrderIdByQueryPurchaseResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetOrderIdByQueryPurchaseResponse struct {
	RequestId string
	Success   bool
	Data      bool
}
