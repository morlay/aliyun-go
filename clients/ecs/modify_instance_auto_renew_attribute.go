package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyInstanceAutoRenewAttributeRequest struct {
	requests.RpcRequest
	Duration             int    `position:"Query" name:"Duration"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	AutoRenew            string `position:"Query" name:"AutoRenew"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	RenewalStatus        string `position:"Query" name:"RenewalStatus"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyInstanceAutoRenewAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyInstanceAutoRenewAttributeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceAutoRenewAttribute", "ecs", "")
	resp = &ModifyInstanceAutoRenewAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyInstanceAutoRenewAttributeResponse struct {
	responses.BaseResponse
	RequestId string
}
