package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeletePublicKeyRequest struct {
	UserPublicKeyId string `position:"Query" name:"UserPublicKeyId"`
	UserName        string `position:"Query" name:"UserName"`
}

func (r DeletePublicKeyRequest) Invoke(client *sdk.Client) (response *DeletePublicKeyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeletePublicKeyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "DeletePublicKey", "", "")

	resp := struct {
		*responses.BaseResponse
		DeletePublicKeyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeletePublicKeyResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeletePublicKeyResponse struct {
	RequestId string
}
