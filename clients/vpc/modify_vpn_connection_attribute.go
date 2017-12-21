package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyVpnConnectionAttributeRequest struct {
	requests.RpcRequest
	IkeConfig            string `position:"Query" name:"IkeConfig"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	RemoteSubnet         string `position:"Query" name:"RemoteSubnet"`
	EffectImmediately    string `position:"Query" name:"EffectImmediately"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	IpsecConfig          string `position:"Query" name:"IpsecConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	LocalSubnet          string `position:"Query" name:"LocalSubnet"`
	VpnConnectionId      string `position:"Query" name:"VpnConnectionId"`
	Name                 string `position:"Query" name:"Name"`
}

func (req *ModifyVpnConnectionAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyVpnConnectionAttributeResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVpnConnectionAttribute", "vpc", "")
	resp = &ModifyVpnConnectionAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyVpnConnectionAttributeResponse struct {
	responses.BaseResponse
	RequestId         string
	VpnConnectionId   string
	CustomerGatewayId string
	VpnGatewayId      string
	Name              string
	Description       string
	LocalSubnet       string
	RemoteSubnet      string
	CreateTime        int64
	EffectImmediately bool
	IkeConfig         ModifyVpnConnectionAttributeIkeConfig
	IpsecConfig       ModifyVpnConnectionAttributeIpsecConfig
}

type ModifyVpnConnectionAttributeIkeConfig struct {
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

type ModifyVpnConnectionAttributeIpsecConfig struct {
	IpsecEncAlg   string
	IpsecAuthAlg  string
	IpsecPfs      string
	IpsecLifetime int64
}
