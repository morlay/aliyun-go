package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DecryptRequest struct {
	CiphertextBlob    string `position:"Query" name:"CiphertextBlob"`
	STSToken          string `position:"Query" name:"STSToken"`
	EncryptionContext string `position:"Query" name:"EncryptionContext"`
}

func (r DecryptRequest) Invoke(client *sdk.Client) (response *DecryptResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DecryptRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Kms", "2016-01-20", "Decrypt", "kms", "")

	resp := struct {
		*responses.BaseResponse
		DecryptResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DecryptResponse

	err = client.DoAction(&req, &resp)
	return
}

type DecryptResponse struct {
	Plaintext string
	KeyId     string
	RequestId string
}
