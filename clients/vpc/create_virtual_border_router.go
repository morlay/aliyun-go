package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateVirtualBorderRouterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	CircuitCode          string `position:"Query" name:"CircuitCode"`
	VlanId               int    `position:"Query" name:"VlanId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PeerGatewayIp        string `position:"Query" name:"PeerGatewayIp"`
	PeeringSubnetMask    string `position:"Query" name:"PeeringSubnetMask"`
	PhysicalConnectionId string `position:"Query" name:"PhysicalConnectionId"`
	Name                 string `position:"Query" name:"Name"`
	LocalGatewayIp       string `position:"Query" name:"LocalGatewayIp"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	VbrOwnerId           int64  `position:"Query" name:"VbrOwnerId"`
}

func (r CreateVirtualBorderRouterRequest) Invoke(client *sdk.Client) (response *CreateVirtualBorderRouterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateVirtualBorderRouterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateVirtualBorderRouter", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		CreateVirtualBorderRouterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateVirtualBorderRouterResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateVirtualBorderRouterResponse struct {
	RequestId string
	VbrId     string
}
