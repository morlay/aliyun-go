package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyScalingGroupRequest struct {
	requests.RpcRequest
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

func (req *ModifyScalingGroupRequest) Invoke(client *sdk.Client) (resp *ModifyScalingGroupResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "ModifyScalingGroup", "ess", "")
	resp = &ModifyScalingGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyScalingGroupResponse struct {
	responses.BaseResponse
	RequestId common.String
}
