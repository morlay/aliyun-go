package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AssociatePhysicalConnectionToVirtualBorderRouterRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	CircuitCode          string `position:"Query" name:"CircuitCode"`
	VlanId               string `position:"Query" name:"VlanId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VbrId                string `position:"Query" name:"VbrId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PeerGatewayIp        string `position:"Query" name:"PeerGatewayIp"`
	PeeringSubnetMask    string `position:"Query" name:"PeeringSubnetMask"`
	PhysicalConnectionId string `position:"Query" name:"PhysicalConnectionId"`
	LocalGatewayIp       string `position:"Query" name:"LocalGatewayIp"`
}

func (req *AssociatePhysicalConnectionToVirtualBorderRouterRequest) Invoke(client *sdk.Client) (resp *AssociatePhysicalConnectionToVirtualBorderRouterResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "AssociatePhysicalConnectionToVirtualBorderRouter", "vpc", "")
	resp = &AssociatePhysicalConnectionToVirtualBorderRouterResponse{}
	err = client.DoAction(req, resp)
	return
}

type AssociatePhysicalConnectionToVirtualBorderRouterResponse struct {
	responses.BaseResponse
	RequestId common.String
}
