package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ExecuteScalingRuleRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ScalingRuleAri       string `position:"Query" name:"ScalingRuleAri"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ExecuteScalingRuleRequest) Invoke(client *sdk.Client) (resp *ExecuteScalingRuleResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "ExecuteScalingRule", "ess", "")
	resp = &ExecuteScalingRuleResponse{}
	err = client.DoAction(req, resp)
	return
}

type ExecuteScalingRuleResponse struct {
	responses.BaseResponse
	ScalingActivityId common.String
	RequestId         common.String
}
