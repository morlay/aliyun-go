package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListPublicKeysRequest struct {
	UserName string `position:"Query" name:"UserName"`
}

func (r ListPublicKeysRequest) Invoke(client *sdk.Client) (response *ListPublicKeysResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListPublicKeysRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "ListPublicKeys", "", "")

	resp := struct {
		*responses.BaseResponse
		ListPublicKeysResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListPublicKeysResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListPublicKeysResponse struct {
	RequestId  string
	PublicKeys ListPublicKeysPublicKeyList
}

type ListPublicKeysPublicKey struct {
	PublicKeyId string
	Status      string
	CreateDate  string
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
