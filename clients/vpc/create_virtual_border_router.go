package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateVirtualBorderRouterRequest struct {
	requests.RpcRequest
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
	VbrOwnerId           int64  `position:"Query" name:"VbrOwnerId"`
}

func (req *CreateVirtualBorderRouterRequest) Invoke(client *sdk.Client) (resp *CreateVirtualBorderRouterResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateVirtualBorderRouter", "vpc", "")
	resp = &CreateVirtualBorderRouterResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateVirtualBorderRouterResponse struct {
	responses.BaseResponse
	RequestId common.String
	VbrId     common.String
}
