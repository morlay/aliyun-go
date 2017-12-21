package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DisableKeyRequest struct {
	KeyId    string `position:"Query" name:"KeyId"`
	STSToken string `position:"Query" name:"STSToken"`
}

func (r DisableKeyRequest) Invoke(client *sdk.Client) (response *DisableKeyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DisableKeyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Kms", "2016-01-20", "DisableKey", "kms", "")

	resp := struct {
		*responses.BaseResponse
		DisableKeyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DisableKeyResponse

	err = client.DoAction(&req, &resp)
	return
}

type DisableKeyResponse struct {
	RequestId string
}
