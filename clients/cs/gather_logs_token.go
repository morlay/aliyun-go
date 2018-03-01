package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GatherLogsTokenRequest struct {
	requests.RoaRequest
	Token string `position:"Path" name:"Token"`
}

func (req *GatherLogsTokenRequest) Invoke(client *sdk.Client) (resp *GatherLogsTokenResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "GatherLogsToken", "/token/[Token]/gather_logs", "", "")
	req.Method = "POST"

	resp = &GatherLogsTokenResponse{}
	err = client.DoAction(req, resp)
	return
}

type GatherLogsTokenResponse struct {
	responses.BaseResponse
}
