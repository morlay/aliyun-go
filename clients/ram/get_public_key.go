package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetPublicKeyRequest struct {
	UserPublicKeyId string `position:"Query" name:"UserPublicKeyId"`
	UserName        string `position:"Query" name:"UserName"`
}

func (r GetPublicKeyRequest) Invoke(client *sdk.Client) (response *GetPublicKeyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetPublicKeyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "GetPublicKey", "", "")

	resp := struct {
		*responses.BaseResponse
		GetPublicKeyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetPublicKeyResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetPublicKeyResponse struct {
	RequestId string
	PublicKey GetPublicKeyPublicKey
}

type GetPublicKeyPublicKey struct {
	PublicKeyId   string
	PublicKeySpec string
	Status        string
	CreateDate    string
}
