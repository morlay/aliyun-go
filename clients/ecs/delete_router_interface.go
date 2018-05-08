package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteRouterInterfaceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	RouterInterfaceId    string `position:"Query" name:"RouterInterfaceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteRouterInterfaceRequest) Invoke(client *sdk.Client) (resp *DeleteRouterInterfaceResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteRouterInterface", "ecs", "")
	resp = &DeleteRouterInterfaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteRouterInterfaceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
