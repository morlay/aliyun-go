package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GenerateDataKeyRequest struct {
	requests.RpcRequest
	EncryptionContext string `position:"Query" name:"EncryptionContext"`
	KeyId             string `position:"Query" name:"KeyId"`
	KeySpec           string `position:"Query" name:"KeySpec"`
	STSToken          string `position:"Query" name:"STSToken"`
	NumberOfBytes     int    `position:"Query" name:"NumberOfBytes"`
}

func (req *GenerateDataKeyRequest) Invoke(client *sdk.Client) (resp *GenerateDataKeyResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "GenerateDataKey", "kms", "")
	resp = &GenerateDataKeyResponse{}
	err = client.DoAction(req, resp)
	return
}

type GenerateDataKeyResponse struct {
	responses.BaseResponse
	CiphertextBlob string
	KeyId          string
	Plaintext      string
	RequestId      string
}
