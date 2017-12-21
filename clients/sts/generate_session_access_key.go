package sts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GenerateSessionAccessKeyRequest struct {
	requests.RpcRequest
	DurationSeconds int64 `position:"Query" name:"DurationSeconds"`
}

func (req *GenerateSessionAccessKeyRequest) Invoke(client *sdk.Client) (resp *GenerateSessionAccessKeyResponse, err error) {
	req.InitWithApiInfo("Sts", "2015-04-01", "GenerateSessionAccessKey", "", "")
	resp = &GenerateSessionAccessKeyResponse{}
	err = client.DoAction(req, resp)
	return
}

type GenerateSessionAccessKeyResponse struct {
	responses.BaseResponse
	RequestId        string
	SessionAccessKey GenerateSessionAccessKeySessionAccessKey
}

type GenerateSessionAccessKeySessionAccessKey struct {
	SessionAccessKeyId     string
	SessionAccessKeySecret string
	Expiration             string
}
