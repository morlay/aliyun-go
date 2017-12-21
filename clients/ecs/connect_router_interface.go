package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ConnectRouterInterfaceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RouterInterfaceId    string `position:"Query" name:"RouterInterfaceId"`
}

func (req *ConnectRouterInterfaceRequest) Invoke(client *sdk.Client) (resp *ConnectRouterInterfaceResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ConnectRouterInterface", "ecs", "")
	resp = &ConnectRouterInterfaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type ConnectRouterInterfaceResponse struct {
	responses.BaseResponse
	RequestId string
}
