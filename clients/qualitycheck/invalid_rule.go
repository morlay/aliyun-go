package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InvalidRuleRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *InvalidRuleRequest) Invoke(client *sdk.Client) (resp *InvalidRuleResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "InvalidRule", "", "")
	resp = &InvalidRuleResponse{}
	err = client.DoAction(req, resp)
	return
}

type InvalidRuleResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      bool
}
