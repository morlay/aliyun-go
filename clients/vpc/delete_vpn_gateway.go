package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteVpnGatewayRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteVpnGatewayRequest) Invoke(client *sdk.Client) (resp *DeleteVpnGatewayResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteVpnGateway", "vpc", "")
	resp = &DeleteVpnGatewayResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteVpnGatewayResponse struct {
	responses.BaseResponse
	RequestId string
}
