package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateAccessKeyRequest struct {
	UserName string `position:"Query" name:"UserName"`
}

func (r CreateAccessKeyRequest) Invoke(client *sdk.Client) (response *CreateAccessKeyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateAccessKeyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "CreateAccessKey", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateAccessKeyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateAccessKeyResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateAccessKeyResponse struct {
	RequestId string
	AccessKey CreateAccessKeyAccessKey
}

type CreateAccessKeyAccessKey struct {
	AccessKeyId     string
	AccessKeySecret string
	Status          string
	CreateDate      string
}
