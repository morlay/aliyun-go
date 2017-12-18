package httpdns

import "encoding/json"

func (c *HttpdnsClient) GetAccountInfo(req *GetAccountInfoArgs) (resp *GetAccountInfoResponse, err error) {
	resp = &GetAccountInfoResponse{}
	err = c.Request("GetAccountInfo", req, resp)
	return
}

type GetAccountInfoAccountInfo struct {
	AccountId              string
	MonthFreeCount         int
	MonthHttpsResolveCount int
	MonthResolveCount      int
	PackageCount           int
	UserStatus             int
	SignSecret             string
	UnsignedEnabled        bool
	SignedCount            int64
	UnsignedCount          int64
}
type GetAccountInfoArgs struct {
}
type GetAccountInfoResponse struct {
	RequestId   string
	AccountInfo GetAccountInfoAccountInfo
}

func (c *HttpdnsClient) GetResolveStatistics(req *GetResolveStatisticsArgs) (resp *GetResolveStatisticsResponse, err error) {
	resp = &GetResolveStatisticsResponse{}
	err = c.Request("GetResolveStatistics", req, resp)
	return
}

type GetResolveStatisticsDataPoint struct {
	Time  int
	Count int
}
type GetResolveStatisticsArgs struct {
	DomainName   string
	TimeSpan     int
	ProtocolName string
	Granularity  string
}
type GetResolveStatisticsResponse struct {
	RequestId  string
	DataPoints GetResolveStatisticsDataPointList
}

type GetResolveStatisticsDataPointList []GetResolveStatisticsDataPoint

func (list *GetResolveStatisticsDataPointList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetResolveStatisticsDataPoint)
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

func (c *HttpdnsClient) AddDomain(req *AddDomainArgs) (resp *AddDomainResponse, err error) {
	resp = &AddDomainResponse{}
	err = c.Request("AddDomain", req, resp)
	return
}

type AddDomainArgs struct {
	AccountId  string
	DomainName string
}
type AddDomainResponse struct {
	RequestId  string
	DomainName string
}

func (c *HttpdnsClient) DescribeDomains(req *DescribeDomainsArgs) (resp *DescribeDomainsResponse, err error) {
	resp = &DescribeDomainsResponse{}
	err = c.Request("DescribeDomains", req, resp)
	return
}

type DescribeDomainsDomain struct {
	DomainName string
}
type DescribeDomainsArgs struct {
	PageSize   int64
	PageNumber int64
	AccountId  string
}
type DescribeDomainsResponse struct {
	RequestId  string
	TotalCount int64
	PageNumber int64
	PageSize   int64
	Domains    DescribeDomainsDomainList
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

func (c *HttpdnsClient) DeleteDomain(req *DeleteDomainArgs) (resp *DeleteDomainResponse, err error) {
	resp = &DeleteDomainResponse{}
	err = c.Request("DeleteDomain", req, resp)
	return
}

type DeleteDomainArgs struct {
	AccountId  string
	DomainName string
}
type DeleteDomainResponse struct {
	RequestId  string
	DomainName string
}

func (c *HttpdnsClient) ListDomains(req *ListDomainsArgs) (resp *ListDomainsResponse, err error) {
	resp = &ListDomainsResponse{}
	err = c.Request("ListDomains", req, resp)
	return
}

type ListDomainsDomainInfo struct {
	DomainName    string
	Resolved      int64
	ResolvedHttps int64
}
type ListDomainsArgs struct {
	PageSize   int
	PageNumber int
}
type ListDomainsResponse struct {
	RequestId   string
	TotalCount  int64
	PageNumber  int64
	PageSize    int64
	DomainInfos ListDomainsDomainInfoList
}

type ListDomainsDomainInfoList []ListDomainsDomainInfo

func (list *ListDomainsDomainInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListDomainsDomainInfo)
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
