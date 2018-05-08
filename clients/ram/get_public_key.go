package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetPublicKeyRequest struct {
	requests.RpcRequest
	UserPublicKeyId string `position:"Query" name:"UserPublicKeyId"`
	UserName        string `position:"Query" name:"UserName"`
}

func (req *GetPublicKeyRequest) Invoke(client *sdk.Client) (resp *GetPublicKeyResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "GetPublicKey", "", "")
	resp = &GetPublicKeyResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetPublicKeyResponse struct {
	responses.BaseResponse
	RequestId common.String
	PublicKey GetPublicKeyPublicKey
}

type GetPublicKeyPublicKey struct {
	PublicKeyId   common.String
	PublicKeySpec common.String
	Status        common.String
	CreateDate    common.String
}
