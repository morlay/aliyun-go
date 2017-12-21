package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RevokeClusterTokenRequest struct {
	requests.RoaRequest
	Token string `position:"Path" name:"Token"`
}

func (req *RevokeClusterTokenRequest) Invoke(client *sdk.Client) (resp *RevokeClusterTokenResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "RevokeClusterToken", "/token/[Token]/revoke", "", "")
	req.Method = "DELETE"

	resp = &RevokeClusterTokenResponse{}
	err = client.DoAction(req, resp)
	return
}

type RevokeClusterTokenResponse struct {
	responses.BaseResponse
}
