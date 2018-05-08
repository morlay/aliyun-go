package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type LeaveSecurityGroupRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *LeaveSecurityGroupRequest) Invoke(client *sdk.Client) (resp *LeaveSecurityGroupResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "LeaveSecurityGroup", "ecs", "")
	resp = &LeaveSecurityGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type LeaveSecurityGroupResponse struct {
	responses.BaseResponse
	RequestId common.String
}
