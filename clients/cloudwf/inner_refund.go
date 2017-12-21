package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type InnerRefundRequest struct {
	Data string `position:"Query" name:"Data"`
}

func (r InnerRefundRequest) Invoke(client *sdk.Client) (response *InnerRefundResponse, err error) {
	req := struct {
		*requests.RpcRequest
		InnerRefundRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "InnerRefund", "", "")

	resp := struct {
		*responses.BaseResponse
		InnerRefundResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.InnerRefundResponse

	err = client.DoAction(&req, &resp)
	return
}

type InnerRefundResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
}
