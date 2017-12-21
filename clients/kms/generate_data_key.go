package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GenerateDataKeyRequest struct {
	KeyId             string `position:"Query" name:"KeyId"`
	KeySpec           string `position:"Query" name:"KeySpec"`
	NumberOfBytes     int    `position:"Query" name:"NumberOfBytes"`
	STSToken          string `position:"Query" name:"STSToken"`
	EncryptionContext string `position:"Query" name:"EncryptionContext"`
}

func (r GenerateDataKeyRequest) Invoke(client *sdk.Client) (response *GenerateDataKeyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GenerateDataKeyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Kms", "2016-01-20", "GenerateDataKey", "kms", "")

	resp := struct {
		*responses.BaseResponse
		GenerateDataKeyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GenerateDataKeyResponse

	err = client.DoAction(&req, &resp)
	return
}

type GenerateDataKeyResponse struct {
	CiphertextBlob string
	KeyId          string
	Plaintext      string
	RequestId      string
}
