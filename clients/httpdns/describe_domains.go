package httpdns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainsRequest struct {
	requests.RpcRequest
	PageSize   int64  `position:"Query" name:"PageSize"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	AccountId  string `position:"Query" name:"AccountId"`
}

func (req *DescribeDomainsRequest) Invoke(client *sdk.Client) (resp *DescribeDomainsResponse, err error) {
	req.InitWithApiInfo("Httpdns", "2016-02-01", "DescribeDomains", "", "")
	resp = &DescribeDomainsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainsResponse struct {
	responses.BaseResponse
	RequestId  string
	TotalCount int64
	PageNumber int64
	PageSize   int64
	Domains    DescribeDomainsDomainList
}

type DescribeDomainsDomain struct {
	DomainName string
}

type DescribeDomainsDomainList []DescribeDomainsDomain

func (list *DescribeDomainsDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainsDomain)
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
