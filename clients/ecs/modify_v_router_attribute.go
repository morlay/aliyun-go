package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyVRouterAttributeRequest struct {
	requests.RpcRequest
	VRouterName          string `position:"Query" name:"VRouterName"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	VRouterId            string `position:"Query" name:"VRouterId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyVRouterAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyVRouterAttributeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyVRouterAttribute", "ecs", "")
	resp = &ModifyVRouterAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyVRouterAttributeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
