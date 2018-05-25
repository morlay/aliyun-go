package market_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InnerVerifyOrderRequest struct {
	requests.RpcRequest
	Data string `position:"Query" name:"Data"`
}

func (req *InnerVerifyOrderRequest) Invoke(client *sdk.Client) (resp *InnerVerifyOrderResponse, err error) {
	req.InitWithApiInfo("Market-Inner", "2016-08-01", "InnerVerifyOrder", "yunmarket", "")
	resp = &InnerVerifyOrderResponse{}
	err = client.DoAction(req, resp)
	return
}

type InnerVerifyOrderResponse struct {
	responses.BaseResponse
	Success   bool
	Data      common.String
	RequestId common.String
	Code      common.String
	Message   common.String
}
