package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdatePublicKeyRequest struct {
	UserPublicKeyId string `position:"Query" name:"UserPublicKeyId"`
	UserName        string `position:"Query" name:"UserName"`
	Status          string `position:"Query" name:"Status"`
}

func (r UpdatePublicKeyRequest) Invoke(client *sdk.Client) (response *UpdatePublicKeyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdatePublicKeyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "UpdatePublicKey", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdatePublicKeyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdatePublicKeyResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdatePublicKeyResponse struct {
	RequestId string
}
