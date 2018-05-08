package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyVirtualBorderRouterAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId               int64  `position:"Query" name:"ResourceOwnerId"`
	CircuitCode                   string `position:"Query" name:"CircuitCode"`
	AssociatedPhysicalConnections string `position:"Query" name:"AssociatedPhysicalConnections"`
	VlanId                        int    `position:"Query" name:"VlanId"`
	ResourceOwnerAccount          string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken                   string `position:"Query" name:"ClientToken"`
	OwnerAccount                  string `position:"Query" name:"OwnerAccount"`
	Description                   string `position:"Query" name:"Description"`
	VbrId                         string `position:"Query" name:"VbrId"`
	OwnerId                       int64  `position:"Query" name:"OwnerId"`
	PeerGatewayIp                 string `position:"Query" name:"PeerGatewayIp"`
	PeeringSubnetMask             string `position:"Query" name:"PeeringSubnetMask"`
	Name                          string `position:"Query" name:"Name"`
	LocalGatewayIp                string `position:"Query" name:"LocalGatewayIp"`
}

func (req *ModifyVirtualBorderRouterAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyVirtualBorderRouterAttributeResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVirtualBorderRouterAttribute", "vpc", "")
	resp = &ModifyVirtualBorderRouterAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyVirtualBorderRouterAttributeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
