package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UploadPublicKeyRequest struct {
	PublicKeySpec string `position:"Query" name:"PublicKeySpec"`
	UserName      string `position:"Query" name:"UserName"`
}

func (r UploadPublicKeyRequest) Invoke(client *sdk.Client) (response *UploadPublicKeyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UploadPublicKeyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "UploadPublicKey", "", "")

	resp := struct {
		*responses.BaseResponse
		UploadPublicKeyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UploadPublicKeyResponse

	err = client.DoAction(&req, &resp)
	return
}

type UploadPublicKeyResponse struct {
	RequestId string
	PublicKey UploadPublicKeyPublicKey
}

type UploadPublicKeyPublicKey struct {
	PublicKeyId   string
	PublicKeySpec string
	Status        string
	CreateDate    string
}
