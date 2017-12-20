package cf

import (
	"github.com/morlay/aliyun-go/core"
)

func (c *CfClient) Authenticate(req *AuthenticateArgs) (resp *AuthenticateResponse, err error) {
	resp = &AuthenticateResponse{}
	err = c.Request("Authenticate", req, resp)
	return
}

type AuthenticateSigAuthenticateResult struct {
	Timestamp int64
	Code      int
	Msg       string
	Risklevel string
}
type AuthenticateArgs struct {
	Sig       string
	RemoteIp  string
	AccessKey string
	AppKey    string
	SessionId string
	Token     string
}
type AuthenticateResponse struct {
	RequestId             string
	Success               core.Bool
	Msg                   string
	Code                  int
	SigAuthenticateResult AuthenticateSigAuthenticateResult
}
