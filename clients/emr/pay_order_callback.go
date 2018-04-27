package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	Success   bool
	Code      string
	Message   string
}
