package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type TerminateVirtualBorderRouterRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	VbrId                string `position:"Query" name:"VbrId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *TerminateVirtualBorderRouterRequest) Invoke(client *sdk.Client) (resp *TerminateVirtualBorderRouterResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "TerminateVirtualBorderRouter", "ecs", "")
	resp = &TerminateVirtualBorderRouterResponse{}
	err = client.DoAction(req, resp)
	return
}

type TerminateVirtualBorderRouterResponse struct {
	responses.BaseResponse
	RequestId string
}
