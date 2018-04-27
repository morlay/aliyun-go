package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifySslVpnClientCertRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SslVpnClientCertId   string `position:"Query" name:"SslVpnClientCertId"`
}

func (req *ModifySslVpnClientCertRequest) Invoke(client *sdk.Client) (resp *ModifySslVpnClientCertResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifySslVpnClientCert", "vpc", "")
	resp = &ModifySslVpnClientCertResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifySslVpnClientCertResponse struct {
	responses.BaseResponse
	RequestId          string
	Name               string
	SslVpnClientCertId string
}
