package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteRouterInterfaceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	RouterInterfaceId    string `position:"Query" name:"RouterInterfaceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteRouterInterfaceRequest) Invoke(client *sdk.Client) (resp *DeleteRouterInterfaceResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteRouterInterface", "vpc", "")
	resp = &DeleteRouterInterfaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteRouterInterfaceResponse struct {
	responses.BaseResponse
	RequestId string
}
