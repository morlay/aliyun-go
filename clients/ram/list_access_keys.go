package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListAccessKeysRequest struct {
	requests.RpcRequest
	UserName string `position:"Query" name:"UserName"`
}

func (req *ListAccessKeysRequest) Invoke(client *sdk.Client) (resp *ListAccessKeysResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "ListAccessKeys", "", "")
	resp = &ListAccessKeysResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAccessKeysResponse struct {
	responses.BaseResponse
	RequestId  common.String
	AccessKeys ListAccessKeysAccessKeyList
}

type ListAccessKeysAccessKey struct {
	AccessKeyId common.String
	Status      common.String
	CreateDate  common.String
}

type ListAccessKeysAccessKeyList []ListAccessKeysAccessKey

func (list *ListAccessKeysAccessKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAccessKeysAccessKey)
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
