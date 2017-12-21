package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyScalingGroupRequest struct {
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupName             string `position:"Query" name:"ScalingGroupName"`
	ScalingGroupId               string `position:"Query" name:"ScalingGroupId"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	ActiveScalingConfigurationId string `position:"Query" name:"ActiveScalingConfigurationId"`
	MinSize                      int    `position:"Query" name:"MinSize"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	MaxSize                      int    `position:"Query" name:"MaxSize"`
	DefaultCooldown              int    `position:"Query" name:"DefaultCooldown"`
	RemovalPolicy1               string `position:"Query" name:"RemovalPolicy.1"`
	RemovalPolicy2               string `position:"Query" name:"RemovalPolicy.2"`
}

func (r ModifyScalingGroupRequest) Invoke(client *sdk.Client) (response *ModifyScalingGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyScalingGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ess", "2014-08-28", "ModifyScalingGroup", "ess", "")

	resp := struct {
		*responses.BaseResponse
		ModifyScalingGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyScalingGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyScalingGroupResponse struct {
	RequestId string
}
