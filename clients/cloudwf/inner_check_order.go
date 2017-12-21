package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type InnerCheckOrderRequest struct {
	requests.RpcRequest
	Data string `position:"Query" name:"Data"`
}

func (req *InnerCheckOrderRequest) Invoke(client *sdk.Client) (resp *InnerCheckOrderResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "InnerCheckOrder", "", "")
	resp = &InnerCheckOrderResponse{}
	err = client.DoAction(req, resp)
	return
}

type InnerCheckOrderResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Code      string
	Data      string
}
