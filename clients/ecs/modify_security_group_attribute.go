package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifySecurityGroupAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SecurityGroupName    string `position:"Query" name:"SecurityGroupName"`
}

func (req *ModifySecurityGroupAttributeRequest) Invoke(client *sdk.Client) (resp *ModifySecurityGroupAttributeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifySecurityGroupAttribute", "ecs", "")
	resp = &ModifySecurityGroupAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifySecurityGroupAttributeResponse struct {
	responses.BaseResponse
	RequestId string
}
