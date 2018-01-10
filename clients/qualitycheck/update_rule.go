package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateRuleRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *UpdateRuleRequest) Invoke(client *sdk.Client) (resp *UpdateRuleResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UpdateRule", "", "")
	resp = &UpdateRuleResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateRuleResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      string
}
