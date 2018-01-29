package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DecryptKeyRequest struct {
	requests.RpcRequest
	Rand                 string `position:"Query" name:"Rand"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	CiphertextBlob       string `position:"Query" name:"CiphertextBlob"`
}

func (req *DecryptKeyRequest) Invoke(client *sdk.Client) (resp *DecryptKeyResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "DecryptKey", "mts", "")
	resp = &DecryptKeyResponse{}
	err = client.DoAction(req, resp)
	return
}

type DecryptKeyResponse struct {
	responses.BaseResponse
	RequestId string
	Plaintext string
	Rand      string
}
