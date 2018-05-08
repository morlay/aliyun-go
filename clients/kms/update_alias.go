package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateAliasRequest struct {
	requests.RpcRequest
	AliasName string `position:"Query" name:"AliasName"`
	KeyId     string `position:"Query" name:"KeyId"`
	STSToken  string `position:"Query" name:"STSToken"`
}

func (req *UpdateAliasRequest) Invoke(client *sdk.Client) (resp *UpdateAliasResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "UpdateAlias", "kms", "")
	resp = &UpdateAliasResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateAliasResponse struct {
	responses.BaseResponse
	RequestId common.String
}
