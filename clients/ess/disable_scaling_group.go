package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DisableScalingGroupRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DisableScalingGroupRequest) Invoke(client *sdk.Client) (resp *DisableScalingGroupResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DisableScalingGroup", "ess", "")
	resp = &DisableScalingGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type DisableScalingGroupResponse struct {
	responses.BaseResponse
	RequestId common.String
}
