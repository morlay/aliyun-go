package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CancelKeyDeletionRequest struct {
	requests.RpcRequest
	KeyId    string `position:"Query" name:"KeyId"`
	STSToken string `position:"Query" name:"STSToken"`
}

func (req *CancelKeyDeletionRequest) Invoke(client *sdk.Client) (resp *CancelKeyDeletionResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "CancelKeyDeletion", "kms", "")
	resp = &CancelKeyDeletionResponse{}
	err = client.DoAction(req, resp)
	return
}

type CancelKeyDeletionResponse struct {
	responses.BaseResponse
	RequestId common.String
}
