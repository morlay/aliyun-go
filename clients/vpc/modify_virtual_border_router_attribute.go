package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyVirtualBorderRouterAttributeRequest struct {
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
	UserCidr                      string `position:"Query" name:"UserCidr"`
}

func (r ModifyVirtualBorderRouterAttributeRequest) Invoke(client *sdk.Client) (response *ModifyVirtualBorderRouterAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyVirtualBorderRouterAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVirtualBorderRouterAttribute", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		ModifyVirtualBorderRouterAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyVirtualBorderRouterAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyVirtualBorderRouterAttributeResponse struct {
	RequestId string
}
