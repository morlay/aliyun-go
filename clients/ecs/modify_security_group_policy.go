package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifySecurityGroupPolicyRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InnerAccessPolicy    string `position:"Query" name:"InnerAccessPolicy"`
}

func (req *ModifySecurityGroupPolicyRequest) Invoke(client *sdk.Client) (resp *ModifySecurityGroupPolicyResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifySecurityGroupPolicy", "ecs", "")
	resp = &ModifySecurityGroupPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifySecurityGroupPolicyResponse struct {
	responses.BaseResponse
	RequestId common.String
}
