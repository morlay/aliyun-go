package dcdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDcdnUserDomainsRequest struct {
	requests.RpcRequest
	FuncFilter       string `position:"Query" name:"FuncFilter"`
	DomainName       string `position:"Query" name:"DomainName"`
	OwnerId          int64  `position:"Query" name:"OwnerId"`
	FuncId           string `position:"Query" name:"FuncId"`
	PageNumber       int    `position:"Query" name:"PageNumber"`
	DomainStatus     string `position:"Query" name:"DomainStatus"`
	DomainSearchType string `position:"Query" name:"DomainSearchType"`
	CheckDomainShow  string `position:"Query" name:"CheckDomainShow"`
	ResourceGroupId  string `position:"Query" name:"ResourceGroupId"`
	SecurityToken    string `position:"Query" name:"SecurityToken"`
	PageSize         int    `position:"Query" name:"PageSize"`
}

func (req *DescribeDcdnUserDomainsRequest) Invoke(client *sdk.Client) (resp *DescribeDcdnUserDomainsResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnUserDomains", "dcdn", "")
	resp = &DescribeDcdnUserDomainsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDcdnUserDomainsResponse struct {
	responses.BaseResponse
	RequestId   string
	PageNumber  int64
	PageSize    int64
	TotalCount  int64
	OnlineCount int64
	Domains     DescribeDcdnUserDomainsPageDataList
}

type DescribeDcdnUserDomainsPageData struct {
	DomainName      string
	Cname           string
	DomainStatus    string
	GmtCreated      string
	GmtModified     string
	Description     string
	SSLProtocol     string
	ResourceGroupId string
	Sandbox         string
	Sources         DescribeDcdnUserDomainsSourceList
}

type DescribeDcdnUserDomainsSource struct {
	Type     string
	Content  string
	Port     int
	Priority string
}

type DescribeDcdnUserDomainsPageDataList []DescribeDcdnUserDomainsPageData

func (list *DescribeDcdnUserDomainsPageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnUserDomainsPageData)
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

type DescribeDcdnUserDomainsSourceList []DescribeDcdnUserDomainsSource

func (list *DescribeDcdnUserDomainsSourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnUserDomainsSource)
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
