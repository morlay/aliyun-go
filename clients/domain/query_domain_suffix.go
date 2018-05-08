package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryDomainSuffixRequest struct {
	requests.RpcRequest
	Lang string `position:"Query" name:"Lang"`
}

func (req *QueryDomainSuffixRequest) Invoke(client *sdk.Client) (resp *QueryDomainSuffixResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "QueryDomainSuffix", "", "")
	resp = &QueryDomainSuffixResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryDomainSuffixResponse struct {
	responses.BaseResponse
	RequestId  common.String
	SuffixList QueryDomainSuffixSuffixListList
}

type QueryDomainSuffixSuffixListList []common.String

func (list *QueryDomainSuffixSuffixListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
