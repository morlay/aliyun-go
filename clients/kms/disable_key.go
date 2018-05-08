package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DisableKeyRequest struct {
	requests.RpcRequest
	KeyId    string `position:"Query" name:"KeyId"`
	STSToken string `position:"Query" name:"STSToken"`
}

func (req *DisableKeyRequest) Invoke(client *sdk.Client) (resp *DisableKeyResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "DisableKey", "kms", "")
	resp = &DisableKeyResponse{}
	err = client.DoAction(req, resp)
	return
}

type DisableKeyResponse struct {
	responses.BaseResponse
	RequestId common.String
}
