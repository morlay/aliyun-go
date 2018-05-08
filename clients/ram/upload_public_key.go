package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	PublicKey UploadPublicKeyPublicKey
}

type UploadPublicKeyPublicKey struct {
	PublicKeyId   common.String
	PublicKeySpec common.String
	Status        common.String
	CreateDate    common.String
}
