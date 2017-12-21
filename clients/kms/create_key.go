package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateKeyRequest struct {
	requests.RpcRequest
	Description string `position:"Query" name:"Description"`
	KeyUsage    string `position:"Query" name:"KeyUsage"`
	STSToken    string `position:"Query" name:"STSToken"`
}

func (req *CreateKeyRequest) Invoke(client *sdk.Client) (resp *CreateKeyResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "CreateKey", "kms", "")
	resp = &CreateKeyResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateKeyResponse struct {
	responses.BaseResponse
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
