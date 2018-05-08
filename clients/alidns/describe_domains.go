package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDomainsRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	KeyWord      string `position:"Query" name:"KeyWord"`
	GroupId      string `position:"Query" name:"GroupId"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
}

func (req *DescribeDomainsRequest) Invoke(client *sdk.Client) (resp *DescribeDomainsResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDomains", "", "")
	resp = &DescribeDomainsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Long
	PageNumber common.Long
	PageSize   common.Long
	Domains    DescribeDomainsDomainList
}

type DescribeDomainsDomain struct {
	DomainId        common.String
	DomainName      common.String
	PunyCode        common.String
	AliDomain       bool
	RegistrantEmail common.String
	GroupId         common.String
	GroupName       common.String
	InstanceId      common.String
	VersionCode     common.String
	VersionName     common.String
	DnsServers      DescribeDomainsDnsServerList
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

type DescribeDomainsDnsServerList []common.String

func (list *DescribeDomainsDnsServerList) UnmarshalJSON(data []byte) error {
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
