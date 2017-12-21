package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ConnectRouterInterfaceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RouterInterfaceId    string `position:"Query" name:"RouterInterfaceId"`
}

func (r ConnectRouterInterfaceRequest) Invoke(client *sdk.Client) (response *ConnectRouterInterfaceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ConnectRouterInterfaceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ConnectRouterInterface", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ConnectRouterInterfaceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ConnectRouterInterfaceResponse

	err = client.DoAction(&req, &resp)
	return
}

type ConnectRouterInterfaceResponse struct {
	RequestId string
}
