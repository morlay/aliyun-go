package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CallbackClusterTokenRequest struct {
	Token   string `position:"Path" name:"Token"`
	ReqOnce string `position:"Path" name:"ReqOnce"`
}

func (r CallbackClusterTokenRequest) Invoke(client *sdk.Client) (response *CallbackClusterTokenResponse, err error) {
	req := struct {
		*requests.RoaRequest
		CallbackClusterTokenRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "CallbackClusterToken", "/token/[Token]/req_once/[ReqOnce]/callback", "", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		CallbackClusterTokenResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CallbackClusterTokenResponse

	err = client.DoAction(&req, &resp)
	return
}

type CallbackClusterTokenResponse struct {
}
