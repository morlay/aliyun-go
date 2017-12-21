package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DisableScalingGroupRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DisableScalingGroupRequest) Invoke(client *sdk.Client) (response *DisableScalingGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DisableScalingGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ess", "2014-08-28", "DisableScalingGroup", "ess", "")

	resp := struct {
		*responses.BaseResponse
		DisableScalingGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DisableScalingGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type DisableScalingGroupResponse struct {
	RequestId string
}
