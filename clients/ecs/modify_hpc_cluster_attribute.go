package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyHpcClusterAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	HpcClusterId         string `position:"Query" name:"HpcClusterId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Description          string `position:"Query" name:"Description"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Name                 string `position:"Query" name:"Name"`
}

func (req *ModifyHpcClusterAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyHpcClusterAttributeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyHpcClusterAttribute", "ecs", "")
	resp = &ModifyHpcClusterAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyHpcClusterAttributeResponse struct {
	responses.BaseResponse
	RequestId string
}
