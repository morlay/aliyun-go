package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RefreshTokenRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (r RefreshTokenRequest) Invoke(client *sdk.Client) (response *RefreshTokenResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RefreshTokenRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CCC", "2017-07-05", "RefreshToken", "ccc", "")

	resp := struct {
		*responses.BaseResponse
		RefreshTokenResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RefreshTokenResponse

	err = client.DoAction(&req, &resp)
	return
}

type RefreshTokenResponse struct {
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
