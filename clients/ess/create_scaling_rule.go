package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateScalingRuleRequest struct {
	requests.RpcRequest
	ScalingRuleName      string `position:"Query" name:"ScalingRuleName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AdjustmentValue      int    `position:"Query" name:"AdjustmentValue"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Cooldown             int    `position:"Query" name:"Cooldown"`
	AdjustmentType       string `position:"Query" name:"AdjustmentType"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CreateScalingRuleRequest) Invoke(client *sdk.Client) (resp *CreateScalingRuleResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "CreateScalingRule", "ess", "")
	resp = &CreateScalingRuleResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateScalingRuleResponse struct {
	responses.BaseResponse
	ScalingRuleId  common.String
	ScalingRuleAri common.String
	RequestId      common.String
}
