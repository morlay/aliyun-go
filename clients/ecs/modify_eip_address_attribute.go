package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyEipAddressAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyEipAddressAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyEipAddressAttributeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyEipAddressAttribute", "ecs", "")
	resp = &ModifyEipAddressAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyEipAddressAttributeResponse struct {
	responses.BaseResponse
	RequestId string
}
