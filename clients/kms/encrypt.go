package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EncryptRequest struct {
	KeyId             string `position:"Query" name:"KeyId"`
	Plaintext         string `position:"Query" name:"Plaintext"`
	STSToken          string `position:"Query" name:"STSToken"`
	EncryptionContext string `position:"Query" name:"EncryptionContext"`
}

func (r EncryptRequest) Invoke(client *sdk.Client) (response *EncryptResponse, err error) {
	req := struct {
		*requests.RpcRequest
		EncryptRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Kms", "2016-01-20", "Encrypt", "kms", "")

	resp := struct {
		*responses.BaseResponse
		EncryptResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.EncryptResponse

	err = client.DoAction(&req, &resp)
	return
}

type EncryptResponse struct {
	CiphertextBlob string
	KeyId          string
	RequestId      string
}
