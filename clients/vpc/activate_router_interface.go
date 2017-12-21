package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ActivateRouterInterfaceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RouterInterfaceId    string `position:"Query" name:"RouterInterfaceId"`
}

func (r ActivateRouterInterfaceRequest) Invoke(client *sdk.Client) (response *ActivateRouterInterfaceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ActivateRouterInterfaceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "ActivateRouterInterface", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		ActivateRouterInterfaceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ActivateRouterInterfaceResponse

	err = client.DoAction(&req, &resp)
	return
}

type ActivateRouterInterfaceResponse struct {
	RequestId string
}
