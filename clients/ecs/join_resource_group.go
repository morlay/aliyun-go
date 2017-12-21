package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type JoinResourceGroupRequest struct {
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceType         string `position:"Query" name:"ResourceType"`
}

func (r JoinResourceGroupRequest) Invoke(client *sdk.Client) (response *JoinResourceGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		JoinResourceGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "JoinResourceGroup", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		JoinResourceGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.JoinResourceGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type JoinResourceGroupResponse struct {
	RequestId string
}
