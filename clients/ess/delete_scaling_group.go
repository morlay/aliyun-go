package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteScalingGroupRequest struct {
	requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	ForceDelete          string `position:"Query" name:"ForceDelete"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteScalingGroupRequest) Invoke(client *sdk.Client) (resp *DeleteScalingGroupResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DeleteScalingGroup", "ess", "")
	resp = &DeleteScalingGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteScalingGroupResponse struct {
	responses.BaseResponse
	RequestId common.String
}
