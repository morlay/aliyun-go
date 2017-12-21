package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeVpnConnectionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpnConnectionId      string `position:"Query" name:"VpnConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeVpnConnectionRequest) Invoke(client *sdk.Client) (response *DescribeVpnConnectionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeVpnConnectionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVpnConnection", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DescribeVpnConnectionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeVpnConnectionResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeVpnConnectionResponse struct {
	RequestId         string
	VpnConnectionId   string
	CustomerGatewayId string
	VpnGatewayId      string
	Name              string
	LocalSubnet       string
	RemoteSubnet      string
	CreateTime        int64
	EffectImmediately bool
	Status            string
	IkeConfig         DescribeVpnConnectionIkeConfig
	IpsecConfig       DescribeVpnConnectionIpsecConfig
}

type DescribeVpnConnectionIkeConfig struct {
	Psk         string
	IkeVersion  string
	IkeMode     string
	IkeEncAlg   string
	IkeAuthAlg  string
	IkePfs      string
	IkeLifetime int64
	LocalId     string
	RemoteId    string
}

type DescribeVpnConnectionIpsecConfig struct {
	IpsecEncAlg   string
	IpsecAuthAlg  string
	IpsecPfs      string
	IpsecLifetime int64
}
