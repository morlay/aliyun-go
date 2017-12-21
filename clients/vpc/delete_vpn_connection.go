package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteVpnConnectionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	VpnConnectionId      string `position:"Query" name:"VpnConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteVpnConnectionRequest) Invoke(client *sdk.Client) (response *DeleteVpnConnectionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteVpnConnectionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteVpnConnection", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DeleteVpnConnectionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteVpnConnectionResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteVpnConnectionResponse struct {
	RequestId string
}
