package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteScalingRuleRequest struct {
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ScalingRuleId        string `position:"Query" name:"ScalingRuleId"`
}

func (r DeleteScalingRuleRequest) Invoke(client *sdk.Client) (response *DeleteScalingRuleResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteScalingRuleRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ess", "2014-08-28", "DeleteScalingRule", "ess", "")

	resp := struct {
		*responses.BaseResponse
		DeleteScalingRuleResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteScalingRuleResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteScalingRuleResponse struct {
	RequestId string
}
