package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeUserDomainsRequest struct {
	requests.RpcRequest
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

func (req *DescribeUserDomainsRequest) Invoke(client *sdk.Client) (resp *DescribeUserDomainsResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeUserDomains", "", "")
	resp = &DescribeUserDomainsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeUserDomainsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	PageNumber common.Long
	PageSize   common.Long
	TotalCount common.Long
	Domains    DescribeUserDomainsPageDataList
}

type DescribeUserDomainsPageData struct {
	DomainName      common.String
	Cname           common.String
	CdnType         common.String
	DomainStatus    common.String
	GmtCreated      common.String
	GmtModified     common.String
	Description     common.String
	SourceType      common.String
	SslProtocol     common.String
	ResourceGroupId common.String
	Sandbox         common.String
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

type DescribeUserDomainsSourceList []common.String

func (list *DescribeUserDomainsSourceList) UnmarshalJSON(data []byte) error {
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
