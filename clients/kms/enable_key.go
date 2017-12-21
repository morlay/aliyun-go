package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EnableKeyRequest struct {
	requests.RpcRequest
	KeyId    string `position:"Query" name:"KeyId"`
	STSToken string `position:"Query" name:"STSToken"`
}

func (req *EnableKeyRequest) Invoke(client *sdk.Client) (resp *EnableKeyResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "EnableKey", "kms", "")
	resp = &EnableKeyResponse{}
	err = client.DoAction(req, resp)
	return
}

type EnableKeyResponse struct {
	responses.BaseResponse
	RequestId string
}
