package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type InnerRefundRequest struct {
	requests.RpcRequest
	Data string `position:"Query" name:"Data"`
}

func (req *InnerRefundRequest) Invoke(client *sdk.Client) (resp *InnerRefundResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "InnerRefund", "", "")
	resp = &InnerRefundResponse{}
	err = client.DoAction(req, resp)
	return
}

type InnerRefundResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
}
