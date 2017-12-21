package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type JoinResourceGroupRequest struct {
	requests.RpcRequest
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceType         string `position:"Query" name:"ResourceType"`
}

func (req *JoinResourceGroupRequest) Invoke(client *sdk.Client) (resp *JoinResourceGroupResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "JoinResourceGroup", "ecs", "")
	resp = &JoinResourceGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type JoinResourceGroupResponse struct {
	responses.BaseResponse
	RequestId string
}
