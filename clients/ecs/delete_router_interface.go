package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteRouterInterfaceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	RouterInterfaceId    string `position:"Query" name:"RouterInterfaceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteRouterInterfaceRequest) Invoke(client *sdk.Client) (response *DeleteRouterInterfaceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteRouterInterfaceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteRouterInterface", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DeleteRouterInterfaceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteRouterInterfaceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteRouterInterfaceResponse struct {
	RequestId string
}
