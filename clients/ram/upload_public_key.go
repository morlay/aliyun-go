package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UploadPublicKeyRequest struct {
	requests.RpcRequest
	PublicKeySpec string `position:"Query" name:"PublicKeySpec"`
	UserName      string `position:"Query" name:"UserName"`
}

func (req *UploadPublicKeyRequest) Invoke(client *sdk.Client) (resp *UploadPublicKeyResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "UploadPublicKey", "", "")
	resp = &UploadPublicKeyResponse{}
	err = client.DoAction(req, resp)
	return
}

type UploadPublicKeyResponse struct {
	responses.BaseResponse
	RequestId string
	PublicKey UploadPublicKeyPublicKey
}

type UploadPublicKeyPublicKey struct {
	PublicKeyId   string
	PublicKeySpec string
	Status        string
	CreateDate    string
}
