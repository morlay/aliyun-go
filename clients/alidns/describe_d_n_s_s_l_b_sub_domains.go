package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDNSSLBSubDomainsRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
}

func (req *DescribeDNSSLBSubDomainsRequest) Invoke(client *sdk.Client) (resp *DescribeDNSSLBSubDomainsResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDNSSLBSubDomains", "", "")
	resp = &DescribeDNSSLBSubDomainsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDNSSLBSubDomainsResponse struct {
	responses.BaseResponse
	RequestId     string
	TotalCount    int64
	PageNumber    int64
	PageSize      int64
	SlbSubDomains DescribeDNSSLBSubDomainsSlbSubDomainList
}

type DescribeDNSSLBSubDomainsSlbSubDomain struct {
	SubDomain   string
	RecordCount int64
	Open        bool
}

type DescribeDNSSLBSubDomainsSlbSubDomainList []DescribeDNSSLBSubDomainsSlbSubDomain

func (list *DescribeDNSSLBSubDomainsSlbSubDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDNSSLBSubDomainsSlbSubDomain)
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
