package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteSecurityGroupRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteSecurityGroupRequest) Invoke(client *sdk.Client) (resp *DeleteSecurityGroupResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteSecurityGroup", "ecs", "")
	resp = &DeleteSecurityGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteSecurityGroupResponse struct {
	responses.BaseResponse
	RequestId string
}
