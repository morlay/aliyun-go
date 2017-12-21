package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DecryptRequest struct {
	requests.RpcRequest
	CiphertextBlob    string `position:"Query" name:"CiphertextBlob"`
	STSToken          string `position:"Query" name:"STSToken"`
	EncryptionContext string `position:"Query" name:"EncryptionContext"`
}

func (req *DecryptRequest) Invoke(client *sdk.Client) (resp *DecryptResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "Decrypt", "kms", "")
	resp = &DecryptResponse{}
	err = client.DoAction(req, resp)
	return
}

type DecryptResponse struct {
	responses.BaseResponse
	Plaintext string
	KeyId     string
	RequestId string
}
