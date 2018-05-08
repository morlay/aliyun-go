package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type PayOrderCallbackRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Data            string `position:"Query" name:"Data"`
}

func (req *PayOrderCallbackRequest) Invoke(client *sdk.Client) (resp *PayOrderCallbackResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "PayOrderCallback", "", "")
	resp = &PayOrderCallbackResponse{}
	err = client.DoAction(req, resp)
	return
}

type PayOrderCallbackResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
}
