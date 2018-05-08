package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ActivateRouterInterfaceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RouterInterfaceId    string `position:"Query" name:"RouterInterfaceId"`
}

func (req *ActivateRouterInterfaceRequest) Invoke(client *sdk.Client) (resp *ActivateRouterInterfaceResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ActivateRouterInterface", "ecs", "")
	resp = &ActivateRouterInterfaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type ActivateRouterInterfaceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
