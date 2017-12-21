package sts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GenerateSessionAccessKeyRequest struct {
	DurationSeconds int64 `position:"Query" name:"DurationSeconds"`
}

func (r GenerateSessionAccessKeyRequest) Invoke(client *sdk.Client) (response *GenerateSessionAccessKeyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GenerateSessionAccessKeyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Sts", "2015-04-01", "GenerateSessionAccessKey", "", "")

	resp := struct {
		*responses.BaseResponse
		GenerateSessionAccessKeyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GenerateSessionAccessKeyResponse

	err = client.DoAction(&req, &resp)
	return
}

type GenerateSessionAccessKeyResponse struct {
	RequestId        string
	SessionAccessKey GenerateSessionAccessKeySessionAccessKey
}

type GenerateSessionAccessKeySessionAccessKey struct {
	SessionAccessKeyId     string
	SessionAccessKeySecret string
	Expiration             string
}
