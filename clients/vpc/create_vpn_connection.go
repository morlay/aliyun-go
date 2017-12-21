package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateVpnConnectionRequest struct {
	requests.RpcRequest
	IkeConfig            string `position:"Query" name:"IkeConfig"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	RemoteSubnet         string `position:"Query" name:"RemoteSubnet"`
	EffectImmediately    string `position:"Query" name:"EffectImmediately"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	IpsecConfig          string `position:"Query" name:"IpsecConfig"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CustomerGatewayId    string `position:"Query" name:"CustomerGatewayId"`
	LocalSubnet          string `position:"Query" name:"LocalSubnet"`
	Name                 string `position:"Query" name:"Name"`
}

func (req *CreateVpnConnectionRequest) Invoke(client *sdk.Client) (resp *CreateVpnConnectionResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateVpnConnection", "vpc", "")
	resp = &CreateVpnConnectionResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateVpnConnectionResponse struct {
	responses.BaseResponse
	RequestId       string
	VpnConnectionId string
	Name            string
	CreateTime      int64
}
