package cf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AuthenticateRequest struct {
	requests.RpcRequest
	Sig       string `position:"Query" name:"Sig"`
	RemoteIp  string `position:"Query" name:"RemoteIp"`
	AccessKey string `position:"Query" name:"AccessKey"`
	AppKey    string `position:"Query" name:"AppKey"`
	SessionId string `position:"Query" name:"SessionId"`
	Token     string `position:"Query" name:"Token"`
}

func (req *AuthenticateRequest) Invoke(client *sdk.Client) (resp *AuthenticateResponse, err error) {
	req.InitWithApiInfo("CF", "2015-12-08", "Authenticate", "cf", "")
	resp = &AuthenticateResponse{}
	err = client.DoAction(req, resp)
	return
}

type AuthenticateResponse struct {
	responses.BaseResponse
	RequestId             common.String
	Success               bool
	Msg                   common.String
	Code                  common.Integer
	SigAuthenticateResult AuthenticateSigAuthenticateResult
}

type AuthenticateSigAuthenticateResult struct {
	Timestamp common.Long
	Code      common.Integer
	Msg       common.String
	Risklevel common.String
}
