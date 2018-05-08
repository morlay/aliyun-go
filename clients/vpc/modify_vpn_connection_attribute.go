package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId         common.String
	VpnConnectionId   common.String
	CustomerGatewayId common.String
	VpnGatewayId      common.String
	Name              common.String
	Description       common.String
	LocalSubnet       common.String
	RemoteSubnet      common.String
	CreateTime        common.Long
	EffectImmediately bool
	IkeConfig         ModifyVpnConnectionAttributeIkeConfig
	IpsecConfig       ModifyVpnConnectionAttributeIpsecConfig
}

type ModifyVpnConnectionAttributeIkeConfig struct {
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

type ModifyVpnConnectionAttributeIpsecConfig struct {
	IpsecEncAlg   common.String
	IpsecAuthAlg  common.String
	IpsecPfs      common.String
	IpsecLifetime common.Long
}
