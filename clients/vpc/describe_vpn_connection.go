package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeVpnConnectionRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpnConnectionId      string `position:"Query" name:"VpnConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeVpnConnectionRequest) Invoke(client *sdk.Client) (resp *DescribeVpnConnectionResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVpnConnection", "vpc", "")
	resp = &DescribeVpnConnectionResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeVpnConnectionResponse struct {
	responses.BaseResponse
	RequestId         common.String
	VpnConnectionId   common.String
	CustomerGatewayId common.String
	VpnGatewayId      common.String
	Name              common.String
	LocalSubnet       common.String
	RemoteSubnet      common.String
	CreateTime        common.Long
	EffectImmediately bool
	Status            common.String
	IkeConfig         DescribeVpnConnectionIkeConfig
	IpsecConfig       DescribeVpnConnectionIpsecConfig
}

type DescribeVpnConnectionIkeConfig struct {
	Psk         common.String
	IkeVersion  common.String
	IkeMode     common.String
	IkeEncAlg   common.String
	IkeAuthAlg  common.String
	IkePfs      common.String
	IkeLifetime common.Long
	LocalId     common.String
	RemoteId    common.String
}

type DescribeVpnConnectionIpsecConfig struct {
	IpsecEncAlg   common.String
	IpsecAuthAlg  common.String
	IpsecPfs      common.String
	IpsecLifetime common.Long
}
