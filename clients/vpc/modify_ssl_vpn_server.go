package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifySslVpnServerRequest struct {
	requests.RpcRequest
	Cipher               string `position:"Query" name:"Cipher"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ClientIpPool         string `position:"Query" name:"ClientIpPool"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Compress             string `position:"Query" name:"Compress"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SslVpnServerId       string `position:"Query" name:"SslVpnServerId"`
	LocalSubnet          string `position:"Query" name:"LocalSubnet"`
	Port                 int    `position:"Query" name:"Port"`
	Proto                string `position:"Query" name:"Proto"`
	Name                 string `position:"Query" name:"Name"`
}

func (req *ModifySslVpnServerRequest) Invoke(client *sdk.Client) (resp *ModifySslVpnServerResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifySslVpnServer", "vpc", "")
	resp = &ModifySslVpnServerResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifySslVpnServerResponse struct {
	responses.BaseResponse
	RequestId      common.String
	RegionId       common.String
	SslVpnServerId common.String
	VpnGatewayId   common.String
	Name           common.String
	LocalSubnet    common.String
	ClientIpPool   common.String
	CreateTime     common.Long
	Cipher         common.String
	Proto          common.String
	Port           common.Integer
	Compress       bool
	Connections    common.Integer
	MaxConnections common.Integer
	InternetIp     common.String
}
