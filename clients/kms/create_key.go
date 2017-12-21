package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateKeyRequest struct {
	Description string `position:"Query" name:"Description"`
	KeyUsage    string `position:"Query" name:"KeyUsage"`
	STSToken    string `position:"Query" name:"STSToken"`
}

func (r CreateKeyRequest) Invoke(client *sdk.Client) (response *CreateKeyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateKeyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Kms", "2016-01-20", "CreateKey", "kms", "")

	resp := struct {
		*responses.BaseResponse
		CreateKeyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateKeyResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateKeyResponse struct {
	RequestId   string
	KeyMetadata CreateKeyKeyMetadata
}

type CreateKeyKeyMetadata struct {
	CreationDate string
	Description  string
	KeyId        string
	KeyState     string
	KeyUsage     string
	DeleteDate   string
	Creator      string
	Arn          string
}
