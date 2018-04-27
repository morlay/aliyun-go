package kms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListKeysRequest struct {
	requests.RpcRequest
	PageSize   int    `position:"Query" name:"PageSize"`
	STSToken   string `position:"Query" name:"STSToken"`
	PageNumber int    `position:"Query" name:"PageNumber"`
}

func (req *ListKeysRequest) Invoke(client *sdk.Client) (resp *ListKeysResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "ListKeys", "kms", "")
	resp = &ListKeysResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListKeysResponse struct {
	responses.BaseResponse
	TotalCount int
	PageNumber int
	PageSize   int
	RequestId  string
	Keys       ListKeysKeyList
}

type ListKeysKey struct {
	KeyId  string
	KeyArn string
}

type ListKeysKeyList []ListKeysKey

func (list *ListKeysKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListKeysKey)
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
