package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ImportKeyMaterialRequest struct {
	requests.RpcRequest
	ImportToken           string `position:"Query" name:"ImportToken"`
	EncryptedKeyMaterial  string `position:"Query" name:"EncryptedKeyMaterial"`
	KeyMaterialExpireUnix int64  `position:"Query" name:"KeyMaterialExpireUnix"`
	KeyId                 string `position:"Query" name:"KeyId"`
	STSToken              string `position:"Query" name:"STSToken"`
}

func (req *ImportKeyMaterialRequest) Invoke(client *sdk.Client) (resp *ImportKeyMaterialResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "ImportKeyMaterial", "kms", "")
	resp = &ImportKeyMaterialResponse{}
	err = client.DoAction(req, resp)
	return
}

type ImportKeyMaterialResponse struct {
	responses.BaseResponse
	RequestId string
}
