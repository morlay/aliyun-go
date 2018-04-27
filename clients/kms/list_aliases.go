package kms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	TotalCount int
	PageNumber int
	PageSize   int
	RequestId  string
	Aliases    ListAliasesAliasList
}

type ListAliasesAlias struct {
	KeyId     string
	AliasName string
	AliasArn  string
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
