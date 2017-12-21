package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateScalingRuleRequest struct {
	ScalingRuleName      string `position:"Query" name:"ScalingRuleName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AdjustmentValue      int    `position:"Query" name:"AdjustmentValue"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Cooldown             int    `position:"Query" name:"Cooldown"`
	AdjustmentType       string `position:"Query" name:"AdjustmentType"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CreateScalingRuleRequest) Invoke(client *sdk.Client) (response *CreateScalingRuleResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateScalingRuleRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ess", "2014-08-28", "CreateScalingRule", "ess", "")

	resp := struct {
		*responses.BaseResponse
		CreateScalingRuleResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateScalingRuleResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateScalingRuleResponse struct {
	ScalingRuleId  string
	ScalingRuleAri string
	RequestId      string
}
