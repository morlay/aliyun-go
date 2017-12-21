package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DecryptKeyRequest struct {
	Rand                 string `position:"Query" name:"Rand"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	CiphertextBlob       string `position:"Query" name:"CiphertextBlob"`
}

func (r DecryptKeyRequest) Invoke(client *sdk.Client) (response *DecryptKeyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DecryptKeyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "DecryptKey", "", "")

	resp := struct {
		*responses.BaseResponse
		DecryptKeyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DecryptKeyResponse

	err = client.DoAction(&req, &resp)
	return
}

type DecryptKeyResponse struct {
	RequestId string
	Plaintext string
	Rand      string
}
