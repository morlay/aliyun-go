package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteSslVpnServerRequest struct {
	requests.RpcRequest
	SslVpnServerId       string `position:"Query" name:"SslVpnServerId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteSslVpnServerRequest) Invoke(client *sdk.Client) (resp *DeleteSslVpnServerResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteSslVpnServer", "vpc", "")
	resp = &DeleteSslVpnServerResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteSslVpnServerResponse struct {
	responses.BaseResponse
	RequestId string
}
