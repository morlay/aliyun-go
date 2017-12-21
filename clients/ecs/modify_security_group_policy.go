package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifySecurityGroupPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InnerAccessPolicy    string `position:"Query" name:"InnerAccessPolicy"`
}

func (r ModifySecurityGroupPolicyRequest) Invoke(client *sdk.Client) (response *ModifySecurityGroupPolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifySecurityGroupPolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifySecurityGroupPolicy", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifySecurityGroupPolicyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifySecurityGroupPolicyResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifySecurityGroupPolicyResponse struct {
	RequestId string
}
