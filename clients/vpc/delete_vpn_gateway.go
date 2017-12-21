package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteVpnGatewayRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteVpnGatewayRequest) Invoke(client *sdk.Client) (response *DeleteVpnGatewayResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteVpnGatewayRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteVpnGateway", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DeleteVpnGatewayResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteVpnGatewayResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteVpnGatewayResponse struct {
	RequestId string
}
