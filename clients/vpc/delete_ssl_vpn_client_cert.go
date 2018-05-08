package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteSslVpnClientCertRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SslVpnClientCertId   string `position:"Query" name:"SslVpnClientCertId"`
}

func (req *DeleteSslVpnClientCertRequest) Invoke(client *sdk.Client) (resp *DeleteSslVpnClientCertResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteSslVpnClientCert", "vpc", "")
	resp = &DeleteSslVpnClientCertResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteSslVpnClientCertResponse struct {
	responses.BaseResponse
	RequestId common.String
}
