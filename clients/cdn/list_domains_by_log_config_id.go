package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListDomainsByLogConfigIdRequest struct {
	requests.RpcRequest
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigId      string `position:"Query" name:"ConfigId"`
}

func (req *ListDomainsByLogConfigIdRequest) Invoke(client *sdk.Client) (resp *ListDomainsByLogConfigIdResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "ListDomainsByLogConfigId", "", "")
	resp = &ListDomainsByLogConfigIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListDomainsByLogConfigIdResponse struct {
	responses.BaseResponse
	RequestId string
	Domains   ListDomainsByLogConfigIdDomainList
}

type ListDomainsByLogConfigIdDomainList []string

func (list *ListDomainsByLogConfigIdDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
