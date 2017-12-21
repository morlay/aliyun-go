package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeactivateRouterInterfaceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RouterInterfaceId    string `position:"Query" name:"RouterInterfaceId"`
}

func (req *DeactivateRouterInterfaceRequest) Invoke(client *sdk.Client) (resp *DeactivateRouterInterfaceResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeactivateRouterInterface", "vpc", "")
	resp = &DeactivateRouterInterfaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeactivateRouterInterfaceResponse struct {
	responses.BaseResponse
	RequestId string
}
