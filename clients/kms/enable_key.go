package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EnableKeyRequest struct {
	KeyId    string `position:"Query" name:"KeyId"`
	STSToken string `position:"Query" name:"STSToken"`
}

func (r EnableKeyRequest) Invoke(client *sdk.Client) (response *EnableKeyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		EnableKeyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Kms", "2016-01-20", "EnableKey", "kms", "")

	resp := struct {
		*responses.BaseResponse
		EnableKeyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.EnableKeyResponse

	err = client.DoAction(&req, &resp)
	return
}

type EnableKeyResponse struct {
	RequestId string
}
