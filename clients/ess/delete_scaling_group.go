package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteScalingGroupRequest struct {
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	ForceDelete          string `position:"Query" name:"ForceDelete"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteScalingGroupRequest) Invoke(client *sdk.Client) (response *DeleteScalingGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteScalingGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ess", "2014-08-28", "DeleteScalingGroup", "ess", "")

	resp := struct {
		*responses.BaseResponse
		DeleteScalingGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteScalingGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteScalingGroupResponse struct {
	RequestId string
}
