package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeactivateRouterInterfaceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RouterInterfaceId    string `position:"Query" name:"RouterInterfaceId"`
}

func (r DeactivateRouterInterfaceRequest) Invoke(client *sdk.Client) (response *DeactivateRouterInterfaceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeactivateRouterInterfaceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeactivateRouterInterface", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DeactivateRouterInterfaceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeactivateRouterInterfaceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeactivateRouterInterfaceResponse struct {
	RequestId string
}
