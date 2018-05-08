package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DecryptRequest struct {
	requests.RpcRequest
	EncryptionContext string `position:"Query" name:"EncryptionContext"`
	STSToken          string `position:"Query" name:"STSToken"`
	CiphertextBlob    string `position:"Query" name:"CiphertextBlob"`
}

func (req *DecryptRequest) Invoke(client *sdk.Client) (resp *DecryptResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "Decrypt", "kms", "")
	resp = &DecryptResponse{}
	err = client.DoAction(req, resp)
	return
}

type DecryptResponse struct {
	responses.BaseResponse
	Plaintext common.String
	KeyId     common.String
	RequestId common.String
}
