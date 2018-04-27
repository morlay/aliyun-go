package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RefreshTokenRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *RefreshTokenRequest) Invoke(client *sdk.Client) (resp *RefreshTokenResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "RefreshToken", "ccc", "")
	resp = &RefreshTokenResponse{}
	err = client.DoAction(req, resp)
	return
}

type RefreshTokenResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	Token          RefreshTokenToken
}

type RefreshTokenToken struct {
	Signature string
	SignData  string
}
