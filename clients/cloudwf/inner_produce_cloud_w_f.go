package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type InnerProduceCloudWFRequest struct {
	Data string `position:"Query" name:"Data"`
}

func (r InnerProduceCloudWFRequest) Invoke(client *sdk.Client) (response *InnerProduceCloudWFResponse, err error) {
	req := struct {
		*requests.RpcRequest
		InnerProduceCloudWFRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "InnerProduceCloudWF", "", "")

	resp := struct {
		*responses.BaseResponse
		InnerProduceCloudWFResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.InnerProduceCloudWFResponse

	err = client.DoAction(&req, &resp)
	return
}

type InnerProduceCloudWFResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
}
