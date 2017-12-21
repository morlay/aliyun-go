package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteVpnConnectionRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	VpnConnectionId      string `position:"Query" name:"VpnConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteVpnConnectionRequest) Invoke(client *sdk.Client) (resp *DeleteVpnConnectionResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteVpnConnection", "vpc", "")
	resp = &DeleteVpnConnectionResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteVpnConnectionResponse struct {
	responses.BaseResponse
	RequestId string
}
