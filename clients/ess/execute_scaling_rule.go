package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ExecuteScalingRuleRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ScalingRuleAri       string `position:"Query" name:"ScalingRuleAri"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ExecuteScalingRuleRequest) Invoke(client *sdk.Client) (response *ExecuteScalingRuleResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ExecuteScalingRuleRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ess", "2014-08-28", "ExecuteScalingRule", "ess", "")

	resp := struct {
		*responses.BaseResponse
		ExecuteScalingRuleResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ExecuteScalingRuleResponse

	err = client.DoAction(&req, &resp)
	return
}

type ExecuteScalingRuleResponse struct {
	ScalingActivityId string
	RequestId         string
}
