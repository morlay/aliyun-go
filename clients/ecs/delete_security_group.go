package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteSecurityGroupRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteSecurityGroupRequest) Invoke(client *sdk.Client) (response *DeleteSecurityGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteSecurityGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteSecurityGroup", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DeleteSecurityGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteSecurityGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteSecurityGroupResponse struct {
	RequestId string
}
