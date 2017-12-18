package cms

import "encoding/json"

func (c *CmsClient) ListAlarm(req *ListAlarmArgs) (resp *ListAlarmResponse, err error) {
	resp = &ListAlarmResponse{}
	err = c.Request("ListAlarm", req, resp)
	return
}

type ListAlarmAlarm struct {
	Id                 string
	Name               string
	Namespace          string
	MetricName         string
	Dimensions         string
	Period             int
	Statistics         string
	ComparisonOperator string
	Threshold          string
	EvaluationCount    int
	StartTime          int
	EndTime            int
	SilenceTime        int
	NotifyType         int
	Enable             bool
	State              string
	ContactGroups      string
	Webhook            string
}
type ListAlarmArgs struct {
	IsEnable         bool
	Callby_cms_owner string
	Name             string
	Namespace        string
	PageSize         int
	Id               string
	State            string
	Dimension        string
	PageNumber       int
}
type ListAlarmResponse struct {
	Success   bool
	Code      string
	Message   string
	NextToken int
	Total     int
	RequestId string
	AlarmList ListAlarmAlarmList
}

type ListAlarmAlarmList []ListAlarmAlarm

func (list *ListAlarmAlarmList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAlarmAlarm)
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

func (c *CmsClient) DeleteCustomMetric(req *DeleteCustomMetricArgs) (resp *DeleteCustomMetricResponse, err error) {
	resp = &DeleteCustomMetricResponse{}
	err = c.Request("DeleteCustomMetric", req, resp)
	return
}

type DeleteCustomMetricArgs struct {
	GroupId    string
	MetricName string
	UUID       string
	Md5        string
}
type DeleteCustomMetricResponse struct {
	Code      string
	Message   string
	RequestId string
	Result    string
}

func (c *CmsClient) ListContactGroup(req *ListContactGroupArgs) (resp *ListContactGroupResponse, err error) {
	resp = &ListContactGroupResponse{}
	err = c.Request("ListContactGroup", req, resp)
	return
}

type ListContactGroupArgs struct {
	Callby_cms_owner string
	PageSize         int
	PageNumber       int
}
type ListContactGroupResponse struct {
	Success       bool
	Code          string
	Message       string
	NextToken     int
	Total         int
	RequestId     string
	ContactGroups ListContactGroupContactGroupList
}

type ListContactGroupContactGroupList []string

func (list *ListContactGroupContactGroupList) UnmarshalJSON(data []byte) error {
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

func (c *CmsClient) PutEvent(req *PutEventArgs) (resp *PutEventResponse, err error) {
	resp = &PutEventResponse{}
	err = c.Request("PutEvent", req, resp)
	return
}

type PutEventArgs struct {
	EventInfo string
}
type PutEventResponse struct {
	Code    string
	Message string
	Data    string
}

func (c *CmsClient) UpdateAlarm(req *UpdateAlarmArgs) (resp *UpdateAlarmResponse, err error) {
	resp = &UpdateAlarmResponse{}
	err = c.Request("UpdateAlarm", req, resp)
	return
}

type UpdateAlarmArgs struct {
	Callby_cms_owner   string
	Period             int
	Webhook            string
	ContactGroups      string
	EndTime            int
	Threshold          string
	StartTime          int
	Name               string
	EvaluationCount    int
	SilenceTime        int
	Id                 string
	NotifyType         int
	ComparisonOperator string
	Statistics         string
}
type UpdateAlarmResponse struct {
	Success   bool
	Code      string
	Message   string
	RequestId string
}

func (c *CmsClient) NodeProcessCreate(req *NodeProcessCreateArgs) (resp *NodeProcessCreateResponse, err error) {
	resp = &NodeProcessCreateResponse{}
	err = c.Request("NodeProcessCreate", req, resp)
	return
}

type NodeProcessCreateArgs struct {
	InstanceId  string
	ProcessName string
	Name        string
	ProcessUser string
	Command     string
}
type NodeProcessCreateResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
}

func (c *CmsClient) ProfileSet(req *ProfileSetArgs) (resp *ProfileSetResponse, err error) {
	resp = &ProfileSetResponse{}
	err = c.Request("ProfileSet", req, resp)
	return
}

type ProfileSetArgs struct {
	AutoInstall bool
	UserId      int64
}
type ProfileSetResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
}

func (c *CmsClient) NodeList(req *NodeListArgs) (resp *NodeListResponse, err error) {
	resp = &NodeListResponse{}
	err = c.Request("NodeList", req, resp)
	return
}

type NodeListNode struct {
	InstanceId       string
	SerialNumber     string
	HostName         string
	AliUid           int64
	OperatingSystem  string
	IpGroup          string
	Region           string
	TianjimonVersion string
	EipAddress       string
	EipId            string
	AliyunHost       bool
	NatIp            string
	NetworkType      string
}
type NodeListArgs struct {
	HostName      string
	InstanceIds   string
	PageSize      int
	KeyWord       string
	UserId        int64
	SerialNumbers string
	PageNumber    int
	Status        string
}
type NodeListResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
	PageNumber   int
	PageSize     int
	PageTotal    int
	Total        int
	Nodes        NodeListNodeList
}

type NodeListNodeList []NodeListNode

func (list *NodeListNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]NodeListNode)
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

func (c *CmsClient) NodeProcessDelete(req *NodeProcessDeleteArgs) (resp *NodeProcessDeleteResponse, err error) {
	resp = &NodeProcessDeleteResponse{}
	err = c.Request("NodeProcessDelete", req, resp)
	return
}

type NodeProcessDeleteArgs struct {
	InstanceId string
	Name       string
	Id         string
}
type NodeProcessDeleteResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
}

func (c *CmsClient) QueryCustomMetricList(req *QueryCustomMetricListArgs) (resp *QueryCustomMetricListResponse, err error) {
	resp = &QueryCustomMetricListResponse{}
	err = c.Request("QueryCustomMetricList", req, resp)
	return
}

type QueryCustomMetricListArgs struct {
	Size       string
	GroupId    string
	Page       string
	MetricName string
	Dimension  string
	UUID       string
	Md5        string
}
type QueryCustomMetricListResponse struct {
	Code      string
	Message   string
	RequestId string
	Result    string
}

func (c *CmsClient) NodeUninstall(req *NodeUninstallArgs) (resp *NodeUninstallResponse, err error) {
	resp = &NodeUninstallResponse{}
	err = c.Request("NodeUninstall", req, resp)
	return
}

type NodeUninstallArgs struct {
	InstanceId string
}
type NodeUninstallResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
}

func (c *CmsClient) ListAlarmHistory(req *ListAlarmHistoryArgs) (resp *ListAlarmHistoryResponse, err error) {
	resp = &ListAlarmHistoryResponse{}
	err = c.Request("ListAlarmHistory", req, resp)
	return
}

type ListAlarmHistoryAlarmHistory struct {
	Id              string
	Name            string
	Namespace       string
	MetricName      string
	Dimension       string
	EvaluationCount int
	Value           string
	AlarmTime       int64
	LastTime        int64
	State           string
	Status          int
	ContactGroups   string
}
type ListAlarmHistoryArgs struct {
	Cursor           string
	Callby_cms_owner string
	Size             int
	EndTime          string
	Id               string
	StartTime        string
}
type ListAlarmHistoryResponse struct {
	Success          bool
	Code             string
	Message          string
	Cursor           string
	RequestId        string
	AlarmHistoryList ListAlarmHistoryAlarmHistoryList
}

type ListAlarmHistoryAlarmHistoryList []ListAlarmHistoryAlarmHistory

func (list *ListAlarmHistoryAlarmHistoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAlarmHistoryAlarmHistory)
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

func (c *CmsClient) CreateAlarm(req *CreateAlarmArgs) (resp *CreateAlarmResponse, err error) {
	resp = &CreateAlarmResponse{}
	err = c.Request("CreateAlarm", req, resp)
	return
}

type CreateAlarmArgs struct {
	Callby_cms_owner   string
	Period             int
	Webhook            string
	ContactGroups      string
	EndTime            int
	Threshold          string
	StartTime          int
	Name               string
	Namespace          string
	EvaluationCount    int
	SilenceTime        int
	MetricName         string
	NotifyType         int
	ComparisonOperator string
	Dimensions         string
	Statistics         string
}
type CreateAlarmResponse struct {
	Success   bool
	Code      string
	Message   string
	RequestId string
	Data      string
}

func (c *CmsClient) QueryMetricLast(req *QueryMetricLastArgs) (resp *QueryMetricLastResponse, err error) {
	resp = &QueryMetricLastResponse{}
	err = c.Request("QueryMetricLast", req, resp)
	return
}

type QueryMetricLastArgs struct {
	Cursor           string
	Callby_cms_owner string
	ResourceOwnerId  int64
	Period           string
	Length           string
	Project          string
	EndTime          string
	Express          string
	StartTime        string
	Metric           string
	Page             string
	Dimensions       string
}
type QueryMetricLastResponse struct {
	Code       string
	Message    string
	Success    bool
	RequestId  string
	Cursor     string
	Datapoints string
	Period     string
}

func (c *CmsClient) DescribeAlarmHistory(req *DescribeAlarmHistoryArgs) (resp *DescribeAlarmHistoryResponse, err error) {
	resp = &DescribeAlarmHistoryResponse{}
	err = c.Request("DescribeAlarmHistory", req, resp)
	return
}

type DescribeAlarmHistoryAlarmHistory struct {
	Id              string
	AlertName       string
	GroupId         string
	Namespace       string
	MetricName      string
	Dimensions      string
	Expression      string
	EvaluationCount int
	Value           string
	AlertTime       int64
	LastTime        int64
	Level           string
	PreLevel        string
	RuleName        string
	State           string
	Status          int
	UserId          string
	Webhooks        string
	ContactGroups   DescribeAlarmHistoryContactGroupList
	Contacts        DescribeAlarmHistoryContactList
	ContactALIIMs   DescribeAlarmHistoryContactALIIMList
	ContactSmses    DescribeAlarmHistoryContactSmseList
	ContactMails    DescribeAlarmHistoryContactMailList
}
type DescribeAlarmHistoryArgs struct {
	AlertName  string
	GroupId    string
	EndTime    string
	RuleName   string
	StartTime  string
	Ascending  bool
	OnlyCount  bool
	Namespace  string
	PageSize   int
	State      string
	Page       int
	MetricName string
	Status     string
}
type DescribeAlarmHistoryResponse struct {
	Success          bool
	Code             string
	Message          string
	Total            string
	RequestId        string
	AlarmHistoryList DescribeAlarmHistoryAlarmHistoryList
}

type DescribeAlarmHistoryContactGroupList []string

func (list *DescribeAlarmHistoryContactGroupList) UnmarshalJSON(data []byte) error {
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

type DescribeAlarmHistoryContactList []string

func (list *DescribeAlarmHistoryContactList) UnmarshalJSON(data []byte) error {
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

type DescribeAlarmHistoryContactALIIMList []string

func (list *DescribeAlarmHistoryContactALIIMList) UnmarshalJSON(data []byte) error {
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

type DescribeAlarmHistoryContactSmseList []string

func (list *DescribeAlarmHistoryContactSmseList) UnmarshalJSON(data []byte) error {
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

type DescribeAlarmHistoryContactMailList []string

func (list *DescribeAlarmHistoryContactMailList) UnmarshalJSON(data []byte) error {
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

type DescribeAlarmHistoryAlarmHistoryList []DescribeAlarmHistoryAlarmHistory

func (list *DescribeAlarmHistoryAlarmHistoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAlarmHistoryAlarmHistory)
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

func (c *CmsClient) AccessKeyGet(req *AccessKeyGetArgs) (resp *AccessKeyGetResponse, err error) {
	resp = &AccessKeyGetResponse{}
	err = c.Request("AccessKeyGet", req, resp)
	return
}

type AccessKeyGetArgs struct {
	UserId int64
}
type AccessKeyGetResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
	UserId       int64
	AccessKey    string
	SecretKey    string
}

func (c *CmsClient) PutCustomMetric(req *PutCustomMetricArgs) (resp *PutCustomMetricResponse, err error) {
	resp = &PutCustomMetricResponse{}
	err = c.Request("PutCustomMetric", req, resp)
	return
}

type PutCustomMetricArgs struct {
	MetricList string
}
type PutCustomMetricResponse struct {
	Code    string
	Message string
	Data    string
}

func (c *CmsClient) DeleteAlarm(req *DeleteAlarmArgs) (resp *DeleteAlarmResponse, err error) {
	resp = &DeleteAlarmResponse{}
	err = c.Request("DeleteAlarm", req, resp)
	return
}

type DeleteAlarmArgs struct {
	Callby_cms_owner string
	Id               string
}
type DeleteAlarmResponse struct {
	Success   bool
	Code      string
	Message   string
	RequestId string
}

func (c *CmsClient) NodeInstall(req *NodeInstallArgs) (resp *NodeInstallResponse, err error) {
	resp = &NodeInstallResponse{}
	err = c.Request("NodeInstall", req, resp)
	return
}

type NodeInstallArgs struct {
	InstanceId string
	Force      bool
	UserId     string
}
type NodeInstallResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
}

func (c *CmsClient) PutMetricData(req *PutMetricDataArgs) (resp *PutMetricDataResponse, err error) {
	resp = &PutMetricDataResponse{}
	err = c.Request("PutMetricData", req, resp)
	return
}

type PutMetricDataArgs struct {
	Callby_cms_owner string
	Body             string
}
type PutMetricDataResponse struct {
	Code      string
	Message   string
	RequestId string
	Success   bool
}

func (c *CmsClient) NodeStatus(req *NodeStatusArgs) (resp *NodeStatusResponse, err error) {
	resp = &NodeStatusResponse{}
	err = c.Request("NodeStatus", req, resp)
	return
}

type NodeStatusArgs struct {
	InstanceId string
}
type NodeStatusResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
	InstanceId   string
	AutoInstall  bool
	Status       string
}

func (c *CmsClient) ProfileGet(req *ProfileGetArgs) (resp *ProfileGetResponse, err error) {
	resp = &ProfileGetResponse{}
	err = c.Request("ProfileGet", req, resp)
	return
}

type ProfileGetArgs struct {
	UserId int64
}
type ProfileGetResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
	UserId       int64
	AutoInstall  bool
}

func (c *CmsClient) PutSystemEvent(req *PutSystemEventArgs) (resp *PutSystemEventResponse, err error) {
	resp = &PutSystemEventResponse{}
	err = c.Request("PutSystemEvent", req, resp)
	return
}

type PutSystemEventArgs struct {
	EventInfo string
}
type PutSystemEventResponse struct {
	Code    string
	Message string
	Data    string
}

func (c *CmsClient) DisableAlarm(req *DisableAlarmArgs) (resp *DisableAlarmResponse, err error) {
	resp = &DisableAlarmResponse{}
	err = c.Request("DisableAlarm", req, resp)
	return
}

type DisableAlarmArgs struct {
	Callby_cms_owner string
	Id               string
}
type DisableAlarmResponse struct {
	Success   bool
	Code      string
	Message   string
	RequestId string
}

func (c *CmsClient) NodeStatusList(req *NodeStatusListArgs) (resp *NodeStatusListResponse, err error) {
	resp = &NodeStatusListResponse{}
	err = c.Request("NodeStatusList", req, resp)
	return
}

type NodeStatusListNodeStatus struct {
	InstanceId  string
	AutoInstall bool
	Status      string
}
type NodeStatusListArgs struct {
	InstanceIds string
}
type NodeStatusListResponse struct {
	ErrorCode      int
	ErrorMessage   string
	Success        bool
	RequestId      string
	NodeStatusList NodeStatusListNodeStatusList
}

type NodeStatusListNodeStatusList []NodeStatusListNodeStatus

func (list *NodeStatusListNodeStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]NodeStatusListNodeStatus)
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

func (c *CmsClient) QueryMetricList(req *QueryMetricListArgs) (resp *QueryMetricListResponse, err error) {
	resp = &QueryMetricListResponse{}
	err = c.Request("QueryMetricList", req, resp)
	return
}

type QueryMetricListArgs struct {
	Cursor           string
	Callby_cms_owner string
	ResourceOwnerId  int64
	Period           string
	Length           string
	Project          string
	EndTime          string
	Express          string
	StartTime        string
	Metric           string
	Page             string
	Dimensions       string
}
type QueryMetricListResponse struct {
	Code       string
	Message    string
	Success    bool
	RequestId  string
	Cursor     string
	Datapoints string
	Period     string
}

func (c *CmsClient) EnableAlarm(req *EnableAlarmArgs) (resp *EnableAlarmResponse, err error) {
	resp = &EnableAlarmResponse{}
	err = c.Request("EnableAlarm", req, resp)
	return
}

type EnableAlarmArgs struct {
	Callby_cms_owner string
	Id               string
}
type EnableAlarmResponse struct {
	Success   bool
	Code      string
	Message   string
	RequestId string
}

func (c *CmsClient) NodeProcesses(req *NodeProcessesArgs) (resp *NodeProcessesResponse, err error) {
	resp = &NodeProcessesResponse{}
	err = c.Request("NodeProcesses", req, resp)
	return
}

type NodeProcessesNodeProcess struct {
	Id          int64
	Name        string
	InstanceId  string
	ProcessName string
	ProcessUser string
	Command     string
}
type NodeProcessesArgs struct {
	InstanceId string
}
type NodeProcessesResponse struct {
	ErrorCode     int
	ErrorMessage  string
	Success       bool
	RequestId     string
	NodeProcesses NodeProcessesNodeProcessList
}

type NodeProcessesNodeProcessList []NodeProcessesNodeProcess

func (list *NodeProcessesNodeProcessList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]NodeProcessesNodeProcess)
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
