package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdatePublicKeyRequest struct {
	requests.RpcRequest
	UserPublicKeyId string `position:"Query" name:"UserPublicKeyId"`
	UserName        string `position:"Query" name:"UserName"`
	Status          string `position:"Query" name:"Status"`
}

func (req *UpdatePublicKeyRequest) Invoke(client *sdk.Client) (resp *UpdatePublicKeyResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "UpdatePublicKey", "", "")
	resp = &UpdatePublicKeyResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdatePublicKeyResponse struct {
	responses.BaseResponse
	RequestId string
}
