package vpc

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
	VbrId                string `position:"Query" name:"VbrId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *TerminateVirtualBorderRouterRequest) Invoke(client *sdk.Client) (resp *TerminateVirtualBorderRouterResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "TerminateVirtualBorderRouter", "vpc", "")
	resp = &TerminateVirtualBorderRouterResponse{}
	err = client.DoAction(req, resp)
	return
}

type TerminateVirtualBorderRouterResponse struct {
	responses.BaseResponse
	RequestId string
}
