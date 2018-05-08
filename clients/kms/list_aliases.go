package kms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListAliasesRequest struct {
	requests.RpcRequest
	PageSize   int    `position:"Query" name:"PageSize"`
	STSToken   string `position:"Query" name:"STSToken"`
	PageNumber int    `position:"Query" name:"PageNumber"`
}

func (req *ListAliasesRequest) Invoke(client *sdk.Client) (resp *ListAliasesResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "ListAliases", "kms", "")
	resp = &ListAliasesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAliasesResponse struct {
	responses.BaseResponse
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	RequestId  common.String
	Aliases    ListAliasesAliasList
}

type ListAliasesAlias struct {
	KeyId     common.String
	AliasName common.String
	AliasArn  common.String
}

type ListAliasesAliasList []ListAliasesAlias

func (list *ListAliasesAliasList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAliasesAlias)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
