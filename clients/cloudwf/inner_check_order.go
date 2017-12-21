package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type InnerCheckOrderRequest struct {
	Data string `position:"Query" name:"Data"`
}

func (r InnerCheckOrderRequest) Invoke(client *sdk.Client) (response *InnerCheckOrderResponse, err error) {
	req := struct {
		*requests.RpcRequest
		InnerCheckOrderRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "InnerCheckOrder", "", "")

	resp := struct {
		*responses.BaseResponse
		InnerCheckOrderResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.InnerCheckOrderResponse

	err = client.DoAction(&req, &resp)
	return
}

type InnerCheckOrderResponse struct {
	RequestId string
	Success   bool
	Message   string
	Code      string
	Data      string
}
