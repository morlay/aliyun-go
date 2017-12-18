package alidns

import "encoding/json"

func (c *AlidnsClient) RetrievalDomainName(req *RetrievalDomainNameArgs) (resp *RetrievalDomainNameResponse, err error) {
	resp = &RetrievalDomainNameResponse{}
	err = c.Request("RetrievalDomainName", req, resp)
	return
}

type RetrievalDomainNameArgs struct {
	Lang         string
	UserClientIp string
	DomainName   string
}
type RetrievalDomainNameResponse struct {
	RequestId  string
	DomainName string
	WhoisEmail string
}

func (c *AlidnsClient) UpdateDNSSLBWeight(req *UpdateDNSSLBWeightArgs) (resp *UpdateDNSSLBWeightResponse, err error) {
	resp = &UpdateDNSSLBWeightResponse{}
	err = c.Request("UpdateDNSSLBWeight", req, resp)
	return
}

type UpdateDNSSLBWeightArgs struct {
	Lang         string
	UserClientIp string
	RecordId     string
	Weight       int
}
type UpdateDNSSLBWeightResponse struct {
	RequestId string
	RecordId  string
	Weight    int
}

func (c *AlidnsClient) DeleteBatchDomains(req *DeleteBatchDomainsArgs) (resp *DeleteBatchDomainsResponse, err error) {
	resp = &DeleteBatchDomainsResponse{}
	err = c.Request("DeleteBatchDomains", req, resp)
	return
}

type DeleteBatchDomainsArgs struct {
	Lang         string
	UserClientIp string
	Domains      string
}
type DeleteBatchDomainsResponse struct {
	RequestId string
	TraceId   string
}

func (c *AlidnsClient) UpdateDomainRecord(req *UpdateDomainRecordArgs) (resp *UpdateDomainRecordResponse, err error) {
	resp = &UpdateDomainRecordResponse{}
	err = c.Request("UpdateDomainRecord", req, resp)
	return
}

type UpdateDomainRecordArgs struct {
	Lang         string
	UserClientIp string
	RecordId     string
	RR           string
	Type         string
	Value        string
	TTL          int64
	Priority     int64
	Line         string
}
type UpdateDomainRecordResponse struct {
	RequestId string
	RecordId  string
}

func (c *AlidnsClient) DescribeDnsProductInstances(req *DescribeDnsProductInstancesArgs) (resp *DescribeDnsProductInstancesResponse, err error) {
	resp = &DescribeDnsProductInstancesResponse{}
	err = c.Request("DescribeDnsProductInstances", req, resp)
	return
}

type DescribeDnsProductInstancesDnsProduct struct {
	InstanceId            string
	VersionCode           string
	VersionName           string
	StartTime             string
	EndTime               string
	StartTimestamp        int64
	EndTimestamp          int64
	Domain                string
	BindCount             int64
	BindUsedCount         int64
	TTLMinValue           int64
	SubDomainLevel        int64
	DnsSLBCount           int64
	URLForwardCount       int64
	DDosDefendFlow        int64
	DDosDefendQuery       int64
	OverseaDDosDefendFlow int64
	SearchEngineLines     string
	ISPLines              string
	ISPRegionLines        string
	OverseaLine           string
	MonitorNodeCount      int64
	MonitorFrequency      int64
	MonitorTaskCount      int64
}
type DescribeDnsProductInstancesArgs struct {
	Lang         string
	UserClientIp string
	PageNumber   int64
	PageSize     int64
	VersionCode  string
}
type DescribeDnsProductInstancesResponse struct {
	RequestId   string
	TotalCount  int64
	PageNumber  int64
	PageSize    int64
	DnsProducts DescribeDnsProductInstancesDnsProductList
}

type DescribeDnsProductInstancesDnsProductList []DescribeDnsProductInstancesDnsProduct

func (list *DescribeDnsProductInstancesDnsProductList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDnsProductInstancesDnsProduct)
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

func (c *AlidnsClient) DescribeSubDomainRecords(req *DescribeSubDomainRecordsArgs) (resp *DescribeSubDomainRecordsResponse, err error) {
	resp = &DescribeSubDomainRecordsResponse{}
	err = c.Request("DescribeSubDomainRecords", req, resp)
	return
}

type DescribeSubDomainRecordsRecord struct {
	DomainName string
	RecordId   string
	RR         string
	Type       string
	Value      string
	TTL        int64
	Priority   int64
	Line       string
	Status     string
	Locked     bool
	Weight     int
}
type DescribeSubDomainRecordsArgs struct {
	Lang         string
	UserClientIp string
	SubDomain    string
	PageNumber   int64
	PageSize     int64
	Type         string
}
type DescribeSubDomainRecordsResponse struct {
	RequestId     string
	TotalCount    int64
	PageNumber    int64
	PageSize      int64
	DomainRecords DescribeSubDomainRecordsRecordList
}

type DescribeSubDomainRecordsRecordList []DescribeSubDomainRecordsRecord

func (list *DescribeSubDomainRecordsRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSubDomainRecordsRecord)
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

func (c *AlidnsClient) DescribeDomainInfo(req *DescribeDomainInfoArgs) (resp *DescribeDomainInfoResponse, err error) {
	resp = &DescribeDomainInfoResponse{}
	err = c.Request("DescribeDomainInfo", req, resp)
	return
}

type DescribeDomainInfoArgs struct {
	Lang         string
	UserClientIp string
	DomainName   string
}
type DescribeDomainInfoResponse struct {
	RequestId       string
	DomainId        string
	DomainName      string
	PunyCode        string
	AliDomain       bool
	RegistrantEmail string
	GroupId         string
	GroupName       string
	InstanceId      string
	VersionCode     string
	VersionName     string
	DnsServers      DescribeDomainInfoDnsServerList
}

type DescribeDomainInfoDnsServerList []string

func (list *DescribeDomainInfoDnsServerList) UnmarshalJSON(data []byte) error {
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

func (c *AlidnsClient) SetDomainRecordStatus(req *SetDomainRecordStatusArgs) (resp *SetDomainRecordStatusResponse, err error) {
	resp = &SetDomainRecordStatusResponse{}
	err = c.Request("SetDomainRecordStatus", req, resp)
	return
}

type SetDomainRecordStatusArgs struct {
	Lang         string
	UserClientIp string
	RecordId     string
	Status       string
}
type SetDomainRecordStatusResponse struct {
	RequestId string
	RecordId  string
	Status    string
}

func (c *AlidnsClient) DescribeDnsProductInstance(req *DescribeDnsProductInstanceArgs) (resp *DescribeDnsProductInstanceResponse, err error) {
	resp = &DescribeDnsProductInstanceResponse{}
	err = c.Request("DescribeDnsProductInstance", req, resp)
	return
}

type DescribeDnsProductInstanceArgs struct {
	Lang         string
	UserClientIp string
	InstanceId   string
}
type DescribeDnsProductInstanceResponse struct {
	RequestId             string
	InstanceId            string
	VersionCode           string
	VersionName           string
	StartTime             string
	StartTimestamp        int64
	EndTime               string
	EndTimestamp          int64
	Domain                string
	BindCount             int64
	BindUsedCount         int64
	TTLMinValue           int64
	SubDomainLevel        int64
	DnsSLBCount           int64
	URLForwardCount       int64
	DDosDefendFlow        int64
	DDosDefendQuery       int64
	OverseaDDosDefendFlow int64
	SearchEngineLines     string
	ISPLines              string
	ISPRegionLines        string
	OverseaLine           string
	MonitorNodeCount      int64
	MonitorFrequency      int64
	MonitorTaskCount      int64
	DnsServers            DescribeDnsProductInstanceDnsServerList
}

type DescribeDnsProductInstanceDnsServerList []string

func (list *DescribeDnsProductInstanceDnsServerList) UnmarshalJSON(data []byte) error {
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

func (c *AlidnsClient) DescribeBatchResult(req *DescribeBatchResultArgs) (resp *DescribeBatchResultResponse, err error) {
	resp = &DescribeBatchResultResponse{}
	err = c.Request("DescribeBatchResult", req, resp)
	return
}

type DescribeBatchResultFailResult struct {
	BatchIndex string
	ErrorCode  string
}
type DescribeBatchResultArgs struct {
	Lang         string
	UserClientIp string
	TraceId      string
}
type DescribeBatchResultResponse struct {
	RequestId     string
	TraceId       string
	Status        int64
	BatchCount    int64
	SuccessNumber int64
	FailResults   DescribeBatchResultFailResultList
}

type DescribeBatchResultFailResultList []DescribeBatchResultFailResult

func (list *DescribeBatchResultFailResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBatchResultFailResult)
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

func (c *AlidnsClient) DescribeDomainWhoisInfo(req *DescribeDomainWhoisInfoArgs) (resp *DescribeDomainWhoisInfoResponse, err error) {
	resp = &DescribeDomainWhoisInfoResponse{}
	err = c.Request("DescribeDomainWhoisInfo", req, resp)
	return
}

type DescribeDomainWhoisInfoArgs struct {
	Lang         string
	UserClientIp string
	DomainName   string
	GroupId      string
}
type DescribeDomainWhoisInfoResponse struct {
	RequestId        string
	RegistrantName   string
	RegistrantEmail  string
	Registrar        string
	RegistrationDate string
	ExpirationDate   string
	StatusList       DescribeDomainWhoisInfoStatusListList
	DnsServers       DescribeDomainWhoisInfoDnsServerList
}

type DescribeDomainWhoisInfoStatusListList []string

func (list *DescribeDomainWhoisInfoStatusListList) UnmarshalJSON(data []byte) error {
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

type DescribeDomainWhoisInfoDnsServerList []string

func (list *DescribeDomainWhoisInfoDnsServerList) UnmarshalJSON(data []byte) error {
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

func (c *AlidnsClient) DescribeDomainRecords(req *DescribeDomainRecordsArgs) (resp *DescribeDomainRecordsResponse, err error) {
	resp = &DescribeDomainRecordsResponse{}
	err = c.Request("DescribeDomainRecords", req, resp)
	return
}

type DescribeDomainRecordsRecord struct {
	DomainName string
	RecordId   string
	RR         string
	Type       string
	Value      string
	TTL        int64
	Priority   int64
	Line       string
	Status     string
	Locked     bool
	Weight     int
}
type DescribeDomainRecordsArgs struct {
	Lang         string
	UserClientIp string
	DomainName   string
	PageNumber   int64
	PageSize     int64
	RRKeyWord    string
	TypeKeyWord  string
	ValueKeyWord string
}
type DescribeDomainRecordsResponse struct {
	RequestId     string
	TotalCount    int64
	PageNumber    int64
	PageSize      int64
	DomainRecords DescribeDomainRecordsRecordList
}

type DescribeDomainRecordsRecordList []DescribeDomainRecordsRecord

func (list *DescribeDomainRecordsRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainRecordsRecord)
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

func (c *AlidnsClient) DeleteSubDomainRecords(req *DeleteSubDomainRecordsArgs) (resp *DeleteSubDomainRecordsResponse, err error) {
	resp = &DeleteSubDomainRecordsResponse{}
	err = c.Request("DeleteSubDomainRecords", req, resp)
	return
}

type DeleteSubDomainRecordsArgs struct {
	Lang         string
	UserClientIp string
	DomainName   string
	RR           string
	Type         string
}
type DeleteSubDomainRecordsResponse struct {
	RequestId  string
	RR         string
	TotalCount string
}

func (c *AlidnsClient) AddDomain(req *AddDomainArgs) (resp *AddDomainResponse, err error) {
	resp = &AddDomainResponse{}
	err = c.Request("AddDomain", req, resp)
	return
}

type AddDomainArgs struct {
	Lang         string
	UserClientIp string
	DomainName   string
	GroupId      string
}
type AddDomainResponse struct {
	RequestId  string
	DomainId   string
	DomainName string
	PunyCode   string
	GroupId    string
	GroupName  string
	DnsServers AddDomainDnsServerList
}

type AddDomainDnsServerList []string

func (list *AddDomainDnsServerList) UnmarshalJSON(data []byte) error {
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

func (c *AlidnsClient) DescribeDomains(req *DescribeDomainsArgs) (resp *DescribeDomainsResponse, err error) {
	resp = &DescribeDomainsResponse{}
	err = c.Request("DescribeDomains", req, resp)
	return
}

type DescribeDomainsDomain struct {
	DomainId        string
	DomainName      string
	PunyCode        string
	AliDomain       bool
	RegistrantEmail string
	GroupId         string
	GroupName       string
	InstanceId      string
	VersionCode     string
	VersionName     string
	DnsServers      DescribeDomainsDnsServerList
}
type DescribeDomainsArgs struct {
	Lang         string
	UserClientIp string
	KeyWord      string
	GroupId      string
	PageNumber   int64
	PageSize     int64
}
type DescribeDomainsResponse struct {
	RequestId  string
	TotalCount int64
	PageNumber int64
	PageSize   int64
	Domains    DescribeDomainsDomainList
}

type DescribeDomainsDnsServerList []string

func (list *DescribeDomainsDnsServerList) UnmarshalJSON(data []byte) error {
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

func (c *AlidnsClient) DeleteBatchDomainRecords(req *DeleteBatchDomainRecordsArgs) (resp *DeleteBatchDomainRecordsResponse, err error) {
	resp = &DeleteBatchDomainRecordsResponse{}
	err = c.Request("DeleteBatchDomainRecords", req, resp)
	return
}

type DeleteBatchDomainRecordsArgs struct {
	Lang         string
	UserClientIp string
	Records      string
}
type DeleteBatchDomainRecordsResponse struct {
	RequestId string
	TraceId   string
}

func (c *AlidnsClient) DeleteDomain(req *DeleteDomainArgs) (resp *DeleteDomainResponse, err error) {
	resp = &DeleteDomainResponse{}
	err = c.Request("DeleteDomain", req, resp)
	return
}

type DeleteDomainArgs struct {
	Lang         string
	UserClientIp string
	DomainName   string
}
type DeleteDomainResponse struct {
	RequestId  string
	DomainName string
}

func (c *AlidnsClient) DescribeDomainNs(req *DescribeDomainNsArgs) (resp *DescribeDomainNsResponse, err error) {
	resp = &DescribeDomainNsResponse{}
	err = c.Request("DescribeDomainNs", req, resp)
	return
}

type DescribeDomainNsArgs struct {
	Lang         string
	UserClientIp string
	DomainName   string
}
type DescribeDomainNsResponse struct {
	RequestId     string
	AllAliDns     bool
	IncludeAliDns bool
	DnsServers    DescribeDomainNsDnsServerList
}

type DescribeDomainNsDnsServerList []string

func (list *DescribeDomainNsDnsServerList) UnmarshalJSON(data []byte) error {
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

func (c *AlidnsClient) DescribeRecordLogs(req *DescribeRecordLogsArgs) (resp *DescribeRecordLogsResponse, err error) {
	resp = &DescribeRecordLogsResponse{}
	err = c.Request("DescribeRecordLogs", req, resp)
	return
}

type DescribeRecordLogsRecordLog struct {
	ActionTime string
	Action     string
	Message    string
	ClientIp   string
}
type DescribeRecordLogsArgs struct {
	Lang         string
	UserClientIp string
	DomainName   string
	PageNumber   int64
	PageSize     int64
	KeyWord      string
}
type DescribeRecordLogsResponse struct {
	RequestId  string
	TotalCount int64
	PageNumber int64
	PageSize   int64
	RecordLogs DescribeRecordLogsRecordLogList
}

type DescribeRecordLogsRecordLogList []DescribeRecordLogsRecordLog

func (list *DescribeRecordLogsRecordLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRecordLogsRecordLog)
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

func (c *AlidnsClient) ChangeDomainGroup(req *ChangeDomainGroupArgs) (resp *ChangeDomainGroupResponse, err error) {
	resp = &ChangeDomainGroupResponse{}
	err = c.Request("ChangeDomainGroup", req, resp)
	return
}

type ChangeDomainGroupArgs struct {
	Lang         string
	UserClientIp string
	DomainName   string
	GroupId      string
}
type ChangeDomainGroupResponse struct {
	RequestId string
	GroupId   string
	GroupName string
}

func (c *AlidnsClient) UpdateBatchDomainRecords(req *UpdateBatchDomainRecordsArgs) (resp *UpdateBatchDomainRecordsResponse, err error) {
	resp = &UpdateBatchDomainRecordsResponse{}
	err = c.Request("UpdateBatchDomainRecords", req, resp)
	return
}

type UpdateBatchDomainRecordsArgs struct {
	Lang         string
	UserClientIp string
	Records      string
}
type UpdateBatchDomainRecordsResponse struct {
	RequestId string
	TraceId   string
}

func (c *AlidnsClient) UpdateDomainGroup(req *UpdateDomainGroupArgs) (resp *UpdateDomainGroupResponse, err error) {
	resp = &UpdateDomainGroupResponse{}
	err = c.Request("UpdateDomainGroup", req, resp)
	return
}

type UpdateDomainGroupArgs struct {
	Lang         string
	UserClientIp string
	GroupId      string
	GroupName    string
}
type UpdateDomainGroupResponse struct {
	RequestId string
	GroupId   string
	GroupName string
}

func (c *AlidnsClient) ApplyForRetrievalDomainName(req *ApplyForRetrievalDomainNameArgs) (resp *ApplyForRetrievalDomainNameResponse, err error) {
	resp = &ApplyForRetrievalDomainNameResponse{}
	err = c.Request("ApplyForRetrievalDomainName", req, resp)
	return
}

type ApplyForRetrievalDomainNameArgs struct {
	Lang         string
	UserClientIp string
	DomainName   string
}
type ApplyForRetrievalDomainNameResponse struct {
	RequestId  string
	DomainName string
}

func (c *AlidnsClient) DeleteDomainRecord(req *DeleteDomainRecordArgs) (resp *DeleteDomainRecordResponse, err error) {
	resp = &DeleteDomainRecordResponse{}
	err = c.Request("DeleteDomainRecord", req, resp)
	return
}

type DeleteDomainRecordArgs struct {
	Lang         string
	UserClientIp string
	RecordId     string
}
type DeleteDomainRecordResponse struct {
	RequestId string
	RecordId  string
}

func (c *AlidnsClient) AddBatchDomainRecords(req *AddBatchDomainRecordsArgs) (resp *AddBatchDomainRecordsResponse, err error) {
	resp = &AddBatchDomainRecordsResponse{}
	err = c.Request("AddBatchDomainRecords", req, resp)
	return
}

type AddBatchDomainRecordsArgs struct {
	Lang         string
	UserClientIp string
	Records      string
}
type AddBatchDomainRecordsResponse struct {
	RequestId string
	TraceId   string
}

func (c *AlidnsClient) AddDomainRecord(req *AddDomainRecordArgs) (resp *AddDomainRecordResponse, err error) {
	resp = &AddDomainRecordResponse{}
	err = c.Request("AddDomainRecord", req, resp)
	return
}

type AddDomainRecordArgs struct {
	Lang         string
	UserClientIp string
	DomainName   string
	RR           string
	Type         string
	Value        string
	TTL          int64
	Priority     int64
	Line         string
}
type AddDomainRecordResponse struct {
	RequestId string
	RecordId  string
}

func (c *AlidnsClient) CheckDomainRecord(req *CheckDomainRecordArgs) (resp *CheckDomainRecordResponse, err error) {
	resp = &CheckDomainRecordResponse{}
	err = c.Request("CheckDomainRecord", req, resp)
	return
}

type CheckDomainRecordArgs struct {
	Lang         string
	UserClientIp string
	DomainName   string
	RR           string
	Type         string
	Value        string
}
type CheckDomainRecordResponse struct {
	RequestId string
	IsExist   bool
}

func (c *AlidnsClient) DeleteDomainGroup(req *DeleteDomainGroupArgs) (resp *DeleteDomainGroupResponse, err error) {
	resp = &DeleteDomainGroupResponse{}
	err = c.Request("DeleteDomainGroup", req, resp)
	return
}

type DeleteDomainGroupArgs struct {
	Lang         string
	UserClientIp string
	GroupId      string
}
type DeleteDomainGroupResponse struct {
	RequestId string
	GroupName string
}

func (c *AlidnsClient) GetMainDomainName(req *GetMainDomainNameArgs) (resp *GetMainDomainNameResponse, err error) {
	resp = &GetMainDomainNameResponse{}
	err = c.Request("GetMainDomainName", req, resp)
	return
}

type GetMainDomainNameArgs struct {
	Lang         string
	UserClientIp string
	InputString  string
}
type GetMainDomainNameResponse struct {
	RequestId   string
	DomainName  string
	RR          string
	DomainLevel int64
}

func (c *AlidnsClient) ModifyHichinaDomainDNS(req *ModifyHichinaDomainDNSArgs) (resp *ModifyHichinaDomainDNSResponse, err error) {
	resp = &ModifyHichinaDomainDNSResponse{}
	err = c.Request("ModifyHichinaDomainDNS", req, resp)
	return
}

type ModifyHichinaDomainDNSArgs struct {
	Lang         string
	UserClientIp string
	DomainName   string
}
type ModifyHichinaDomainDNSResponse struct {
	RequestId          string
	OriginalDnsServers ModifyHichinaDomainDNSOriginalDnsServerList
	NewDnsServers      ModifyHichinaDomainDNSNewDnsServerList
}

type ModifyHichinaDomainDNSOriginalDnsServerList []string

func (list *ModifyHichinaDomainDNSOriginalDnsServerList) UnmarshalJSON(data []byte) error {
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

type ModifyHichinaDomainDNSNewDnsServerList []string

func (list *ModifyHichinaDomainDNSNewDnsServerList) UnmarshalJSON(data []byte) error {
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

func (c *AlidnsClient) DescribeDomainGroups(req *DescribeDomainGroupsArgs) (resp *DescribeDomainGroupsResponse, err error) {
	resp = &DescribeDomainGroupsResponse{}
	err = c.Request("DescribeDomainGroups", req, resp)
	return
}

type DescribeDomainGroupsDomainGroup struct {
	GroupId   string
	GroupName string
}
type DescribeDomainGroupsArgs struct {
	Lang         string
	UserClientIp string
	KeyWord      string
	PageNumber   int64
	PageSize     int64
}
type DescribeDomainGroupsResponse struct {
	RequestId    string
	TotalCount   int64
	PageNumber   int64
	PageSize     int64
	DomainGroups DescribeDomainGroupsDomainGroupList
}

type DescribeDomainGroupsDomainGroupList []DescribeDomainGroupsDomainGroup

func (list *DescribeDomainGroupsDomainGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainGroupsDomainGroup)
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

func (c *AlidnsClient) DescribeDomainRecordInfo(req *DescribeDomainRecordInfoArgs) (resp *DescribeDomainRecordInfoResponse, err error) {
	resp = &DescribeDomainRecordInfoResponse{}
	err = c.Request("DescribeDomainRecordInfo", req, resp)
	return
}

type DescribeDomainRecordInfoArgs struct {
	Lang         string
	UserClientIp string
	RecordId     string
}
type DescribeDomainRecordInfoResponse struct {
	RequestId  string
	DomainId   string
	DomainName string
	PunyCode   string
	GroupId    string
	GroupName  string
	RecordId   string
	RR         string
	Type       string
	Value      string
	TTL        int64
	Priority   int64
	Line       string
	Status     string
	Locked     bool
}

func (c *AlidnsClient) DescribeDNSSLBSubDomains(req *DescribeDNSSLBSubDomainsArgs) (resp *DescribeDNSSLBSubDomainsResponse, err error) {
	resp = &DescribeDNSSLBSubDomainsResponse{}
	err = c.Request("DescribeDNSSLBSubDomains", req, resp)
	return
}

type DescribeDNSSLBSubDomainsSlbSubDomain struct {
	SubDomain   string
	RecordCount int64
	Open        bool
}
type DescribeDNSSLBSubDomainsArgs struct {
	Lang         string
	UserClientIp string
	DomainName   string
	PageNumber   int64
	PageSize     int64
}
type DescribeDNSSLBSubDomainsResponse struct {
	RequestId     string
	TotalCount    int64
	PageNumber    int64
	PageSize      int64
	SlbSubDomains DescribeDNSSLBSubDomainsSlbSubDomainList
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

func (c *AlidnsClient) DescribeDnsProductAttributes(req *DescribeDnsProductAttributesArgs) (resp *DescribeDnsProductAttributesResponse, err error) {
	resp = &DescribeDnsProductAttributesResponse{}
	err = c.Request("DescribeDnsProductAttributes", req, resp)
	return
}

type DescribeDnsProductAttributesRecordLine struct {
	LineCode   string
	FatherCode string
	LineName   string
}
type DescribeDnsProductAttributesArgs struct {
	VersionCode string
}
type DescribeDnsProductAttributesResponse struct {
	RequestId   string
	TtlMinValue string
	TtlMaxValue string
	RecordLines DescribeDnsProductAttributesRecordLineList
	RecordTypes DescribeDnsProductAttributesRecordTypeList
}

type DescribeDnsProductAttributesRecordLineList []DescribeDnsProductAttributesRecordLine

func (list *DescribeDnsProductAttributesRecordLineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDnsProductAttributesRecordLine)
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

type DescribeDnsProductAttributesRecordTypeList []string

func (list *DescribeDnsProductAttributesRecordTypeList) UnmarshalJSON(data []byte) error {
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

func (c *AlidnsClient) AddDomainGroup(req *AddDomainGroupArgs) (resp *AddDomainGroupResponse, err error) {
	resp = &AddDomainGroupResponse{}
	err = c.Request("AddDomainGroup", req, resp)
	return
}

type AddDomainGroupArgs struct {
	Lang         string
	UserClientIp string
	GroupName    string
}
type AddDomainGroupResponse struct {
	RequestId string
	GroupId   string
	GroupName string
}

func (c *AlidnsClient) ChangeDomainOfDnsProduct(req *ChangeDomainOfDnsProductArgs) (resp *ChangeDomainOfDnsProductResponse, err error) {
	resp = &ChangeDomainOfDnsProductResponse{}
	err = c.Request("ChangeDomainOfDnsProduct", req, resp)
	return
}

type ChangeDomainOfDnsProductArgs struct {
	Lang         string
	UserClientIp string
	InstanceId   string
	NewDomain    string
}
type ChangeDomainOfDnsProductResponse struct {
	RequestId      string
	OriginalDomain string
}

func (c *AlidnsClient) DescribeDomainLogs(req *DescribeDomainLogsArgs) (resp *DescribeDomainLogsResponse, err error) {
	resp = &DescribeDomainLogsResponse{}
	err = c.Request("DescribeDomainLogs", req, resp)
	return
}

type DescribeDomainLogsDomainLog struct {
	ActionTime string
	DomainName string
	Action     string
	Message    string
	ClientIp   string
}
type DescribeDomainLogsArgs struct {
	Lang         string
	UserClientIp string
	KeyWord      string
	GroupId      string
	PageNumber   int64
	PageSize     int64
}
type DescribeDomainLogsResponse struct {
	RequestId  string
	TotalCount int64
	PageNumber int64
	PageSize   int64
	DomainLogs DescribeDomainLogsDomainLogList
}

type DescribeDomainLogsDomainLogList []DescribeDomainLogsDomainLog

func (list *DescribeDomainLogsDomainLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainLogsDomainLog)
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

func (c *AlidnsClient) SetDNSSLBStatus(req *SetDNSSLBStatusArgs) (resp *SetDNSSLBStatusResponse, err error) {
	resp = &SetDNSSLBStatusResponse{}
	err = c.Request("SetDNSSLBStatus", req, resp)
	return
}

type SetDNSSLBStatusArgs struct {
	Lang         string
	UserClientIp string
	SubDomain    string
	Open         bool
}
type SetDNSSLBStatusResponse struct {
	RequestId   string
	RecordCount int64
	Open        bool
}
