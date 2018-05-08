package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteScalingRuleRequest struct {
	requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ScalingRuleId        string `position:"Query" name:"ScalingRuleId"`
}

func (req *DeleteScalingRuleRequest) Invoke(client *sdk.Client) (resp *DeleteScalingRuleResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DeleteScalingRule", "ess", "")
	resp = &DeleteScalingRuleResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteScalingRuleResponse struct {
	responses.BaseResponse
	RequestId common.String
}
