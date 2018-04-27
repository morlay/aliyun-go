package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateSslVpnClientCertRequest struct {
	requests.RpcRequest
	SslVpnServerId       string `position:"Query" name:"SslVpnServerId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CreateSslVpnClientCertRequest) Invoke(client *sdk.Client) (resp *CreateSslVpnClientCertResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateSslVpnClientCert", "vpc", "")
	resp = &CreateSslVpnClientCertResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateSslVpnClientCertResponse struct {
	responses.BaseResponse
	RequestId          string
	Name               string
	SslVpnClientCertId string
}
