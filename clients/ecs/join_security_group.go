package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type JoinSecurityGroupRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *JoinSecurityGroupRequest) Invoke(client *sdk.Client) (resp *JoinSecurityGroupResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "JoinSecurityGroup", "ecs", "")
	resp = &JoinSecurityGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type JoinSecurityGroupResponse struct {
	responses.BaseResponse
	RequestId string
}
