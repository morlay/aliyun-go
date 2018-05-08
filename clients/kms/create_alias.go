package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateAliasRequest struct {
	requests.RpcRequest
	AliasName string `position:"Query" name:"AliasName"`
	KeyId     string `position:"Query" name:"KeyId"`
	STSToken  string `position:"Query" name:"STSToken"`
}

func (req *CreateAliasRequest) Invoke(client *sdk.Client) (resp *CreateAliasResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "CreateAlias", "kms", "")
	resp = &CreateAliasResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateAliasResponse struct {
	responses.BaseResponse
	RequestId common.String
}
