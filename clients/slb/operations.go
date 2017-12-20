package slb

import (
	"encoding/json"

	"github.com/morlay/aliyun-go/core"
)

func (c *SlbClient) SetLoadBalancerUDPListenerAttribute(req *SetLoadBalancerUDPListenerAttributeArgs) (resp *SetLoadBalancerUDPListenerAttributeResponse, err error) {
	resp = &SetLoadBalancerUDPListenerAttributeResponse{}
	err = c.Request("SetLoadBalancerUDPListenerAttribute", req, resp)
	return
}

type SetLoadBalancerUDPListenerAttributeArgs struct {
	Access_key_id             string
	HealthCheckConnectTimeout int
	ResourceOwnerId           int64
	UnhealthyThreshold        int
	HealthyThreshold          int
	Scheduler                 string
	MasterSlaveServerGroup    string
	MaxConnection             int
	PersistenceTimeout        int
	VServerGroupId            string
	ListenerPort              int
	ResourceOwnerAccount      string
	Bandwidth                 int
	OwnerAccount              string
	OwnerId                   int64
	Tags                      string
	LoadBalancerId            string
	MasterSlaveServerGroupId  string
	HealthCheckReq            string
	HealthCheckInterval       int
	HealthCheckExp            string
	HealthCheckConnectPort    int
	VServerGroup              string
}
type SetLoadBalancerUDPListenerAttributeResponse struct {
	RequestId string
}

func (c *SlbClient) AddListenerWhiteListItem(req *AddListenerWhiteListItemArgs) (resp *AddListenerWhiteListItemResponse, err error) {
	resp = &AddListenerWhiteListItemResponse{}
	err = c.Request("AddListenerWhiteListItem", req, resp)
	return
}

type AddListenerWhiteListItemArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ListenerPort         int
	LoadBalancerId       string
	SourceItems          string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type AddListenerWhiteListItemResponse struct {
	RequestId string
}

func (c *SlbClient) DeleteCACertificate(req *DeleteCACertificateArgs) (resp *DeleteCACertificateResponse, err error) {
	resp = &DeleteCACertificateResponse{}
	err = c.Request("DeleteCACertificate", req, resp)
	return
}

type DeleteCACertificateArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	CACertificateId      string
}
type DeleteCACertificateResponse struct {
	RequestId string
}

func (c *SlbClient) SetCACertificateName(req *SetCACertificateNameArgs) (resp *SetCACertificateNameResponse, err error) {
	resp = &SetCACertificateNameResponse{}
	err = c.Request("SetCACertificateName", req, resp)
	return
}

type SetCACertificateNameArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	CACertificateName    string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	CACertificateId      string
}
type SetCACertificateNameResponse struct {
	RequestId string
}

func (c *SlbClient) DescribeTags(req *DescribeTagsArgs) (resp *DescribeTagsResponse, err error) {
	resp = &DescribeTagsResponse{}
	err = c.Request("DescribeTags", req, resp)
	return
}

type DescribeTagsTagSet struct {
	TagKey        string
	TagValue      string
	InstanceCount int
}
type DescribeTagsArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DistinctKey          core.Bool
	OwnerId              int64
	PageNumber           int
	Tags                 string
	LoadBalancerId       string
	PageSize             int
}
type DescribeTagsResponse struct {
	RequestId  string
	PageSize   int
	PageNumber int
	TotalCount int
	TagSets    DescribeTagsTagSetList
}

type DescribeTagsTagSetList []DescribeTagsTagSet

func (list *DescribeTagsTagSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTagsTagSet)
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

func (c *SlbClient) DeleteRules(req *DeleteRulesArgs) (resp *DeleteRulesResponse, err error) {
	resp = &DeleteRulesResponse{}
	err = c.Request("DeleteRules", req, resp)
	return
}

type DeleteRulesArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	RuleIds              string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DeleteRulesResponse struct {
	RequestId string
}

func (c *SlbClient) DescribeLoadBalancers(req *DescribeLoadBalancersArgs) (resp *DescribeLoadBalancersResponse, err error) {
	resp = &DescribeLoadBalancersResponse{}
	err = c.Request("DescribeLoadBalancers", req, resp)
	return
}

type DescribeLoadBalancersLoadBalancer struct {
	LoadBalancerId     string
	LoadBalancerName   string
	LoadBalancerStatus string
	Address            string
	AddressType        string
	RegionId           string
	RegionIdAlias      string
	VSwitchId          string
	VpcId              string
	NetworkType        string
	MasterZoneId       string
	SlaveZoneId        string
	InternetChargeType string
	CreateTime         string
	CreateTimeStamp    int64
	PayType            string
	ResourceGroupId    string
}
type DescribeLoadBalancersArgs struct {
	Access_key_id         string
	ResourceOwnerId       int64
	NetworkType           string
	MasterZoneId          string
	PageNumber            int
	ResourceGroupId       string
	LoadBalancerName      string
	PageSize              int
	AddressType           string
	SlaveZoneId           string
	Address               string
	ResourceOwnerAccount  string
	OwnerAccount          string
	OwnerId               int64
	ServerId              string
	Tags                  string
	ServerIntranetAddress string
	VSwitchId             string
	LoadBalancerId        string
	InternetChargeType    string
	VpcId                 string
	PayType               string
}
type DescribeLoadBalancersResponse struct {
	RequestId     string
	PageNumber    int
	PageSize      int
	TotalCount    int
	LoadBalancers DescribeLoadBalancersLoadBalancerList
}

type DescribeLoadBalancersLoadBalancerList []DescribeLoadBalancersLoadBalancer

func (list *DescribeLoadBalancersLoadBalancerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersLoadBalancer)
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

func (c *SlbClient) CreateLoadBalancerHTTPListener(req *CreateLoadBalancerHTTPListenerArgs) (resp *CreateLoadBalancerHTTPListenerResponse, err error) {
	resp = &CreateLoadBalancerHTTPListenerResponse{}
	err = c.Request("CreateLoadBalancerHTTPListener", req, resp)
	return
}

type CreateLoadBalancerHTTPListenerArgs struct {
	Access_key_id          string
	ResourceOwnerId        int64
	HealthCheckTimeout     int
	XForwardedFor          string
	HealthCheckURI         string
	UnhealthyThreshold     int
	HealthyThreshold       int
	Scheduler              string
	HealthCheck            string
	MaxConnection          int
	CookieTimeout          int
	StickySessionType      string
	VServerGroupId         string
	ListenerPort           int
	Cookie                 string
	ResourceOwnerAccount   string
	Bandwidth              int
	StickySession          string
	HealthCheckDomain      string
	OwnerAccount           string
	Gzip                   string
	OwnerId                int64
	Tags                   string
	LoadBalancerId         string
	XForwardedFor_SLBIP    string
	BackendServerPort      int
	HealthCheckInterval    int
	XForwardedFor_proto    string
	XForwardedFor_SLBID    string
	HealthCheckConnectPort int
	HealthCheckHttpCode    string
}
type CreateLoadBalancerHTTPListenerResponse struct {
	RequestId string
}

func (c *SlbClient) DescribeLoadBalancerHTTPSListenerAttribute(req *DescribeLoadBalancerHTTPSListenerAttributeArgs) (resp *DescribeLoadBalancerHTTPSListenerAttributeResponse, err error) {
	resp = &DescribeLoadBalancerHTTPSListenerAttributeResponse{}
	err = c.Request("DescribeLoadBalancerHTTPSListenerAttribute", req, resp)
	return
}

type DescribeLoadBalancerHTTPSListenerAttributeArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ListenerPort         int
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DescribeLoadBalancerHTTPSListenerAttributeResponse struct {
	RequestId              string
	ListenerPort           int
	BackendServerPort      int
	Bandwidth              int
	Status                 string
	SecurityStatus         string
	XForwardedFor          string
	Scheduler              string
	StickySession          string
	StickySessionType      string
	CookieTimeout          int
	Cookie                 string
	HealthCheck            string
	HealthCheckDomain      string
	HealthCheckURI         string
	HealthyThreshold       int
	UnhealthyThreshold     int
	HealthCheckTimeout     int
	HealthCheckInterval    int
	HealthCheckConnectPort int
	HealthCheckHttpCode    string
	ServerCertificateId    string
	CACertificateId        string
	MaxConnection          int
	VServerGroupId         string
	Gzip                   string
	XForwardedFor_SLBIP    string
	XForwardedFor_SLBID    string
	XForwardedFor_proto    string
}

func (c *SlbClient) DescribeServerCertificates(req *DescribeServerCertificatesArgs) (resp *DescribeServerCertificatesResponse, err error) {
	resp = &DescribeServerCertificatesResponse{}
	err = c.Request("DescribeServerCertificates", req, resp)
	return
}

type DescribeServerCertificatesServerCertificate struct {
	ServerCertificateId     string
	Fingerprint             string
	ServerCertificateName   string
	RegionId                string
	RegionIdAlias           string
	AliCloudCertificateId   string
	AliCloudCertificateName string
	IsAliCloudCertificate   int
	ResourceGroupId         string
	CreateTime              string
	CreateTimeStamp         int64
}
type DescribeServerCertificatesArgs struct {
	Access_key_id        string
	ResourceGroupId      string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	ServerCertificateId  string
	Tags                 string
}
type DescribeServerCertificatesResponse struct {
	RequestId          string
	ServerCertificates DescribeServerCertificatesServerCertificateList
}

type DescribeServerCertificatesServerCertificateList []DescribeServerCertificatesServerCertificate

func (list *DescribeServerCertificatesServerCertificateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeServerCertificatesServerCertificate)
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

func (c *SlbClient) DescribeVServerGroups(req *DescribeVServerGroupsArgs) (resp *DescribeVServerGroupsResponse, err error) {
	resp = &DescribeVServerGroupsResponse{}
	err = c.Request("DescribeVServerGroups", req, resp)
	return
}

type DescribeVServerGroupsVServerGroup struct {
	VServerGroupId   string
	VServerGroupName string
}
type DescribeVServerGroupsArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DescribeVServerGroupsResponse struct {
	RequestId     string
	VServerGroups DescribeVServerGroupsVServerGroupList
}

type DescribeVServerGroupsVServerGroupList []DescribeVServerGroupsVServerGroup

func (list *DescribeVServerGroupsVServerGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVServerGroupsVServerGroup)
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

func (c *SlbClient) DescribeLoadBalancerAutoReleaseTime(req *DescribeLoadBalancerAutoReleaseTimeArgs) (resp *DescribeLoadBalancerAutoReleaseTimeResponse, err error) {
	resp = &DescribeLoadBalancerAutoReleaseTimeResponse{}
	err = c.Request("DescribeLoadBalancerAutoReleaseTime", req, resp)
	return
}

type DescribeLoadBalancerAutoReleaseTimeArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DescribeLoadBalancerAutoReleaseTimeResponse struct {
	RequestId       string
	AutoReleaseTime int64
}

func (c *SlbClient) SetListenerAccessControlStatus(req *SetListenerAccessControlStatusArgs) (resp *SetListenerAccessControlStatusResponse, err error) {
	resp = &SetListenerAccessControlStatusResponse{}
	err = c.Request("SetListenerAccessControlStatus", req, resp)
	return
}

type SetListenerAccessControlStatusArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ListenerPort         int
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	AccessControlStatus  string
	OwnerId              int64
	Tags                 string
}
type SetListenerAccessControlStatusResponse struct {
	RequestId string
}

func (c *SlbClient) MoveResourceGroup(req *MoveResourceGroupArgs) (resp *MoveResourceGroupResponse, err error) {
	resp = &MoveResourceGroupResponse{}
	err = c.Request("MoveResourceGroup", req, resp)
	return
}

type MoveResourceGroupArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ResourceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	ResourceType         string
	Tags                 string
	ResourceGroupId      string
	NewResourceGroupId   string
}
type MoveResourceGroupResponse struct {
	RequestId string
}

func (c *SlbClient) SetServerCertificateName(req *SetServerCertificateNameArgs) (resp *SetServerCertificateNameResponse, err error) {
	resp = &SetServerCertificateNameResponse{}
	err = c.Request("SetServerCertificateName", req, resp)
	return
}

type SetServerCertificateNameArgs struct {
	Access_key_id         string
	ResourceOwnerId       int64
	ResourceOwnerAccount  string
	OwnerAccount          string
	OwnerId               int64
	ServerCertificateId   string
	ServerCertificateName string
	Tags                  string
}
type SetServerCertificateNameResponse struct {
	RequestId string
}

func (c *SlbClient) DescribeMasterSlaveVServerGroupAttribute(req *DescribeMasterSlaveVServerGroupAttributeArgs) (resp *DescribeMasterSlaveVServerGroupAttributeResponse, err error) {
	resp = &DescribeMasterSlaveVServerGroupAttributeResponse{}
	err = c.Request("DescribeMasterSlaveVServerGroupAttribute", req, resp)
	return
}

type DescribeMasterSlaveVServerGroupAttributeMasterSlaveBackendServer struct {
	ServerId string
	Port     int
	Weight   int
	IsBackup int
}
type DescribeMasterSlaveVServerGroupAttributeArgs struct {
	Access_key_id             string
	ResourceOwnerId           int64
	MasterSlaveVServerGroupId string
	ResourceOwnerAccount      string
	OwnerAccount              string
	OwnerId                   int64
	Tags                      string
}
type DescribeMasterSlaveVServerGroupAttributeResponse struct {
	RequestId                   string
	MasterSlaveVServerGroupId   string
	MasterSlaveVServerGroupName string
	MasterSlaveBackendServers   DescribeMasterSlaveVServerGroupAttributeMasterSlaveBackendServerList
}

type DescribeMasterSlaveVServerGroupAttributeMasterSlaveBackendServerList []DescribeMasterSlaveVServerGroupAttributeMasterSlaveBackendServer

func (list *DescribeMasterSlaveVServerGroupAttributeMasterSlaveBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMasterSlaveVServerGroupAttributeMasterSlaveBackendServer)
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

func (c *SlbClient) DescribeAliCloudCertificates(req *DescribeAliCloudCertificatesArgs) (resp *DescribeAliCloudCertificatesResponse, err error) {
	resp = &DescribeAliCloudCertificatesResponse{}
	err = c.Request("DescribeAliCloudCertificates", req, resp)
	return
}

type DescribeAliCloudCertificatesAliCloudCertificate struct {
	AliCloudCertificateId   string
	AliCloudCertificateName string
	Fingerprint             string
	DomainName              string
	Issuer                  string
}
type DescribeAliCloudCertificatesArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeAliCloudCertificatesResponse struct {
	RequestId            string
	AliCloudCertificates DescribeAliCloudCertificatesAliCloudCertificateList
}

type DescribeAliCloudCertificatesAliCloudCertificateList []DescribeAliCloudCertificatesAliCloudCertificate

func (list *DescribeAliCloudCertificatesAliCloudCertificateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAliCloudCertificatesAliCloudCertificate)
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

func (c *SlbClient) SetRule(req *SetRuleArgs) (resp *SetRuleResponse, err error) {
	resp = &SetRuleResponse{}
	err = c.Request("SetRule", req, resp)
	return
}

type SetRuleArgs struct {
	Access_key_id        string
	VServerGroupId       string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	RuleId               string
	Tags                 string
}
type SetRuleResponse struct {
	RequestId string
}

func (c *SlbClient) SetBackendServers(req *SetBackendServersArgs) (resp *SetBackendServersResponse, err error) {
	resp = &SetBackendServersResponse{}
	err = c.Request("SetBackendServers", req, resp)
	return
}

type SetBackendServersBackendServer struct {
	ServerId string
	Weight   string
}
type SetBackendServersArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	BackendServers       string
	Tags                 string
}
type SetBackendServersResponse struct {
	RequestId      string
	LoadBalancerId string
	BackendServers SetBackendServersBackendServerList
}

type SetBackendServersBackendServerList []SetBackendServersBackendServer

func (list *SetBackendServersBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SetBackendServersBackendServer)
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

func (c *SlbClient) DescribeLoadBalancerTCPListenerAttribute(req *DescribeLoadBalancerTCPListenerAttributeArgs) (resp *DescribeLoadBalancerTCPListenerAttributeResponse, err error) {
	resp = &DescribeLoadBalancerTCPListenerAttributeResponse{}
	err = c.Request("DescribeLoadBalancerTCPListenerAttribute", req, resp)
	return
}

type DescribeLoadBalancerTCPListenerAttributeArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ListenerPort         int
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DescribeLoadBalancerTCPListenerAttributeResponse struct {
	RequestId                 string
	ListenerPort              int
	BackendServerPort         int
	Status                    string
	Bandwidth                 int
	Scheduler                 string
	SynProxy                  string
	PersistenceTimeout        int
	EstablishedTimeout        int
	HealthCheck               string
	HealthyThreshold          int
	UnhealthyThreshold        int
	HealthCheckConnectTimeout int
	HealthCheckConnectPort    int
	HealthCheckInterval       int
	HealthCheckHttpCode       string
	HealthCheckDomain         string
	HealthCheckURI            string
	HealthCheckType           string
	MaxConnection             int
	VServerGroupId            string
	MasterSlaveServerGroupId  string
}

func (c *SlbClient) DescribeListenerAccessControlAttribute(req *DescribeListenerAccessControlAttributeArgs) (resp *DescribeListenerAccessControlAttributeResponse, err error) {
	resp = &DescribeListenerAccessControlAttributeResponse{}
	err = c.Request("DescribeListenerAccessControlAttribute", req, resp)
	return
}

type DescribeListenerAccessControlAttributeArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ListenerPort         int
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DescribeListenerAccessControlAttributeResponse struct {
	RequestId           string
	AccessControlStatus string
	SourceItems         string
}

func (c *SlbClient) CreateLoadBalancerHTTPSListener(req *CreateLoadBalancerHTTPSListenerArgs) (resp *CreateLoadBalancerHTTPSListenerResponse, err error) {
	resp = &CreateLoadBalancerHTTPSListenerResponse{}
	err = c.Request("CreateLoadBalancerHTTPSListener", req, resp)
	return
}

type CreateLoadBalancerHTTPSListenerArgs struct {
	Access_key_id          string
	ResourceOwnerId        int64
	HealthCheckTimeout     int
	XForwardedFor          string
	HealthCheckURI         string
	UnhealthyThreshold     int
	HealthyThreshold       int
	Scheduler              string
	HealthCheck            string
	MaxConnection          int
	CookieTimeout          int
	StickySessionType      string
	VServerGroupId         string
	ListenerPort           int
	Cookie                 string
	ResourceOwnerAccount   string
	Bandwidth              int
	StickySession          string
	HealthCheckDomain      string
	OwnerAccount           string
	Gzip                   string
	OwnerId                int64
	ServerCertificateId    string
	CACertificateId        string
	Tags                   string
	LoadBalancerId         string
	XForwardedFor_SLBIP    string
	BackendServerPort      int
	HealthCheckInterval    int
	XForwardedFor_proto    string
	XForwardedFor_SLBID    string
	HealthCheckConnectPort int
	HealthCheckHttpCode    string
}
type CreateLoadBalancerHTTPSListenerResponse struct {
	RequestId string
}

func (c *SlbClient) SetLoadBalancerHTTPSListenerAttribute(req *SetLoadBalancerHTTPSListenerAttributeArgs) (resp *SetLoadBalancerHTTPSListenerAttributeResponse, err error) {
	resp = &SetLoadBalancerHTTPSListenerAttributeResponse{}
	err = c.Request("SetLoadBalancerHTTPSListenerAttribute", req, resp)
	return
}

type SetLoadBalancerHTTPSListenerAttributeArgs struct {
	Access_key_id          string
	ResourceOwnerId        int64
	HealthCheckTimeout     int
	XForwardedFor          string
	HealthCheckURI         string
	UnhealthyThreshold     int
	HealthyThreshold       int
	Scheduler              string
	HealthCheck            string
	MaxConnection          int
	CookieTimeout          int
	StickySessionType      string
	VServerGroupId         string
	ListenerPort           int
	Cookie                 string
	ResourceOwnerAccount   string
	Bandwidth              int
	StickySession          string
	HealthCheckDomain      string
	OwnerAccount           string
	Gzip                   string
	OwnerId                int64
	ServerCertificateId    string
	CACertificateId        string
	Tags                   string
	LoadBalancerId         string
	XForwardedFor_SLBIP    string
	HealthCheckInterval    int
	XForwardedFor_proto    string
	XForwardedFor_SLBID    string
	HealthCheckConnectPort int
	HealthCheckHttpCode    string
	VServerGroup           string
}
type SetLoadBalancerHTTPSListenerAttributeResponse struct {
	RequestId string
}

func (c *SlbClient) DescribeMasterSlaveServerGroupAttribute(req *DescribeMasterSlaveServerGroupAttributeArgs) (resp *DescribeMasterSlaveServerGroupAttributeResponse, err error) {
	resp = &DescribeMasterSlaveServerGroupAttributeResponse{}
	err = c.Request("DescribeMasterSlaveServerGroupAttribute", req, resp)
	return
}

type DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServer struct {
	ServerId   string
	Port       int
	Weight     int
	ServerType string
}
type DescribeMasterSlaveServerGroupAttributeArgs struct {
	Access_key_id            string
	ResourceOwnerId          int64
	MasterSlaveServerGroupId string
	ResourceOwnerAccount     string
	OwnerAccount             string
	OwnerId                  int64
	Tags                     string
}
type DescribeMasterSlaveServerGroupAttributeResponse struct {
	RequestId                  string
	MasterSlaveServerGroupId   string
	MasterSlaveServerGroupName string
	MasterSlaveBackendServers  DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServerList
}

type DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServerList []DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServer

func (list *DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServer)
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

func (c *SlbClient) CreateLoadBalancerTCPListener(req *CreateLoadBalancerTCPListenerArgs) (resp *CreateLoadBalancerTCPListenerResponse, err error) {
	resp = &CreateLoadBalancerTCPListenerResponse{}
	err = c.Request("CreateLoadBalancerTCPListener", req, resp)
	return
}

type CreateLoadBalancerTCPListenerArgs struct {
	Access_key_id             string
	HealthCheckConnectTimeout int
	ResourceOwnerId           int64
	HealthCheckURI            string
	UnhealthyThreshold        int
	HealthyThreshold          int
	Scheduler                 string
	EstablishedTimeout        int
	MaxConnection             int
	PersistenceTimeout        int
	VServerGroupId            string
	ListenerPort              int
	HealthCheckType           string
	ResourceOwnerAccount      string
	Bandwidth                 int
	HealthCheckDomain         string
	OwnerAccount              string
	OwnerId                   int64
	Tags                      string
	LoadBalancerId            string
	MasterSlaveServerGroupId  string
	BackendServerPort         int
	HealthCheckInterval       int
	HealthCheckConnectPort    int
	HealthCheckHttpCode       string
}
type CreateLoadBalancerTCPListenerResponse struct {
	RequestId string
}

func (c *SlbClient) DescribeMasterSlaveVServerGroups(req *DescribeMasterSlaveVServerGroupsArgs) (resp *DescribeMasterSlaveVServerGroupsResponse, err error) {
	resp = &DescribeMasterSlaveVServerGroupsResponse{}
	err = c.Request("DescribeMasterSlaveVServerGroups", req, resp)
	return
}

type DescribeMasterSlaveVServerGroupsMasterSlaveVServerGroup struct {
	MasterSlaveVServerGroupId   string
	MasterSlaveVServerGroupName string
}
type DescribeMasterSlaveVServerGroupsArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DescribeMasterSlaveVServerGroupsResponse struct {
	RequestId                string
	MasterSlaveVServerGroups DescribeMasterSlaveVServerGroupsMasterSlaveVServerGroupList
}

type DescribeMasterSlaveVServerGroupsMasterSlaveVServerGroupList []DescribeMasterSlaveVServerGroupsMasterSlaveVServerGroup

func (list *DescribeMasterSlaveVServerGroupsMasterSlaveVServerGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMasterSlaveVServerGroupsMasterSlaveVServerGroup)
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

func (c *SlbClient) DescribeLoadBalancerUDPListenerAttribute(req *DescribeLoadBalancerUDPListenerAttributeArgs) (resp *DescribeLoadBalancerUDPListenerAttributeResponse, err error) {
	resp = &DescribeLoadBalancerUDPListenerAttributeResponse{}
	err = c.Request("DescribeLoadBalancerUDPListenerAttribute", req, resp)
	return
}

type DescribeLoadBalancerUDPListenerAttributeArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ListenerPort         int
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DescribeLoadBalancerUDPListenerAttributeResponse struct {
	RequestId                 string
	ListenerPort              int
	BackendServerPort         int
	Status                    string
	Bandwidth                 int
	Scheduler                 string
	PersistenceTimeout        int
	HealthCheck               string
	HealthyThreshold          int
	UnhealthyThreshold        int
	HealthCheckConnectTimeout int
	HealthCheckConnectPort    int
	HealthCheckInterval       int
	HealthCheckReq            string
	HealthCheckExp            string
	MaxConnection             int
	VServerGroupId            string
	MasterSlaveServerGroupId  string
}

func (c *SlbClient) DeleteVServerGroup(req *DeleteVServerGroupArgs) (resp *DeleteVServerGroupResponse, err error) {
	resp = &DeleteVServerGroupResponse{}
	err = c.Request("DeleteVServerGroup", req, resp)
	return
}

type DeleteVServerGroupArgs struct {
	Access_key_id        string
	VServerGroupId       string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DeleteVServerGroupResponse struct {
	RequestId string
}

func (c *SlbClient) SetLogsDownloadAttribute(req *SetLogsDownloadAttributeArgs) (resp *SetLogsDownloadAttributeResponse, err error) {
	resp = &SetLogsDownloadAttributeResponse{}
	err = c.Request("SetLogsDownloadAttribute", req, resp)
	return
}

type SetLogsDownloadAttributeArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	LogType              string
	ResourceOwnerAccount string
	RoleName             string
	OwnerAccount         string
	OSSBucketName        string
	OwnerId              int64
	Tags                 string
}
type SetLogsDownloadAttributeResponse struct {
	RequestId string
}

func (c *SlbClient) DescribeVServerGroupAttribute(req *DescribeVServerGroupAttributeArgs) (resp *DescribeVServerGroupAttributeResponse, err error) {
	resp = &DescribeVServerGroupAttributeResponse{}
	err = c.Request("DescribeVServerGroupAttribute", req, resp)
	return
}

type DescribeVServerGroupAttributeBackendServer struct {
	ServerId string
	Port     int
	Weight   int
}
type DescribeVServerGroupAttributeArgs struct {
	Access_key_id        string
	VServerGroupId       string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DescribeVServerGroupAttributeResponse struct {
	RequestId        string
	VServerGroupId   string
	VServerGroupName string
	BackendServers   DescribeVServerGroupAttributeBackendServerList
}

type DescribeVServerGroupAttributeBackendServerList []DescribeVServerGroupAttributeBackendServer

func (list *DescribeVServerGroupAttributeBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVServerGroupAttributeBackendServer)
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

func (c *SlbClient) RemoveBackendServers(req *RemoveBackendServersArgs) (resp *RemoveBackendServersResponse, err error) {
	resp = &RemoveBackendServersResponse{}
	err = c.Request("RemoveBackendServers", req, resp)
	return
}

type RemoveBackendServersBackendServer struct {
	ServerId string
	Weight   int
}
type RemoveBackendServersArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	BackendServers       string
	Tags                 string
}
type RemoveBackendServersResponse struct {
	RequestId      string
	LoadBalancerId string
	BackendServers RemoveBackendServersBackendServerList
}

type RemoveBackendServersBackendServerList []RemoveBackendServersBackendServer

func (list *RemoveBackendServersBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveBackendServersBackendServer)
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

func (c *SlbClient) DeleteLogsDownloadAttribute(req *DeleteLogsDownloadAttributeArgs) (resp *DeleteLogsDownloadAttributeResponse, err error) {
	resp = &DeleteLogsDownloadAttributeResponse{}
	err = c.Request("DeleteLogsDownloadAttribute", req, resp)
	return
}

type DeleteLogsDownloadAttributeArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	RoleName             string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DeleteLogsDownloadAttributeResponse struct {
	RequestId string
}

func (c *SlbClient) CreateLoadBalancerUDPListener(req *CreateLoadBalancerUDPListenerArgs) (resp *CreateLoadBalancerUDPListenerResponse, err error) {
	resp = &CreateLoadBalancerUDPListenerResponse{}
	err = c.Request("CreateLoadBalancerUDPListener", req, resp)
	return
}

type CreateLoadBalancerUDPListenerArgs struct {
	Access_key_id             string
	HealthCheckConnectTimeout int
	ResourceOwnerId           int64
	UnhealthyThreshold        int
	HealthyThreshold          int
	Scheduler                 string
	MaxConnection             int
	PersistenceTimeout        int
	VServerGroupId            string
	ListenerPort              int
	ResourceOwnerAccount      string
	Bandwidth                 int
	OwnerAccount              string
	OwnerId                   int64
	Tags                      string
	LoadBalancerId            string
	MasterSlaveServerGroupId  string
	HealthCheckReq            string
	BackendServerPort         int
	HealthCheckInterval       int
	HealthCheckExp            string
	HealthCheckConnectPort    int
}
type CreateLoadBalancerUDPListenerResponse struct {
	RequestId string
}

func (c *SlbClient) DeleteLoadBalancer(req *DeleteLoadBalancerArgs) (resp *DeleteLoadBalancerResponse, err error) {
	resp = &DeleteLoadBalancerResponse{}
	err = c.Request("DeleteLoadBalancer", req, resp)
	return
}

type DeleteLoadBalancerArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DeleteLoadBalancerResponse struct {
	RequestId string
}

func (c *SlbClient) SetLogsDownloadStatus(req *SetLogsDownloadStatusArgs) (resp *SetLogsDownloadStatusResponse, err error) {
	resp = &SetLogsDownloadStatusResponse{}
	err = c.Request("SetLogsDownloadStatus", req, resp)
	return
}

type SetLogsDownloadStatusArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	RoleName             string
	OwnerAccount         string
	OwnerId              int64
	LogsDownloadStatus   string
	Tags                 string
}
type SetLogsDownloadStatusResponse struct {
	RequestId string
}

func (c *SlbClient) SetLoadBalancerTCPListenerAttribute(req *SetLoadBalancerTCPListenerAttributeArgs) (resp *SetLoadBalancerTCPListenerAttributeResponse, err error) {
	resp = &SetLoadBalancerTCPListenerAttributeResponse{}
	err = c.Request("SetLoadBalancerTCPListenerAttribute", req, resp)
	return
}

type SetLoadBalancerTCPListenerAttributeArgs struct {
	Access_key_id             string
	HealthCheckConnectTimeout int
	ResourceOwnerId           int64
	HealthCheckURI            string
	UnhealthyThreshold        int
	HealthyThreshold          int
	Scheduler                 string
	MasterSlaveServerGroup    string
	EstablishedTimeout        int
	MaxConnection             int
	PersistenceTimeout        int
	VServerGroupId            string
	ListenerPort              int
	HealthCheckType           string
	ResourceOwnerAccount      string
	Bandwidth                 int
	HealthCheckDomain         string
	OwnerAccount              string
	SynProxy                  string
	OwnerId                   int64
	Tags                      string
	LoadBalancerId            string
	MasterSlaveServerGroupId  string
	HealthCheckInterval       int
	HealthCheckConnectPort    int
	HealthCheckHttpCode       string
	VServerGroup              string
}
type SetLoadBalancerTCPListenerAttributeResponse struct {
	RequestId string
}

func (c *SlbClient) DescribeLogsDownloadAttribute(req *DescribeLogsDownloadAttributeArgs) (resp *DescribeLogsDownloadAttributeResponse, err error) {
	resp = &DescribeLogsDownloadAttributeResponse{}
	err = c.Request("DescribeLogsDownloadAttribute", req, resp)
	return
}

type DescribeLogsDownloadAttributeArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DescribeLogsDownloadAttributeResponse struct {
	RequestId     string
	LogType       string
	OSSBucketName string
	RoleName      string
}

func (c *SlbClient) SetVServerGroupAttribute(req *SetVServerGroupAttributeArgs) (resp *SetVServerGroupAttributeResponse, err error) {
	resp = &SetVServerGroupAttributeResponse{}
	err = c.Request("SetVServerGroupAttribute", req, resp)
	return
}

type SetVServerGroupAttributeBackendServer struct {
	ServerId string
	Port     int
	Weight   int
}
type SetVServerGroupAttributeArgs struct {
	Access_key_id        string
	VServerGroupId       string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	BackendServers       string
	Tags                 string
	VServerGroupName     string
}
type SetVServerGroupAttributeResponse struct {
	RequestId        string
	VServerGroupId   string
	VServerGroupName string
	BackendServers   SetVServerGroupAttributeBackendServerList
}

type SetVServerGroupAttributeBackendServerList []SetVServerGroupAttributeBackendServer

func (list *SetVServerGroupAttributeBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SetVServerGroupAttributeBackendServer)
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

func (c *SlbClient) SetLoadBalancerStatus(req *SetLoadBalancerStatusArgs) (resp *SetLoadBalancerStatusResponse, err error) {
	resp = &SetLoadBalancerStatusResponse{}
	err = c.Request("SetLoadBalancerStatus", req, resp)
	return
}

type SetLoadBalancerStatusArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	LoadBalancerStatus   string
	Tags                 string
}
type SetLoadBalancerStatusResponse struct {
	RequestId string
}

func (c *SlbClient) RemoveVServerGroupBackendServers(req *RemoveVServerGroupBackendServersArgs) (resp *RemoveVServerGroupBackendServersResponse, err error) {
	resp = &RemoveVServerGroupBackendServersResponse{}
	err = c.Request("RemoveVServerGroupBackendServers", req, resp)
	return
}

type RemoveVServerGroupBackendServersBackendServer struct {
	ServerId string
	Port     int
	Weight   int
}
type RemoveVServerGroupBackendServersArgs struct {
	Access_key_id        string
	VServerGroupId       string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	BackendServers       string
	Tags                 string
}
type RemoveVServerGroupBackendServersResponse struct {
	RequestId      string
	VServerGroupId string
	BackendServers RemoveVServerGroupBackendServersBackendServerList
}

type RemoveVServerGroupBackendServersBackendServerList []RemoveVServerGroupBackendServersBackendServer

func (list *RemoveVServerGroupBackendServersBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveVServerGroupBackendServersBackendServer)
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

func (c *SlbClient) RemoveTags(req *RemoveTagsArgs) (resp *RemoveTagsResponse, err error) {
	resp = &RemoveTagsResponse{}
	err = c.Request("RemoveTags", req, resp)
	return
}

type RemoveTagsArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type RemoveTagsResponse struct {
	RequestId string
}

func (c *SlbClient) ModifyVServerGroupBackendServers(req *ModifyVServerGroupBackendServersArgs) (resp *ModifyVServerGroupBackendServersResponse, err error) {
	resp = &ModifyVServerGroupBackendServersResponse{}
	err = c.Request("ModifyVServerGroupBackendServers", req, resp)
	return
}

type ModifyVServerGroupBackendServersBackendServer struct {
	ServerId string
	Port     int
	Weight   int
}
type ModifyVServerGroupBackendServersArgs struct {
	Access_key_id        string
	VServerGroupId       string
	ResourceOwnerId      int64
	OldBackendServers    string
	ResourceOwnerAccount string
	NewBackendServers    string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type ModifyVServerGroupBackendServersResponse struct {
	RequestId      string
	VServerGroupId string
	BackendServers ModifyVServerGroupBackendServersBackendServerList
}

type ModifyVServerGroupBackendServersBackendServerList []ModifyVServerGroupBackendServersBackendServer

func (list *ModifyVServerGroupBackendServersBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyVServerGroupBackendServersBackendServer)
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

func (c *SlbClient) DescribeRuleAttribute(req *DescribeRuleAttributeArgs) (resp *DescribeRuleAttributeResponse, err error) {
	resp = &DescribeRuleAttributeResponse{}
	err = c.Request("DescribeRuleAttribute", req, resp)
	return
}

type DescribeRuleAttributeArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	RuleId               string
	Tags                 string
}
type DescribeRuleAttributeResponse struct {
	RequestId      string
	RuleName       string
	LoadBalancerId string
	ListenerPort   string
	Domain         string
	Url            string
	VServerGroupId string
}

func (c *SlbClient) CreateLoadBalancer(req *CreateLoadBalancerArgs) (resp *CreateLoadBalancerResponse, err error) {
	resp = &CreateLoadBalancerResponse{}
	err = c.Request("CreateLoadBalancer", req, resp)
	return
}

type CreateLoadBalancerArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ClientToken          string
	MasterZoneId         string
	Duration             int
	ResourceGroupId      string
	LoadBalancerName     string
	AddressType          string
	SlaveZoneId          string
	LoadBalancerSpec     string
	AutoPay              core.Bool
	ResourceOwnerAccount string
	Bandwidth            int
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
	VSwitchId            string
	EnableVpcVipFlow     string
	InternetChargeType   string
	VpcId                string
	PayType              string
	PricingCycle         string
}
type CreateLoadBalancerResponse struct {
	RequestId        string
	LoadBalancerId   string
	ResourceGroupId  string
	Address          string
	LoadBalancerName string
	VpcId            string
	VSwitchId        string
	NetworkType      string
	OrderId          int64
}

func (c *SlbClient) CreateMasterSlaveVServerGroup(req *CreateMasterSlaveVServerGroupArgs) (resp *CreateMasterSlaveVServerGroupResponse, err error) {
	resp = &CreateMasterSlaveVServerGroupResponse{}
	err = c.Request("CreateMasterSlaveVServerGroup", req, resp)
	return
}

type CreateMasterSlaveVServerGroupMasterSlaveBackendServer struct {
	ServerId string
	Port     int
	Weight   int
	IsBackup int
}
type CreateMasterSlaveVServerGroupArgs struct {
	Access_key_id               string
	ResourceOwnerId             int64
	MasterSlaveBackendServers   string
	LoadBalancerId              string
	ResourceOwnerAccount        string
	OwnerAccount                string
	MasterSlaveVServerGroupName string
	OwnerId                     int64
	Tags                        string
}
type CreateMasterSlaveVServerGroupResponse struct {
	RequestId                 string
	MasterSlaveVServerGroupId string
	MasterSlaveBackendServers CreateMasterSlaveVServerGroupMasterSlaveBackendServerList
}

type CreateMasterSlaveVServerGroupMasterSlaveBackendServerList []CreateMasterSlaveVServerGroupMasterSlaveBackendServer

func (list *CreateMasterSlaveVServerGroupMasterSlaveBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateMasterSlaveVServerGroupMasterSlaveBackendServer)
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

func (c *SlbClient) DeleteMasterSlaveServerGroup(req *DeleteMasterSlaveServerGroupArgs) (resp *DeleteMasterSlaveServerGroupResponse, err error) {
	resp = &DeleteMasterSlaveServerGroupResponse{}
	err = c.Request("DeleteMasterSlaveServerGroup", req, resp)
	return
}

type DeleteMasterSlaveServerGroupArgs struct {
	Access_key_id            string
	ResourceOwnerId          int64
	MasterSlaveServerGroupId string
	ResourceOwnerAccount     string
	OwnerAccount             string
	OwnerId                  int64
	Tags                     string
}
type DeleteMasterSlaveServerGroupResponse struct {
	RequestId string
}

func (c *SlbClient) CreateMasterSlaveServerGroup(req *CreateMasterSlaveServerGroupArgs) (resp *CreateMasterSlaveServerGroupResponse, err error) {
	resp = &CreateMasterSlaveServerGroupResponse{}
	err = c.Request("CreateMasterSlaveServerGroup", req, resp)
	return
}

type CreateMasterSlaveServerGroupMasterSlaveBackendServer struct {
	ServerId   string
	Port       int
	Weight     int
	ServerType string
}
type CreateMasterSlaveServerGroupArgs struct {
	Access_key_id              string
	ResourceOwnerId            int64
	MasterSlaveBackendServers  string
	LoadBalancerId             string
	ResourceOwnerAccount       string
	OwnerAccount               string
	MasterSlaveServerGroupName string
	OwnerId                    int64
	Tags                       string
}
type CreateMasterSlaveServerGroupResponse struct {
	RequestId                 string
	MasterSlaveServerGroupId  string
	MasterSlaveBackendServers CreateMasterSlaveServerGroupMasterSlaveBackendServerList
}

type CreateMasterSlaveServerGroupMasterSlaveBackendServerList []CreateMasterSlaveServerGroupMasterSlaveBackendServer

func (list *CreateMasterSlaveServerGroupMasterSlaveBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateMasterSlaveServerGroupMasterSlaveBackendServer)
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

func (c *SlbClient) ModifyLoadBalancerInstanceSpec(req *ModifyLoadBalancerInstanceSpecArgs) (resp *ModifyLoadBalancerInstanceSpecResponse, err error) {
	resp = &ModifyLoadBalancerInstanceSpecResponse{}
	err = c.Request("ModifyLoadBalancerInstanceSpec", req, resp)
	return
}

type ModifyLoadBalancerInstanceSpecArgs struct {
	Access_key_id        string
	LoadBalancerSpec     string
	ResourceOwnerId      int64
	LoadBalancerId       string
	AutoPay              core.Bool
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type ModifyLoadBalancerInstanceSpecResponse struct {
	RequestId string
	OrderId   int64
}

func (c *SlbClient) StartLoadBalancerListener(req *StartLoadBalancerListenerArgs) (resp *StartLoadBalancerListenerResponse, err error) {
	resp = &StartLoadBalancerListenerResponse{}
	err = c.Request("StartLoadBalancerListener", req, resp)
	return
}

type StartLoadBalancerListenerArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ListenerPort         int
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type StartLoadBalancerListenerResponse struct {
	RequestId string
}

func (c *SlbClient) DeleteServerCertificate(req *DeleteServerCertificateArgs) (resp *DeleteServerCertificateResponse, err error) {
	resp = &DeleteServerCertificateResponse{}
	err = c.Request("DeleteServerCertificate", req, resp)
	return
}

type DeleteServerCertificateArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	ServerCertificateId  string
	Tags                 string
}
type DeleteServerCertificateResponse struct {
	RequestId string
}

func (c *SlbClient) UploadServerCertificate(req *UploadServerCertificateArgs) (resp *UploadServerCertificateResponse, err error) {
	resp = &UploadServerCertificateResponse{}
	err = c.Request("UploadServerCertificate", req, resp)
	return
}

type UploadServerCertificateArgs struct {
	Access_key_id           string
	ResourceOwnerId         int64
	ServerCertificate       string
	ResourceOwnerAccount    string
	OwnerAccount            string
	AliCloudCertificateName string
	AliCloudCertificateId   string
	OwnerId                 int64
	Tags                    string
	PrivateKey              string
	ResourceGroupId         string
	ServerCertificateName   string
}
type UploadServerCertificateResponse struct {
	RequestId               string
	ServerCertificateId     string
	Fingerprint             string
	ServerCertificateName   string
	RegionId                string
	RegionIdAlias           string
	AliCloudCertificateId   string
	AliCloudCertificateName string
	IsAliCloudCertificate   int
	ResourceGroupId         string
	CreateTime              string
	CreateTimeStamp         int64
}

func (c *SlbClient) DeleteMasterSlaveVServerGroup(req *DeleteMasterSlaveVServerGroupArgs) (resp *DeleteMasterSlaveVServerGroupResponse, err error) {
	resp = &DeleteMasterSlaveVServerGroupResponse{}
	err = c.Request("DeleteMasterSlaveVServerGroup", req, resp)
	return
}

type DeleteMasterSlaveVServerGroupArgs struct {
	Access_key_id             string
	ResourceOwnerId           int64
	MasterSlaveVServerGroupId string
	ResourceOwnerAccount      string
	OwnerAccount              string
	OwnerId                   int64
	Tags                      string
}
type DeleteMasterSlaveVServerGroupResponse struct {
	RequestId string
}

func (c *SlbClient) DescribeLoadBalancerHTTPListenerAttribute(req *DescribeLoadBalancerHTTPListenerAttributeArgs) (resp *DescribeLoadBalancerHTTPListenerAttributeResponse, err error) {
	resp = &DescribeLoadBalancerHTTPListenerAttributeResponse{}
	err = c.Request("DescribeLoadBalancerHTTPListenerAttribute", req, resp)
	return
}

type DescribeLoadBalancerHTTPListenerAttributeArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ListenerPort         int
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DescribeLoadBalancerHTTPListenerAttributeResponse struct {
	RequestId              string
	ListenerPort           int
	BackendServerPort      int
	Bandwidth              int
	Status                 string
	SecurityStatus         string
	XForwardedFor          string
	Scheduler              string
	StickySession          string
	StickySessionType      string
	CookieTimeout          int
	Cookie                 string
	HealthCheck            string
	HealthCheckDomain      string
	HealthCheckURI         string
	HealthyThreshold       int
	UnhealthyThreshold     int
	HealthCheckTimeout     int
	HealthCheckInterval    int
	HealthCheckConnectPort int
	HealthCheckHttpCode    string
	MaxConnection          int
	VServerGroupId         string
	Gzip                   string
	XForwardedFor_SLBIP    string
	XForwardedFor_SLBID    string
	XForwardedFor_proto    string
}

func (c *SlbClient) CreateVServerGroup(req *CreateVServerGroupArgs) (resp *CreateVServerGroupResponse, err error) {
	resp = &CreateVServerGroupResponse{}
	err = c.Request("CreateVServerGroup", req, resp)
	return
}

type CreateVServerGroupBackendServer struct {
	ServerId string
	Port     int
	Weight   int
}
type CreateVServerGroupArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	BackendServers       string
	Tags                 string
	VServerGroupName     string
}
type CreateVServerGroupResponse struct {
	RequestId      string
	VServerGroupId string
	BackendServers CreateVServerGroupBackendServerList
}

type CreateVServerGroupBackendServerList []CreateVServerGroupBackendServer

func (list *CreateVServerGroupBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateVServerGroupBackendServer)
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

func (c *SlbClient) DeleteLoadBalancerListener(req *DeleteLoadBalancerListenerArgs) (resp *DeleteLoadBalancerListenerResponse, err error) {
	resp = &DeleteLoadBalancerListenerResponse{}
	err = c.Request("DeleteLoadBalancerListener", req, resp)
	return
}

type DeleteLoadBalancerListenerArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ListenerPort         int
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DeleteLoadBalancerListenerResponse struct {
	RequestId string
}

func (c *SlbClient) DescribeLoadBalancerAttribute(req *DescribeLoadBalancerAttributeArgs) (resp *DescribeLoadBalancerAttributeResponse, err error) {
	resp = &DescribeLoadBalancerAttributeResponse{}
	err = c.Request("DescribeLoadBalancerAttribute", req, resp)
	return
}

type DescribeLoadBalancerAttributeListenerPortAndProtocal struct {
	ListenerPort     int
	ListenerProtocal string
}

type DescribeLoadBalancerAttributeListenerPortAndProtocol struct {
	ListenerPort     int
	ListenerProtocol string
}

type DescribeLoadBalancerAttributeBackendServer struct {
	ServerId string
	Weight   int
}
type DescribeLoadBalancerAttributeArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DescribeLoadBalancerAttributeResponse struct {
	RequestId                string
	LoadBalancerId           string
	ResourceGroupId          string
	LoadBalancerName         string
	LoadBalancerStatus       string
	RegionId                 string
	RegionIdAlias            string
	Address                  string
	AddressType              string
	VpcId                    string
	VSwitchId                string
	NetworkType              string
	InternetChargeType       string
	AutoReleaseTime          int64
	Bandwidth                int
	LoadBalancerSpec         string
	CreateTime               string
	CreateTimeStamp          int64
	EndTime                  string
	EndTimeStamp             int64
	PayType                  string
	MasterZoneId             string
	SlaveZoneId              string
	ListenerPortsAndProtocal DescribeLoadBalancerAttributeListenerPortAndProtocalList
	ListenerPortsAndProtocol DescribeLoadBalancerAttributeListenerPortAndProtocolList
	BackendServers           DescribeLoadBalancerAttributeBackendServerList
	ListenerPorts            DescribeLoadBalancerAttributeListenerPortList
}

type DescribeLoadBalancerAttributeListenerPortAndProtocalList []DescribeLoadBalancerAttributeListenerPortAndProtocal

func (list *DescribeLoadBalancerAttributeListenerPortAndProtocalList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancerAttributeListenerPortAndProtocal)
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

type DescribeLoadBalancerAttributeListenerPortAndProtocolList []DescribeLoadBalancerAttributeListenerPortAndProtocol

func (list *DescribeLoadBalancerAttributeListenerPortAndProtocolList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancerAttributeListenerPortAndProtocol)
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

type DescribeLoadBalancerAttributeBackendServerList []DescribeLoadBalancerAttributeBackendServer

func (list *DescribeLoadBalancerAttributeBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancerAttributeBackendServer)
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

type DescribeLoadBalancerAttributeListenerPortList []string

func (list *DescribeLoadBalancerAttributeListenerPortList) UnmarshalJSON(data []byte) error {
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

func (c *SlbClient) ModifyLoadBalancerInternetSpec(req *ModifyLoadBalancerInternetSpecArgs) (resp *ModifyLoadBalancerInternetSpecResponse, err error) {
	resp = &ModifyLoadBalancerInternetSpecResponse{}
	err = c.Request("ModifyLoadBalancerInternetSpec", req, resp)
	return
}

type ModifyLoadBalancerInternetSpecArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	AutoPay              core.Bool
	ResourceOwnerAccount string
	Bandwidth            int
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
	LoadBalancerId       string
	InternetChargeType   string
}
type ModifyLoadBalancerInternetSpecResponse struct {
	RequestId string
	OrderId   int64
}

func (c *SlbClient) DescribeLoadBalancersRelatedEcs(req *DescribeLoadBalancersRelatedEcsArgs) (resp *DescribeLoadBalancersRelatedEcsResponse, err error) {
	resp = &DescribeLoadBalancersRelatedEcsResponse{}
	err = c.Request("DescribeLoadBalancersRelatedEcs", req, resp)
	return
}

type DescribeLoadBalancersRelatedEcsLoadBalancer struct {
	LoadBalancerId           string
	Count                    int
	MasterSlaveVServerGroups DescribeLoadBalancersRelatedEcsMasterSlaveVServerGroupList
	VServerGroups            DescribeLoadBalancersRelatedEcsVServerGroupList
	BackendServers           DescribeLoadBalancersRelatedEcsBackendServer4List
}

type DescribeLoadBalancersRelatedEcsMasterSlaveVServerGroup struct {
	GroupId         string
	GroupName       string
	BackendServers1 DescribeLoadBalancersRelatedEcsBackendServerList
}

type DescribeLoadBalancersRelatedEcsBackendServer struct {
	VmName      string
	Weight      int
	Port        int
	NetworkType string
}

type DescribeLoadBalancersRelatedEcsVServerGroup struct {
	GroupId         string
	GroupName       string
	BackendServers2 DescribeLoadBalancersRelatedEcsBackendServer3List
}

type DescribeLoadBalancersRelatedEcsBackendServer3 struct {
	VmName      string
	Weight      int
	Port        int
	NetworkType string
}

type DescribeLoadBalancersRelatedEcsBackendServer4 struct {
	VmName      string
	Weight      int
	Port        int
	NetworkType string
}
type DescribeLoadBalancersRelatedEcsArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DescribeLoadBalancersRelatedEcsResponse struct {
	Message       string
	Success       core.Bool
	RequestId     string
	LoadBalancers DescribeLoadBalancersRelatedEcsLoadBalancerList
}

type DescribeLoadBalancersRelatedEcsMasterSlaveVServerGroupList []DescribeLoadBalancersRelatedEcsMasterSlaveVServerGroup

func (list *DescribeLoadBalancersRelatedEcsMasterSlaveVServerGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersRelatedEcsMasterSlaveVServerGroup)
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

type DescribeLoadBalancersRelatedEcsVServerGroupList []DescribeLoadBalancersRelatedEcsVServerGroup

func (list *DescribeLoadBalancersRelatedEcsVServerGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersRelatedEcsVServerGroup)
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

type DescribeLoadBalancersRelatedEcsBackendServer4List []DescribeLoadBalancersRelatedEcsBackendServer4

func (list *DescribeLoadBalancersRelatedEcsBackendServer4List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersRelatedEcsBackendServer4)
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

type DescribeLoadBalancersRelatedEcsBackendServerList []DescribeLoadBalancersRelatedEcsBackendServer

func (list *DescribeLoadBalancersRelatedEcsBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersRelatedEcsBackendServer)
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

type DescribeLoadBalancersRelatedEcsBackendServer3List []DescribeLoadBalancersRelatedEcsBackendServer3

func (list *DescribeLoadBalancersRelatedEcsBackendServer3List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersRelatedEcsBackendServer3)
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

type DescribeLoadBalancersRelatedEcsLoadBalancerList []DescribeLoadBalancersRelatedEcsLoadBalancer

func (list *DescribeLoadBalancersRelatedEcsLoadBalancerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersRelatedEcsLoadBalancer)
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

func (c *SlbClient) DescribeZones(req *DescribeZonesArgs) (resp *DescribeZonesResponse, err error) {
	resp = &DescribeZonesResponse{}
	err = c.Request("DescribeZones", req, resp)
	return
}

type DescribeZonesZone struct {
	ZoneId     string
	LocalName  string
	SlaveZones DescribeZonesSlaveZoneList
}

type DescribeZonesSlaveZone struct {
	ZoneId    string
	LocalName string
}
type DescribeZonesArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DescribeZonesResponse struct {
	RequestId string
	Zones     DescribeZonesZoneList
}

type DescribeZonesSlaveZoneList []DescribeZonesSlaveZone

func (list *DescribeZonesSlaveZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZonesSlaveZone)
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

type DescribeZonesZoneList []DescribeZonesZone

func (list *DescribeZonesZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZonesZone)
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

func (c *SlbClient) CreateRules(req *CreateRulesArgs) (resp *CreateRulesResponse, err error) {
	resp = &CreateRulesResponse{}
	err = c.Request("CreateRules", req, resp)
	return
}

type CreateRulesRule struct {
	RuleId   string
	RuleName string
}
type CreateRulesArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ListenerPort         int
	LoadBalancerId       string
	ResourceOwnerAccount string
	RuleList             string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type CreateRulesResponse struct {
	RequestId string
	Rules     CreateRulesRuleList
}

type CreateRulesRuleList []CreateRulesRule

func (list *CreateRulesRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateRulesRule)
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

func (c *SlbClient) DescribeLogsDownloadStatus(req *DescribeLogsDownloadStatusArgs) (resp *DescribeLogsDownloadStatusResponse, err error) {
	resp = &DescribeLogsDownloadStatusResponse{}
	err = c.Request("DescribeLogsDownloadStatus", req, resp)
	return
}

type DescribeLogsDownloadStatusArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DescribeLogsDownloadStatusResponse struct {
	RequestId          string
	LogsDownloadStatus string
}

func (c *SlbClient) RemoveListenerWhiteListItem(req *RemoveListenerWhiteListItemArgs) (resp *RemoveListenerWhiteListItemResponse, err error) {
	resp = &RemoveListenerWhiteListItemResponse{}
	err = c.Request("RemoveListenerWhiteListItem", req, resp)
	return
}

type RemoveListenerWhiteListItemArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ListenerPort         int
	LoadBalancerId       string
	SourceItems          string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type RemoveListenerWhiteListItemResponse struct {
	RequestId string
}

func (c *SlbClient) DescribeRegions(req *DescribeRegionsArgs) (resp *DescribeRegionsResponse, err error) {
	resp = &DescribeRegionsResponse{}
	err = c.Request("DescribeRegions", req, resp)
	return
}

type DescribeRegionsRegion struct {
	RegionId  string
	LocalName string
}
type DescribeRegionsArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DescribeRegionsResponse struct {
	RequestId string
	Regions   DescribeRegionsRegionList
}

type DescribeRegionsRegionList []DescribeRegionsRegion

func (list *DescribeRegionsRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsRegion)
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

func (c *SlbClient) SetLoadBalancerName(req *SetLoadBalancerNameArgs) (resp *SetLoadBalancerNameResponse, err error) {
	resp = &SetLoadBalancerNameResponse{}
	err = c.Request("SetLoadBalancerName", req, resp)
	return
}

type SetLoadBalancerNameArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	LoadBalancerName     string
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type SetLoadBalancerNameResponse struct {
	RequestId string
}

func (c *SlbClient) UploadCACertificate(req *UploadCACertificateArgs) (resp *UploadCACertificateResponse, err error) {
	resp = &UploadCACertificateResponse{}
	err = c.Request("UploadCACertificate", req, resp)
	return
}

type UploadCACertificateArgs struct {
	Access_key_id        string
	ResourceGroupId      string
	ResourceOwnerId      int64
	CACertificate        string
	CACertificateName    string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type UploadCACertificateResponse struct {
	RequestId         string
	CACertificateId   string
	CACertificateName string
	Fingerprint       string
	ResourceGroupId   string
	CreateTime        string
	CreateTimeStamp   int64
}

func (c *SlbClient) DescribeHealthStatus(req *DescribeHealthStatusArgs) (resp *DescribeHealthStatusResponse, err error) {
	resp = &DescribeHealthStatusResponse{}
	err = c.Request("DescribeHealthStatus", req, resp)
	return
}

type DescribeHealthStatusBackendServer struct {
	ListenerPort       int
	ServerId           string
	Port               int
	ServerHealthStatus string
}
type DescribeHealthStatusArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ListenerPort         int
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DescribeHealthStatusResponse struct {
	RequestId      string
	BackendServers DescribeHealthStatusBackendServerList
}

type DescribeHealthStatusBackendServerList []DescribeHealthStatusBackendServer

func (list *DescribeHealthStatusBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeHealthStatusBackendServer)
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

func (c *SlbClient) DescribeMasterSlaveServerGroups(req *DescribeMasterSlaveServerGroupsArgs) (resp *DescribeMasterSlaveServerGroupsResponse, err error) {
	resp = &DescribeMasterSlaveServerGroupsResponse{}
	err = c.Request("DescribeMasterSlaveServerGroups", req, resp)
	return
}

type DescribeMasterSlaveServerGroupsMasterSlaveServerGroup struct {
	MasterSlaveServerGroupId   string
	MasterSlaveServerGroupName string
}
type DescribeMasterSlaveServerGroupsArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DescribeMasterSlaveServerGroupsResponse struct {
	RequestId               string
	MasterSlaveServerGroups DescribeMasterSlaveServerGroupsMasterSlaveServerGroupList
}

type DescribeMasterSlaveServerGroupsMasterSlaveServerGroupList []DescribeMasterSlaveServerGroupsMasterSlaveServerGroup

func (list *DescribeMasterSlaveServerGroupsMasterSlaveServerGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMasterSlaveServerGroupsMasterSlaveServerGroup)
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

func (c *SlbClient) SetLoadBalancerHTTPListenerAttribute(req *SetLoadBalancerHTTPListenerAttributeArgs) (resp *SetLoadBalancerHTTPListenerAttributeResponse, err error) {
	resp = &SetLoadBalancerHTTPListenerAttributeResponse{}
	err = c.Request("SetLoadBalancerHTTPListenerAttribute", req, resp)
	return
}

type SetLoadBalancerHTTPListenerAttributeArgs struct {
	Access_key_id          string
	ResourceOwnerId        int64
	HealthCheckTimeout     int
	XForwardedFor          string
	HealthCheckURI         string
	UnhealthyThreshold     int
	HealthyThreshold       int
	Scheduler              string
	HealthCheck            string
	MaxConnection          int
	CookieTimeout          int
	StickySessionType      string
	VServerGroupId         string
	ListenerPort           int
	Cookie                 string
	ResourceOwnerAccount   string
	Bandwidth              int
	StickySession          string
	HealthCheckDomain      string
	OwnerAccount           string
	Gzip                   string
	OwnerId                int64
	Tags                   string
	LoadBalancerId         string
	XForwardedFor_SLBIP    string
	HealthCheckInterval    int
	XForwardedFor_proto    string
	XForwardedFor_SLBID    string
	HealthCheckConnectPort int
	HealthCheckHttpCode    string
	VServerGroup           string
}
type SetLoadBalancerHTTPListenerAttributeResponse struct {
	RequestId string
}

func (c *SlbClient) AddVServerGroupBackendServers(req *AddVServerGroupBackendServersArgs) (resp *AddVServerGroupBackendServersResponse, err error) {
	resp = &AddVServerGroupBackendServersResponse{}
	err = c.Request("AddVServerGroupBackendServers", req, resp)
	return
}

type AddVServerGroupBackendServersBackendServer struct {
	ServerId string
	Port     int
	Weight   int
}
type AddVServerGroupBackendServersArgs struct {
	Access_key_id        string
	VServerGroupId       string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	BackendServers       string
	Tags                 string
}
type AddVServerGroupBackendServersResponse struct {
	RequestId      string
	VServerGroupId string
	BackendServers AddVServerGroupBackendServersBackendServerList
}

type AddVServerGroupBackendServersBackendServerList []AddVServerGroupBackendServersBackendServer

func (list *AddVServerGroupBackendServersBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddVServerGroupBackendServersBackendServer)
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

func (c *SlbClient) StopLoadBalancerListener(req *StopLoadBalancerListenerArgs) (resp *StopLoadBalancerListenerResponse, err error) {
	resp = &StopLoadBalancerListenerResponse{}
	err = c.Request("StopLoadBalancerListener", req, resp)
	return
}

type StopLoadBalancerListenerArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ListenerPort         int
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type StopLoadBalancerListenerResponse struct {
	RequestId string
}

func (c *SlbClient) SetLoadBalancerAutoReleaseTime(req *SetLoadBalancerAutoReleaseTimeArgs) (resp *SetLoadBalancerAutoReleaseTimeResponse, err error) {
	resp = &SetLoadBalancerAutoReleaseTimeResponse{}
	err = c.Request("SetLoadBalancerAutoReleaseTime", req, resp)
	return
}

type SetLoadBalancerAutoReleaseTimeArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	AutoReleaseTime      int64
	OwnerId              int64
	Tags                 string
}
type SetLoadBalancerAutoReleaseTimeResponse struct {
	RequestId string
}

func (c *SlbClient) AddTags(req *AddTagsArgs) (resp *AddTagsResponse, err error) {
	resp = &AddTagsResponse{}
	err = c.Request("AddTags", req, resp)
	return
}

type AddTagsArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type AddTagsResponse struct {
	RequestId string
}

func (c *SlbClient) AddBackendServers(req *AddBackendServersArgs) (resp *AddBackendServersResponse, err error) {
	resp = &AddBackendServersResponse{}
	err = c.Request("AddBackendServers", req, resp)
	return
}

type AddBackendServersBackendServer struct {
	ServerId string
	Weight   string
}
type AddBackendServersArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	BackendServers       string
	Tags                 string
}
type AddBackendServersResponse struct {
	RequestId      string
	LoadBalancerId string
	BackendServers AddBackendServersBackendServerList
}

type AddBackendServersBackendServerList []AddBackendServersBackendServer

func (list *AddBackendServersBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddBackendServersBackendServer)
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

func (c *SlbClient) ModifyLoadBalancerPayType(req *ModifyLoadBalancerPayTypeArgs) (resp *ModifyLoadBalancerPayTypeResponse, err error) {
	resp = &ModifyLoadBalancerPayTypeResponse{}
	err = c.Request("ModifyLoadBalancerPayType", req, resp)
	return
}

type ModifyLoadBalancerPayTypeArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	AutoPay              core.Bool
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
	Duration             int
	LoadBalancerId       string
	PayType              string
	PricingCycle         string
}
type ModifyLoadBalancerPayTypeResponse struct {
	RequestId string
	OrderId   int64
}

func (c *SlbClient) DescribeCACertificates(req *DescribeCACertificatesArgs) (resp *DescribeCACertificatesResponse, err error) {
	resp = &DescribeCACertificatesResponse{}
	err = c.Request("DescribeCACertificates", req, resp)
	return
}

type DescribeCACertificatesCACertificate struct {
	RegionId          string
	CACertificateId   string
	CACertificateName string
	Fingerprint       string
	ResourceGroupId   string
	CreateTime        string
	CreateTimeStamp   int64
}
type DescribeCACertificatesArgs struct {
	Access_key_id        string
	ResourceGroupId      string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	CACertificateId      string
}
type DescribeCACertificatesResponse struct {
	RequestId      string
	CACertificates DescribeCACertificatesCACertificateList
}

type DescribeCACertificatesCACertificateList []DescribeCACertificatesCACertificate

func (list *DescribeCACertificatesCACertificateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCACertificatesCACertificate)
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

func (c *SlbClient) DescribeRules(req *DescribeRulesArgs) (resp *DescribeRulesResponse, err error) {
	resp = &DescribeRulesResponse{}
	err = c.Request("DescribeRules", req, resp)
	return
}

type DescribeRulesRule struct {
	RuleId         string
	RuleName       string
	Domain         string
	Url            string
	VServerGroupId string
}
type DescribeRulesArgs struct {
	Access_key_id        string
	ResourceOwnerId      int64
	ListenerPort         int
	LoadBalancerId       string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tags                 string
}
type DescribeRulesResponse struct {
	RequestId string
	Rules     DescribeRulesRuleList
}

type DescribeRulesRuleList []DescribeRulesRule

func (list *DescribeRulesRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRulesRule)
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

func (c *SlbClient) DescribeRealtimeLogs(req *DescribeRealtimeLogsArgs) (resp *DescribeRealtimeLogsResponse, err error) {
	resp = &DescribeRealtimeLogsResponse{}
	err = c.Request("DescribeRealtimeLogs", req, resp)
	return
}

type DescribeRealtimeLogsLBRealTimeLog struct {
	LogDetail string
}
type DescribeRealtimeLogsArgs struct {
	Access_key_id        string
	LogStartTime         string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	LogEndTime           string
	OwnerAccount         string
	OwnerId              int64
	PageNumber           int
	Tags                 string
	LogType              string
	LoadBalancerId       string
	PageSize             int
}
type DescribeRealtimeLogsResponse struct {
	RequestId         string
	PageNumber        int
	PageSize          int
	TotalCount        int64
	Progress          string
	LBRealTimeLogsSet DescribeRealtimeLogsLBRealTimeLogList
}

type DescribeRealtimeLogsLBRealTimeLogList []DescribeRealtimeLogsLBRealTimeLog

func (list *DescribeRealtimeLogsLBRealTimeLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRealtimeLogsLBRealTimeLog)
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
