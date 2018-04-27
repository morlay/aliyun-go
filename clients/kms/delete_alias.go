package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteAliasRequest struct {
	requests.RpcRequest
	AliasName string `position:"Query" name:"AliasName"`
	STSToken  string `position:"Query" name:"STSToken"`
}

func (req *DeleteAliasRequest) Invoke(client *sdk.Client) (resp *DeleteAliasResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "DeleteAlias", "kms", "")
	resp = &DeleteAliasResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteAliasResponse struct {
	responses.BaseResponse
	RequestId string
}
