package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RevokeClusterTokenRequest struct {
	Token string `position:"Path" name:"Token"`
}

func (r RevokeClusterTokenRequest) Invoke(client *sdk.Client) (response *RevokeClusterTokenResponse, err error) {
	req := struct {
		*requests.RoaRequest
		RevokeClusterTokenRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "RevokeClusterToken", "/token/[Token]/revoke", "", "")
	req.Method = "DELETE"

	resp := struct {
		*responses.BaseResponse
		RevokeClusterTokenResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RevokeClusterTokenResponse

	err = client.DoAction(&req, &resp)
	return
}

type RevokeClusterTokenResponse struct {
}
