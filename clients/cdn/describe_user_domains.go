package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeUserDomainsRequest struct {
	FuncFilter       string `position:"Query" name:"FuncFilter"`
	Sources          string `position:"Query" name:"Sources"`
	DomainName       string `position:"Query" name:"DomainName"`
	OwnerId          int64  `position:"Query" name:"OwnerId"`
	FuncId           string `position:"Query" name:"FuncId"`
	PageNumber       int    `position:"Query" name:"PageNumber"`
	DomainStatus     string `position:"Query" name:"DomainStatus"`
	DomainSearchType string `position:"Query" name:"DomainSearchType"`
	CheckDomainShow  string `position:"Query" name:"CheckDomainShow"`
	ResourceGroupId  string `position:"Query" name:"ResourceGroupId"`
	SecurityToken    string `position:"Query" name:"SecurityToken"`
	CdnType          string `position:"Query" name:"CdnType"`
	PageSize         int    `position:"Query" name:"PageSize"`
}

func (r DescribeUserDomainsRequest) Invoke(client *sdk.Client) (response *DescribeUserDomainsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeUserDomainsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeUserDomains", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeUserDomainsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeUserDomainsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeUserDomainsResponse struct {
	RequestId  string
	PageNumber int64
	PageSize   int64
	TotalCount int64
	Domains    DescribeUserDomainsPageDataList
}

type DescribeUserDomainsPageData struct {
	DomainName      string
	Cname           string
	CdnType         string
	DomainStatus    string
	GmtCreated      string
	GmtModified     string
	Description     string
	SourceType      string
	SslProtocol     string
	ResourceGroupId string
	Sandbox         string
	Sources         DescribeUserDomainsSourceList
}

type DescribeUserDomainsPageDataList []DescribeUserDomainsPageData

func (list *DescribeUserDomainsPageDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeUserDomainsPageData)
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

type DescribeUserDomainsSourceList []string

func (list *DescribeUserDomainsSourceList) UnmarshalJSON(data []byte) error {
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
