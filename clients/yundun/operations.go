package yundun

import "encoding/json"

func (c *YundunClient) DetectVulById(req *DetectVulByIdArgs) (resp *DetectVulByIdResponse, err error) {
	resp = &DetectVulByIdResponse{}
	err = c.Request("DetectVulById", req, resp)
	return
}

type DetectVulByIdArgs struct {
	InstanceId string
	VulId      int
}
type DetectVulByIdResponse struct {
	RequestId string
}

func (c *YundunClient) ConfirmLogin(req *ConfirmLoginArgs) (resp *ConfirmLoginResponse, err error) {
	resp = &ConfirmLoginResponse{}
	err = c.Request("ConfirmLogin", req, resp)
	return
}

type ConfirmLoginArgs struct {
	InstanceId string
	SourceIp   string
	Time       string
}
type ConfirmLoginResponse struct {
	RequestId string
}

func (c *YundunClient) AddCNameWaf(req *AddCNameWafArgs) (resp *AddCNameWafResponse, err error) {
	resp = &AddCNameWafResponse{}
	err = c.Request("AddCNameWaf", req, resp)
	return
}

type AddCNameWafWafInfo struct {
	Id     int
	Domain string
	Cname  string
	Status int
}
type AddCNameWafArgs struct {
	InstanceId   string
	InstanceType string
	Domain       string
}
type AddCNameWafResponse struct {
	RequestId   string
	WafInfoList AddCNameWafWafInfoList
}

type AddCNameWafWafInfoList []AddCNameWafWafInfo

func (list *AddCNameWafWafInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddCNameWafWafInfo)
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

func (c *YundunClient) SecureCheck(req *SecureCheckArgs) (resp *SecureCheckResponse, err error) {
	resp = &SecureCheckResponse{}
	err = c.Request("SecureCheck", req, resp)
	return
}

type SecureCheckInfo struct {
	Ip         string
	Status     string
	VulNum     string
	InstanceId string
}
type SecureCheckArgs struct {
	JstOwnerId  int64
	InstanceIds string
}
type SecureCheckResponse struct {
	RequestId        string
	RecentInstanceId string
	ProblemList      SecureCheckInfoList
	NoProblemList    SecureCheckInfoList
	NoScanList       SecureCheckInfoList
	ScanningList     SecureCheckInfoList
	InnerIpList      SecureCheckInfoList
}

type SecureCheckInfoList []SecureCheckInfo

func (list *SecureCheckInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SecureCheckInfo)
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

func (c *YundunClient) QueryDdosConfig(req *QueryDdosConfigArgs) (resp *QueryDdosConfigResponse, err error) {
	resp = &QueryDdosConfigResponse{}
	err = c.Request("QueryDdosConfig", req, resp)
	return
}

type QueryDdosConfigArgs struct {
	InstanceId   string
	InstanceType string
}
type QueryDdosConfigResponse struct {
	RequestId        string
	Bps              int64
	Pps              int64
	Qps              int64
	Sipconn          int64
	Sipnew           int64
	Layer7Config     bool
	FlowPosition     int
	QpsPosition      int
	StrategyPosition int
	Level            int
	HoleBps          string
	ConfigType       string
}

func (c *YundunClient) CloseCCProtect(req *CloseCCProtectArgs) (resp *CloseCCProtectResponse, err error) {
	resp = &CloseCCProtectResponse{}
	err = c.Request("CloseCCProtect", req, resp)
	return
}

type CloseCCProtectArgs struct {
	InstanceId   string
	InstanceType string
}
type CloseCCProtectResponse struct {
	RequestId string
}

func (c *YundunClient) WebshellLog(req *WebshellLogArgs) (resp *WebshellLogResponse, err error) {
	resp = &WebshellLogResponse{}
	err = c.Request("WebshellLog", req, resp)
	return
}

type WebshellLogWebshellLog struct {
	Id     string
	Path   string
	Status int
	Time   string
}
type WebshellLogArgs struct {
	JstOwnerId int64
	PageNumber int
	PageSize   int
	InstanceId string
	RecordType int
}
type WebshellLogResponse struct {
	RequestId  string
	PageNumber int
	PageSize   int
	TotalCount int
	LogList    WebshellLogWebshellLogList
}

type WebshellLogWebshellLogList []WebshellLogWebshellLog

func (list *WebshellLogWebshellLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]WebshellLogWebshellLog)
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

func (c *YundunClient) DeleteCNameWaf(req *DeleteCNameWafArgs) (resp *DeleteCNameWafResponse, err error) {
	resp = &DeleteCNameWafResponse{}
	err = c.Request("DeleteCNameWaf", req, resp)
	return
}

type DeleteCNameWafWafInfo struct {
	Id     int
	Domain string
	Cname  string
	Status int
}
type DeleteCNameWafArgs struct {
	InstanceId   string
	Domain       string
	CnameId      int
	InstanceType string
}
type DeleteCNameWafResponse struct {
	RequestId   string
	WafInfoList DeleteCNameWafWafInfoList
}

type DeleteCNameWafWafInfoList []DeleteCNameWafWafInfo

func (list *DeleteCNameWafWafInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteCNameWafWafInfo)
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

func (c *YundunClient) BruteforceLog(req *BruteforceLogArgs) (resp *BruteforceLogResponse, err error) {
	resp = &BruteforceLogResponse{}
	err = c.Request("BruteforceLog", req, resp)
	return
}

type BruteforceLogBruteforceLog struct {
	BlockTimes int
	SourceIp   string
	Status     int
	Time       string
	Location   string
}
type BruteforceLogArgs struct {
	JstOwnerId int64
	PageNumber int
	PageSize   int
	InstanceId string
	RecordType int
}
type BruteforceLogResponse struct {
	RequestId  string
	StartTime  string
	EndTime    string
	PageNumber int
	PageSize   int
	TotalCount int
	LogList    BruteforceLogBruteforceLogList
}

type BruteforceLogBruteforceLogList []BruteforceLogBruteforceLog

func (list *BruteforceLogBruteforceLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]BruteforceLogBruteforceLog)
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

func (c *YundunClient) OpenCCProtect(req *OpenCCProtectArgs) (resp *OpenCCProtectResponse, err error) {
	resp = &OpenCCProtectResponse{}
	err = c.Request("OpenCCProtect", req, resp)
	return
}

type OpenCCProtectArgs struct {
	InstanceId   string
	InstanceType string
}
type OpenCCProtectResponse struct {
	RequestId string
}

func (c *YundunClient) DdosLog(req *DdosLogArgs) (resp *DdosLogResponse, err error) {
	resp = &DdosLogResponse{}
	err = c.Request("DdosLog", req, resp)
	return
}

type DdosLogDdosLog struct {
	StartTime    string
	EndTime      string
	Reason       string
	Status       int
	Bps          int64
	Pps          int64
	Qps          int64
	AttackType   string
	AttackIpList string
	Type         int
}
type DdosLogArgs struct {
	JstOwnerId   int64
	PageNumber   int
	PageSize     int
	InstanceId   string
	InstanceType string
}
type DdosLogResponse struct {
	RequestId    string
	AttackStatus int
	StartTime    string
	EndTime      string
	PageNumber   int
	PageSize     int
	TotalCount   int
	LogList      DdosLogDdosLogList
}

type DdosLogDdosLogList []DdosLogDdosLog

func (list *DdosLogDdosLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DdosLogDdosLog)
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

func (c *YundunClient) ServiceStatus(req *ServiceStatusArgs) (resp *ServiceStatusResponse, err error) {
	resp = &ServiceStatusResponse{}
	err = c.Request("ServiceStatus", req, resp)
	return
}

type ServiceStatusArgs struct {
	InstanceId string
}
type ServiceStatusResponse struct {
	RequestId string
	PortScan  bool
	VulScan   bool
}

func (c *YundunClient) SetDdosQps(req *SetDdosQpsArgs) (resp *SetDdosQpsResponse, err error) {
	resp = &SetDdosQpsResponse{}
	err = c.Request("SetDdosQps", req, resp)
	return
}

type SetDdosQpsArgs struct {
	InstanceId   string
	InstanceType string
	QpsPosition  int
	Level        int
}
type SetDdosQpsResponse struct {
	RequestId string
}

func (c *YundunClient) VulScanLog(req *VulScanLogArgs) (resp *VulScanLogResponse, err error) {
	resp = &VulScanLogResponse{}
	err = c.Request("VulScanLog", req, resp)
	return
}

type VulScanLogVulScanLog struct {
	Id           int
	Type         string
	Url          string
	HelpAddress  string
	VulParameter string
	Status       int
}
type VulScanLogArgs struct {
	JstOwnerId int64
	PageNumber int
	PageSize   int
	InstanceId string
	VulStatus  int
}
type VulScanLogResponse struct {
	RequestId  string
	StartTime  string
	EndTime    string
	PageNumber int
	PageSize   int
	TotalCount int
	LogList    VulScanLogVulScanLogList
}

type VulScanLogVulScanLogList []VulScanLogVulScanLog

func (list *VulScanLogVulScanLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]VulScanLogVulScanLog)
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

func (c *YundunClient) WafInfo(req *WafInfoArgs) (resp *WafInfoResponse, err error) {
	resp = &WafInfoResponse{}
	err = c.Request("WafInfo", req, resp)
	return
}

type WafInfoWafInfo struct {
	Id     int
	Domain string
	Cname  string
	Status int
}
type WafInfoArgs struct {
	InstanceId   string
	InstanceType string
}
type WafInfoResponse struct {
	RequestId    string
	WafDomainNum int
	WafInfos     WafInfoWafInfoList
}

type WafInfoWafInfoList []WafInfoWafInfo

func (list *WafInfoWafInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]WafInfoWafInfo)
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

func (c *YundunClient) CloseVulScan(req *CloseVulScanArgs) (resp *CloseVulScanResponse, err error) {
	resp = &CloseVulScanResponse{}
	err = c.Request("CloseVulScan", req, resp)
	return
}

type CloseVulScanArgs struct {
	InstanceId string
}
type CloseVulScanResponse struct {
	RequestId string
}

func (c *YundunClient) DdosFlowGraph(req *DdosFlowGraphArgs) (resp *DdosFlowGraphResponse, err error) {
	resp = &DdosFlowGraphResponse{}
	err = c.Request("DdosFlowGraph", req, resp)
	return
}

type DdosFlowGraphNormalFlow struct {
	Time    int64
	BitRecv int64
	BitSend int64
	PktRecv int64
	PktSend int64
}

type DdosFlowGraphTotalFlow struct {
	Time    int64
	BitRecv int64
	PktRecv int64
}
type DdosFlowGraphArgs struct {
	InstanceId   string
	InstanceType string
}
type DdosFlowGraphResponse struct {
	RequestId   string
	NormalFlows DdosFlowGraphNormalFlowList
	TotalFlows  DdosFlowGraphTotalFlowList
}

type DdosFlowGraphNormalFlowList []DdosFlowGraphNormalFlow

func (list *DdosFlowGraphNormalFlowList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DdosFlowGraphNormalFlow)
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

type DdosFlowGraphTotalFlowList []DdosFlowGraphTotalFlow

func (list *DdosFlowGraphTotalFlowList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DdosFlowGraphTotalFlow)
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

func (c *YundunClient) LogineventLog(req *LogineventLogArgs) (resp *LogineventLogResponse, err error) {
	resp = &LogineventLogResponse{}
	err = c.Request("LogineventLog", req, resp)
	return
}

type LogineventLogLoginEventLog struct {
	BlockTimes int
	SourceIp   string
	Status     int
	Time       string
	Location   string
}
type LogineventLogArgs struct {
	JstOwnerId int64
	PageNumber int
	PageSize   int
	InstanceId string
	RecordType int
}
type LogineventLogResponse struct {
	RequestId  string
	StartTime  string
	EndTime    string
	PageNumber int
	PageSize   int
	TotalCount int
	LogList    LogineventLogLoginEventLogList
}

type LogineventLogLoginEventLogList []LogineventLogLoginEventLog

func (list *LogineventLogLoginEventLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]LogineventLogLoginEventLog)
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

func (c *YundunClient) OpenVulScan(req *OpenVulScanArgs) (resp *OpenVulScanResponse, err error) {
	resp = &OpenVulScanResponse{}
	err = c.Request("OpenVulScan", req, resp)
	return
}

type OpenVulScanArgs struct {
	InstanceId string
}
type OpenVulScanResponse struct {
	RequestId string
}

func (c *YundunClient) GetDdosConfigOptions(req *GetDdosConfigOptionsArgs) (resp *GetDdosConfigOptionsResponse, err error) {
	resp = &GetDdosConfigOptionsResponse{}
	err = c.Request("GetDdosConfigOptions", req, resp)
	return
}

type GetDdosConfigOptionsRequestThresholdOption struct {
	Bps int64
	Pps int64
}

type GetDdosConfigOptionsConnectionThresholdOption struct {
	Sipconn int64
	Sipnew  int64
}
type GetDdosConfigOptionsArgs struct {
}
type GetDdosConfigOptionsResponse struct {
	RequestId                  string
	RequestThresholdOptions1   GetDdosConfigOptionsRequestThresholdOptionList
	RequestThresholdOptions2   GetDdosConfigOptionsRequestThresholdOptionList
	ConnectionThresholdOptions GetDdosConfigOptionsConnectionThresholdOptionList
	QpsOptions1                GetDdosConfigOptionsQpsOptions1List
	QpsOptions2                GetDdosConfigOptionsQpsOptions2List
}

type GetDdosConfigOptionsRequestThresholdOptionList []GetDdosConfigOptionsRequestThresholdOption

func (list *GetDdosConfigOptionsRequestThresholdOptionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetDdosConfigOptionsRequestThresholdOption)
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

type GetDdosConfigOptionsConnectionThresholdOptionList []GetDdosConfigOptionsConnectionThresholdOption

func (list *GetDdosConfigOptionsConnectionThresholdOptionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetDdosConfigOptionsConnectionThresholdOption)
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

type GetDdosConfigOptionsQpsOptions1List []string

func (list *GetDdosConfigOptionsQpsOptions1List) UnmarshalJSON(data []byte) error {
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

type GetDdosConfigOptionsQpsOptions2List []string

func (list *GetDdosConfigOptionsQpsOptions2List) UnmarshalJSON(data []byte) error {
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

func (c *YundunClient) Summary(req *SummaryArgs) (resp *SummaryResponse, err error) {
	resp = &SummaryResponse{}
	err = c.Request("Summary", req, resp)
	return
}

type SummaryDdos struct {
	Count     int64
	HostCount int64
}

type SummaryBruteForce struct {
	Count     int64
	HostCount int64
}

type SummaryWebshell struct {
	Count     int64
	HostCount int64
}

type SummaryRemoteLogin struct {
	Count     int64
	HostCount int64
}

type SummaryWebAttack struct {
	Count     int64
	HostCount int64
}

type SummaryWebLeak struct {
	Count     int64
	HostCount int64
}
type SummaryArgs struct {
	JstOwnerId  int64
	InstanceIds string
}
type SummaryResponse struct {
	RequestId         string
	Status            int64
	AbnormalHostCount int64
	Ddos              SummaryDdos
	BruteForce        SummaryBruteForce
	Webshell          SummaryWebshell
	RemoteLogin       SummaryRemoteLogin
	WebAttack         SummaryWebAttack
	WebLeak           SummaryWebLeak
}

func (c *YundunClient) ClosePortScan(req *ClosePortScanArgs) (resp *ClosePortScanResponse, err error) {
	resp = &ClosePortScanResponse{}
	err = c.Request("ClosePortScan", req, resp)
	return
}

type ClosePortScanArgs struct {
	InstanceId string
}
type ClosePortScanResponse struct {
	RequestId string
}

func (c *YundunClient) ListInstanceInfos(req *ListInstanceInfosArgs) (resp *ListInstanceInfosResponse, err error) {
	resp = &ListInstanceInfosResponse{}
	err = c.Request("ListInstanceInfos", req, resp)
	return
}

type ListInstanceInfosInstanceInfo struct {
	Region       string
	RegionName   string
	RegionEnName string
	InstanceName string
	InstanceId   string
	Ip           string
	InternetIp   string
	IntranetIp   string
	Ddos         int
	HostEvent    int
	SecureCheck  int
	AegisStatus  int
	Waf          int
	IsLock       bool
	LockType     string
	UnLockTimes  int
	TriggerTime  string
}
type ListInstanceInfosArgs struct {
	JstOwnerId   int64
	PageNumber   int
	PageSize     int
	Region       string
	EventType    string
	InstanceName string
	InstanceType string
	InstanceIds  string
}
type ListInstanceInfosResponse struct {
	RequestId  string
	PageNumber int
	PageSize   int
	TotalCount int
	InfosList  ListInstanceInfosInstanceInfoList
}

type ListInstanceInfosInstanceInfoList []ListInstanceInfosInstanceInfo

func (list *ListInstanceInfosInstanceInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListInstanceInfosInstanceInfo)
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

func (c *YundunClient) WafLog(req *WafLogArgs) (resp *WafLogResponse, err error) {
	resp = &WafLogResponse{}
	err = c.Request("WafLog", req, resp)
	return
}

type WafLogWafLog struct {
	SourceIp string
	Time     string
	Url      string
	Type     string
	Status   int
}
type WafLogArgs struct {
	JstOwnerId   int64
	PageNumber   int
	PageSize     int
	InstanceId   string
	InstanceType string
}
type WafLogResponse struct {
	RequestId   string
	WebAttack   int
	NewWafUser  bool
	WafOpened   bool
	InWhiteList bool
	DomainCount int
	StartTime   string
	EndTime     string
	PageNumber  int
	PageSize    int
	TotalCount  int
	LogList     WafLogWafLogList
}

type WafLogWafLogList []WafLogWafLog

func (list *WafLogWafLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]WafLogWafLog)
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

func (c *YundunClient) ConfigDdos(req *ConfigDdosArgs) (resp *ConfigDdosResponse, err error) {
	resp = &ConfigDdosResponse{}
	err = c.Request("ConfigDdos", req, resp)
	return
}

type ConfigDdosArgs struct {
	InstanceId       string
	InstanceType     string
	FlowPosition     int
	StrategyPosition int
	Level            int
}
type ConfigDdosResponse struct {
	RequestId string
}

func (c *YundunClient) SetDdosAuto(req *SetDdosAutoArgs) (resp *SetDdosAutoResponse, err error) {
	resp = &SetDdosAutoResponse{}
	err = c.Request("SetDdosAuto", req, resp)
	return
}

type SetDdosAutoArgs struct {
	InstanceId   string
	InstanceType string
}
type SetDdosAutoResponse struct {
	RequestId string
}

func (c *YundunClient) DeleteBackDoorFile(req *DeleteBackDoorFileArgs) (resp *DeleteBackDoorFileResponse, err error) {
	resp = &DeleteBackDoorFileResponse{}
	err = c.Request("DeleteBackDoorFile", req, resp)
	return
}

type DeleteBackDoorFileArgs struct {
	JstOwnerId int64
	InstanceId string
	Path       string
}
type DeleteBackDoorFileResponse struct {
	RequestId string
}

func (c *YundunClient) DetectVulByIp(req *DetectVulByIpArgs) (resp *DetectVulByIpResponse, err error) {
	resp = &DetectVulByIpResponse{}
	err = c.Request("DetectVulByIp", req, resp)
	return
}

type DetectVulByIpArgs struct {
	InstanceId string
	VulIp      string
}
type DetectVulByIpResponse struct {
	RequestId string
}

func (c *YundunClient) OpenPortScan(req *OpenPortScanArgs) (resp *OpenPortScanResponse, err error) {
	resp = &OpenPortScanResponse{}
	err = c.Request("OpenPortScan", req, resp)
	return
}

type OpenPortScanArgs struct {
	InstanceId string
}
type OpenPortScanResponse struct {
	RequestId string
}
