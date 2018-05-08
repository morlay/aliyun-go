package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type EncryptRequest struct {
	requests.RpcRequest
	EncryptionContext string `position:"Query" name:"EncryptionContext"`
	KeyId             string `position:"Query" name:"KeyId"`
	STSToken          string `position:"Query" name:"STSToken"`
	Plaintext         string `position:"Query" name:"Plaintext"`
}

func (req *EncryptRequest) Invoke(client *sdk.Client) (resp *EncryptResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "Encrypt", "kms", "")
	resp = &EncryptResponse{}
	err = client.DoAction(req, resp)
	return
}

type EncryptResponse struct {
	responses.BaseResponse
	CiphertextBlob common.String
	KeyId          common.String
	RequestId      common.String
}
