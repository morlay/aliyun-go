package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListPublicKeysRequest struct {
	requests.RpcRequest
	UserName string `position:"Query" name:"UserName"`
}

func (req *ListPublicKeysRequest) Invoke(client *sdk.Client) (resp *ListPublicKeysResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "ListPublicKeys", "", "")
	resp = &ListPublicKeysResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListPublicKeysResponse struct {
	responses.BaseResponse
	RequestId  common.String
	PublicKeys ListPublicKeysPublicKeyList
}

type ListPublicKeysPublicKey struct {
	PublicKeyId common.String
	Status      common.String
	CreateDate  common.String
}

type ListPublicKeysPublicKeyList []ListPublicKeysPublicKey

func (list *ListPublicKeysPublicKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPublicKeysPublicKey)
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
