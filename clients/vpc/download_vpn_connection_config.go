package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DownloadVpnConnectionConfigRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpnConnectionId      string `position:"Query" name:"VpnConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DownloadVpnConnectionConfigRequest) Invoke(client *sdk.Client) (resp *DownloadVpnConnectionConfigResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DownloadVpnConnectionConfig", "vpc", "")
	resp = &DownloadVpnConnectionConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DownloadVpnConnectionConfigResponse struct {
	responses.BaseResponse
	RequestId           common.String
	VpnConnectionConfig DownloadVpnConnectionConfigVpnConnectionConfig
}

type DownloadVpnConnectionConfigVpnConnectionConfig struct {
	LocalSubnet  common.String
	RemoteSubnet common.String
	Local        common.String
	Remote       common.String
	IkeConfig    DownloadVpnConnectionConfigIkeConfig
	IpsecConfig  DownloadVpnConnectionConfigIpsecConfig
}

type DownloadVpnConnectionConfigIkeConfig struct {
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

type DownloadVpnConnectionConfigIpsecConfig struct {
	IpsecEncAlg   common.String
	IpsecAuthAlg  common.String
	IpsecPfs      common.String
	IpsecLifetime common.Long
}
