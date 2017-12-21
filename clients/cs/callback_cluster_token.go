package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CallbackClusterTokenRequest struct {
	requests.RoaRequest
	Token   string `position:"Path" name:"Token"`
	ReqOnce string `position:"Path" name:"ReqOnce"`
}

func (req *CallbackClusterTokenRequest) Invoke(client *sdk.Client) (resp *CallbackClusterTokenResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "CallbackClusterToken", "/token/[Token]/req_once/[ReqOnce]/callback", "", "")
	req.Method = "POST"

	resp = &CallbackClusterTokenResponse{}
	err = client.DoAction(req, resp)
	return
}

type CallbackClusterTokenResponse struct {
	responses.BaseResponse
}
