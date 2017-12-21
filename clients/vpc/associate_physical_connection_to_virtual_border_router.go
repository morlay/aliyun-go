package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AssociatePhysicalConnectionToVirtualBorderRouterRequest struct {
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
	UserCidr             string `position:"Query" name:"UserCidr"`
}

func (r AssociatePhysicalConnectionToVirtualBorderRouterRequest) Invoke(client *sdk.Client) (response *AssociatePhysicalConnectionToVirtualBorderRouterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AssociatePhysicalConnectionToVirtualBorderRouterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "AssociatePhysicalConnectionToVirtualBorderRouter", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		AssociatePhysicalConnectionToVirtualBorderRouterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AssociatePhysicalConnectionToVirtualBorderRouterResponse

	err = client.DoAction(&req, &resp)
	return
}

type AssociatePhysicalConnectionToVirtualBorderRouterResponse struct {
	RequestId string
}
