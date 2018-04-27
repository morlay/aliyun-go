package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateSslVpnServerRequest struct {
	requests.RpcRequest
	Cipher               string `position:"Query" name:"Cipher"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ClientIpPool         string `position:"Query" name:"ClientIpPool"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Compress             string `position:"Query" name:"Compress"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	LocalSubnet          string `position:"Query" name:"LocalSubnet"`
	Port                 int    `position:"Query" name:"Port"`
	Proto                string `position:"Query" name:"Proto"`
	Name                 string `position:"Query" name:"Name"`
}

func (req *CreateSslVpnServerRequest) Invoke(client *sdk.Client) (resp *CreateSslVpnServerResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateSslVpnServer", "vpc", "")
	resp = &CreateSslVpnServerResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateSslVpnServerResponse struct {
	responses.BaseResponse
	RequestId      string
	SslVpnServerId string
	Name           string
}
