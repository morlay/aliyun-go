package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyHealthRuleConfigRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RuleName        string `position:"Query" name:"RuleName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RuleId          int64  `position:"Query" name:"RuleId"`
	Status          string `position:"Query" name:"Status"`
}

func (req *ModifyHealthRuleConfigRequest) Invoke(client *sdk.Client) (resp *ModifyHealthRuleConfigResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyHealthRuleConfig", "", "")
	resp = &ModifyHealthRuleConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyHealthRuleConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
