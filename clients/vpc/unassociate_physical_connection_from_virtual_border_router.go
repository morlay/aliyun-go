package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UnassociatePhysicalConnectionFromVirtualBorderRouterRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	PhysicalConnectionId string `position:"Query" name:"PhysicalConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VbrId                string `position:"Query" name:"VbrId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) Invoke(client *sdk.Client) (resp *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "UnassociatePhysicalConnectionFromVirtualBorderRouter", "vpc", "")
	resp = &UnassociatePhysicalConnectionFromVirtualBorderRouterResponse{}
	err = client.DoAction(req, resp)
	return
}

type UnassociatePhysicalConnectionFromVirtualBorderRouterResponse struct {
	responses.BaseResponse
	RequestId common.String
}
