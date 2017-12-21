package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyScalingRuleRequest struct {
	requests.RpcRequest
	ScalingRuleName      string `position:"Query" name:"ScalingRuleName"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AdjustmentValue      int    `position:"Query" name:"AdjustmentValue"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Cooldown             int    `position:"Query" name:"Cooldown"`
	AdjustmentType       string `position:"Query" name:"AdjustmentType"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ScalingRuleId        string `position:"Query" name:"ScalingRuleId"`
}

func (req *ModifyScalingRuleRequest) Invoke(client *sdk.Client) (resp *ModifyScalingRuleResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "ModifyScalingRule", "ess", "")
	resp = &ModifyScalingRuleResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyScalingRuleResponse struct {
	responses.BaseResponse
	RequestId string
}
