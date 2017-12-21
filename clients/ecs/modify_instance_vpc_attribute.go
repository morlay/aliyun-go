package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyInstanceVpcAttributeRequest struct {
	requests.RpcRequest
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	PrivateIpAddress     string `position:"Query" name:"PrivateIpAddress"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyInstanceVpcAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyInstanceVpcAttributeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceVpcAttribute", "ecs", "")
	resp = &ModifyInstanceVpcAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyInstanceVpcAttributeResponse struct {
	responses.BaseResponse
	RequestId string
}
