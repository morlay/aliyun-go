package commondriver

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetOrderIdByCheckBeforePayRequest struct {
	OrderId string `position:"Query" name:"OrderId"`
}

func (r GetOrderIdByCheckBeforePayRequest) Invoke(client *sdk.Client) (response *GetOrderIdByCheckBeforePayResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetOrderIdByCheckBeforePayRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Commondriver", "2015-12-29", "GetOrderIdByCheckBeforePay", "", "")

	resp := struct {
		*responses.BaseResponse
		GetOrderIdByCheckBeforePayResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetOrderIdByCheckBeforePayResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetOrderIdByCheckBeforePayResponse struct {
	RequestId string
	Success   bool
	Data      bool
}
