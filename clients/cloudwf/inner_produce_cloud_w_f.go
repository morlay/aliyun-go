package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type InnerProduceCloudWFRequest struct {
	requests.RpcRequest
	Data string `position:"Query" name:"Data"`
}

func (req *InnerProduceCloudWFRequest) Invoke(client *sdk.Client) (resp *InnerProduceCloudWFResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "InnerProduceCloudWF", "", "")
	resp = &InnerProduceCloudWFResponse{}
	err = client.DoAction(req, resp)
	return
}

type InnerProduceCloudWFResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
}
