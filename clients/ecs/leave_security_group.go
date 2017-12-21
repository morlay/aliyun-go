package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type LeaveSecurityGroupRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r LeaveSecurityGroupRequest) Invoke(client *sdk.Client) (response *LeaveSecurityGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		LeaveSecurityGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "LeaveSecurityGroup", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		LeaveSecurityGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.LeaveSecurityGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type LeaveSecurityGroupResponse struct {
	RequestId string
}
