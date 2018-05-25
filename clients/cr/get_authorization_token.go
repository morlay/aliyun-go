package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetAuthorizationTokenRequest struct {
	requests.RoaRequest
}

func (req *GetAuthorizationTokenRequest) Invoke(client *sdk.Client) (resp *GetAuthorizationTokenResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "GetAuthorizationToken", "/tokens", "", "")
	req.Method = "GET"

	resp = &GetAuthorizationTokenResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetAuthorizationTokenResponse struct {
	responses.BaseResponse
}
