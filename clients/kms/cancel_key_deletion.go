package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CancelKeyDeletionRequest struct {
	KeyId    string `position:"Query" name:"KeyId"`
	STSToken string `position:"Query" name:"STSToken"`
}

func (r CancelKeyDeletionRequest) Invoke(client *sdk.Client) (response *CancelKeyDeletionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CancelKeyDeletionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Kms", "2016-01-20", "CancelKeyDeletion", "kms", "")

	resp := struct {
		*responses.BaseResponse
		CancelKeyDeletionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CancelKeyDeletionResponse

	err = client.DoAction(&req, &resp)
	return
}

type CancelKeyDeletionResponse struct {
	RequestId string
}
