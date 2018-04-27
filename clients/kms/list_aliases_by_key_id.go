package kms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListAliasesByKeyIdRequest struct {
	requests.RpcRequest
	PageSize   int    `position:"Query" name:"PageSize"`
	KeyId      string `position:"Query" name:"KeyId"`
	STSToken   string `position:"Query" name:"STSToken"`
	PageNumber int    `position:"Query" name:"PageNumber"`
}

func (req *ListAliasesByKeyIdRequest) Invoke(client *sdk.Client) (resp *ListAliasesByKeyIdResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "ListAliasesByKeyId", "kms", "")
	resp = &ListAliasesByKeyIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAliasesByKeyIdResponse struct {
	responses.BaseResponse
	TotalCount int
	PageNumber int
	PageSize   int
	RequestId  string
	Aliases    ListAliasesByKeyIdAliasList
}

type ListAliasesByKeyIdAlias struct {
	KeyId     string
	AliasName string
	AliasArn  string
}

type ListAliasesByKeyIdAliasList []ListAliasesByKeyIdAlias

func (list *ListAliasesByKeyIdAliasList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAliasesByKeyIdAlias)
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
