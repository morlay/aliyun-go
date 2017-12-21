package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DownloadVpnConnectionConfigRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpnConnectionId      string `position:"Query" name:"VpnConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DownloadVpnConnectionConfigRequest) Invoke(client *sdk.Client) (response *DownloadVpnConnectionConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DownloadVpnConnectionConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DownloadVpnConnectionConfig", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DownloadVpnConnectionConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DownloadVpnConnectionConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DownloadVpnConnectionConfigResponse struct {
	RequestId           string
	VpnConnectionConfig DownloadVpnConnectionConfigVpnConnectionConfig
}

type DownloadVpnConnectionConfigVpnConnectionConfig struct {
	LocalSubnet  string
	RemoteSubnet string
	Local        string
	Remote       string
	IkeConfig    DownloadVpnConnectionConfigIkeConfig
	IpsecConfig  DownloadVpnConnectionConfigIpsecConfig
}

type DownloadVpnConnectionConfigIkeConfig struct {
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

type DownloadVpnConnectionConfigIpsecConfig struct {
	IpsecEncAlg   string
	IpsecAuthAlg  string
	IpsecPfs      string
	IpsecLifetime int64
}
