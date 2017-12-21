package cf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AuthenticateRequest struct {
	Sig       string `position:"Query" name:"Sig"`
	RemoteIp  string `position:"Query" name:"RemoteIp"`
	AccessKey string `position:"Query" name:"AccessKey"`
	AppKey    string `position:"Query" name:"AppKey"`
	SessionId string `position:"Query" name:"SessionId"`
	Token     string `position:"Query" name:"Token"`
}

func (r AuthenticateRequest) Invoke(client *sdk.Client) (response *AuthenticateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AuthenticateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CF", "2015-12-08", "Authenticate", "cf", "")

	resp := struct {
		*responses.BaseResponse
		AuthenticateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AuthenticateResponse

	err = client.DoAction(&req, &resp)
	return
}

type AuthenticateResponse struct {
	RequestId             string
	Success               bool
	Msg                   string
	Code                  int
	SigAuthenticateResult AuthenticateSigAuthenticateResult
}

type AuthenticateSigAuthenticateResult struct {
	Timestamp int64
	Code      int
	Msg       string
	Risklevel string
}
