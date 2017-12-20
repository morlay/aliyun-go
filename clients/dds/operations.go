package dds

import (
	"encoding/json"

	"github.com/morlay/aliyun-go/core"
)

func (c *DdsClient) DescribeShardingNetworkAddress(req *DescribeShardingNetworkAddressArgs) (resp *DescribeShardingNetworkAddressResponse, err error) {
	resp = &DescribeShardingNetworkAddressResponse{}
	err = c.Request("DescribeShardingNetworkAddress", req, resp)
	return
}

type DescribeShardingNetworkAddressNetworkAddress struct {
	NetworkAddress string
	IPAddress      string
	NetworkType    string
	Port           string
	VPCId          string
	VswitchId      string
	NodeId         string
	ExpiredTime    string
}
type DescribeShardingNetworkAddressArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
	NodeId               string
}
type DescribeShardingNetworkAddressResponse struct {
	RequestId        string
	NetworkAddresses DescribeShardingNetworkAddressNetworkAddressList
}

type DescribeShardingNetworkAddressNetworkAddressList []DescribeShardingNetworkAddressNetworkAddress

func (list *DescribeShardingNetworkAddressNetworkAddressList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeShardingNetworkAddressNetworkAddress)
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

func (c *DdsClient) DescribeBackupPolicy(req *DescribeBackupPolicyArgs) (resp *DescribeBackupPolicyResponse, err error) {
	resp = &DescribeBackupPolicyResponse{}
	err = c.Request("DescribeBackupPolicy", req, resp)
	return
}

type DescribeBackupPolicyArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeBackupPolicyResponse struct {
	RequestId             string
	BackupRetentionPeriod string
	PreferredBackupTime   string
	PreferredBackupPeriod string
}

func (c *DdsClient) DescribeReplicaUsage(req *DescribeReplicaUsageArgs) (resp *DescribeReplicaUsageResponse, err error) {
	resp = &DescribeReplicaUsageResponse{}
	err = c.Request("DescribeReplicaUsage", req, resp)
	return
}

type DescribeReplicaUsagePerformanceKeys struct {
	PerformanceKey DescribeReplicaUsagePerformanceKeyItemList
}

type DescribeReplicaUsagePerformanceKeyItem struct {
	Key               string
	Unit              string
	ValueFormat       string
	PerformanceValues DescribeReplicaUsagePerformanceValues
}

type DescribeReplicaUsagePerformanceValues struct {
	PerformanceValue DescribeReplicaUsagePerformanceValueItemList
}

type DescribeReplicaUsagePerformanceValueItem struct {
	Value string
	Date  string
}
type DescribeReplicaUsageArgs struct {
	ResourceOwnerId      int64
	SourceDBInstanceId   string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	ReplicaId            string
	OwnerId              int64
}
type DescribeReplicaUsageResponse struct {
	RequestId       string
	StartTime       string
	EndTime         string
	ReplicaId       string
	PerformanceKeys DescribeReplicaUsagePerformanceKeys
}

type DescribeReplicaUsagePerformanceKeyItemList []DescribeReplicaUsagePerformanceKeyItem

func (list *DescribeReplicaUsagePerformanceKeyItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicaUsagePerformanceKeyItem)
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

type DescribeReplicaUsagePerformanceValueItemList []DescribeReplicaUsagePerformanceValueItem

func (list *DescribeReplicaUsagePerformanceValueItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicaUsagePerformanceValueItem)
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

func (c *DdsClient) ModifyReplicaDescription(req *ModifyReplicaDescriptionArgs) (resp *ModifyReplicaDescriptionResponse, err error) {
	resp = &ModifyReplicaDescriptionResponse{}
	err = c.Request("ModifyReplicaDescription", req, resp)
	return
}

type ModifyReplicaDescriptionArgs struct {
	ReplicaDescription   string
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	ReplicaId            string
	OwnerId              int64
}
type ModifyReplicaDescriptionResponse struct {
	RequestId string
}

func (c *DdsClient) ModifyAccountDescription(req *ModifyAccountDescriptionArgs) (resp *ModifyAccountDescriptionResponse, err error) {
	resp = &ModifyAccountDescriptionResponse{}
	err = c.Request("ModifyAccountDescription", req, resp)
	return
}

type ModifyAccountDescriptionArgs struct {
	ResourceOwnerId      int64
	AccountName          string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
	AccountDescription   string
}
type ModifyAccountDescriptionResponse struct {
	RequestId string
}

func (c *DdsClient) ModifyDBInstanceSpec(req *ModifyDBInstanceSpecArgs) (resp *ModifyDBInstanceSpecResponse, err error) {
	resp = &ModifyDBInstanceSpecResponse{}
	err = c.Request("ModifyDBInstanceSpec", req, resp)
	return
}

type ModifyDBInstanceSpecArgs struct {
	ResourceOwnerId      int64
	DBInstanceStorage    string
	AutoPay              core.Bool
	FromApp              string
	ResourceOwnerAccount string
	OwnerAccount         string
	CouponNo             string
	OwnerId              int64
	DBInstanceClass      string
	SecurityToken        string
	DBInstanceId         string
	BusinessInfo         string
	OrderType            string
}
type ModifyDBInstanceSpecResponse struct {
	RequestId string
	OrderId   string
}

func (c *DdsClient) ModifyDBInstanceMaintainTime(req *ModifyDBInstanceMaintainTimeArgs) (resp *ModifyDBInstanceMaintainTimeResponse, err error) {
	resp = &ModifyDBInstanceMaintainTimeResponse{}
	err = c.Request("ModifyDBInstanceMaintainTime", req, resp)
	return
}

type ModifyDBInstanceMaintainTimeArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	MaintainStartTime    string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
	MaintainEndTime      string
}
type ModifyDBInstanceMaintainTimeResponse struct {
	RequestId string
}

func (c *DdsClient) RenewDBInstance(req *RenewDBInstanceArgs) (resp *RenewDBInstanceResponse, err error) {
	resp = &RenewDBInstanceResponse{}
	err = c.Request("RenewDBInstance", req, resp)
	return
}

type RenewDBInstanceArgs struct {
	ResourceOwnerId      int64
	Period               int
	AutoPay              core.Bool
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	CouponNo             string
	OwnerId              int64
	SecurityToken        string
	DBInstanceId         string
	BusinessInfo         string
}
type RenewDBInstanceResponse struct {
	RequestId string
	OrderId   string
}

func (c *DdsClient) CreateDBInstance(req *CreateDBInstanceArgs) (resp *CreateDBInstanceResponse, err error) {
	resp = &CreateDBInstanceResponse{}
	err = c.Request("CreateDBInstance", req, resp)
	return
}

type CreateDBInstanceArgs struct {
	ResourceOwnerId       int64
	DBInstanceStorage     int
	ClientToken           string
	CouponNo              string
	EngineVersion         string
	NetworkType           string
	StorageEngine         string
	SecurityToken         string
	Engine                string
	DBInstanceDescription string
	BusinessInfo          string
	Period                int
	RestoreTime           string
	ResourceOwnerAccount  string
	SrcDBInstanceId       string
	OwnerAccount          string
	BackupId              string
	OwnerId               int64
	DBInstanceClass       string
	SecurityIPList        string
	VSwitchId             string
	AccountPassword       string
	VpcId                 string
	ZoneId                string
	ChargeType            string
}
type CreateDBInstanceResponse struct {
	RequestId    string
	OrderId      string
	DBInstanceId string
}

func (c *DdsClient) ResetAccountPassword(req *ResetAccountPasswordArgs) (resp *ResetAccountPasswordResponse, err error) {
	resp = &ResetAccountPasswordResponse{}
	err = c.Request("ResetAccountPassword", req, resp)
	return
}

type ResetAccountPasswordArgs struct {
	ResourceOwnerId      int64
	AccountPassword      string
	AccountName          string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type ResetAccountPasswordResponse struct {
	RequestId string
}

func (c *DdsClient) CreateBackup(req *CreateBackupArgs) (resp *CreateBackupResponse, err error) {
	resp = &CreateBackupResponse{}
	err = c.Request("CreateBackup", req, resp)
	return
}

type CreateBackupArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type CreateBackupResponse struct {
	RequestId string
	BackupId  string
}

func (c *DdsClient) DescribeBackups(req *DescribeBackupsArgs) (resp *DescribeBackupsResponse, err error) {
	resp = &DescribeBackupsResponse{}
	err = c.Request("DescribeBackups", req, resp)
	return
}

type DescribeBackupsBackup struct {
	BackupDBNames     string
	BackupId          int
	BackupStatus      string
	BackupStartTime   string
	BackupEndTime     string
	BackupType        string
	BackupMode        string
	BackupMethod      string
	BackupDownloadURL string
	BackupSize        int64
}
type DescribeBackupsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	BackupId             int
	EndTime              string
	StartTime            string
	OwnerId              int64
	PageNumber           int
	SecurityToken        string
	PageSize             int
	DBInstanceId         string
	NodeId               string
}
type DescribeBackupsResponse struct {
	RequestId  string
	PageNumber int
	PageSize   int
	TotalCount int
	Backups    DescribeBackupsBackupList
}

type DescribeBackupsBackupList []DescribeBackupsBackup

func (list *DescribeBackupsBackupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBackupsBackup)
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

func (c *DdsClient) RestartDBInstance(req *RestartDBInstanceArgs) (resp *RestartDBInstanceResponse, err error) {
	resp = &RestartDBInstanceResponse{}
	err = c.Request("RestartDBInstance", req, resp)
	return
}

type RestartDBInstanceArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
	NodeId               string
}
type RestartDBInstanceResponse struct {
	RequestId string
}

func (c *DdsClient) DescribeDBInstances(req *DescribeDBInstancesArgs) (resp *DescribeDBInstancesResponse, err error) {
	resp = &DescribeDBInstancesResponse{}
	err = c.Request("DescribeDBInstances", req, resp)
	return
}

type DescribeDBInstancesDBInstance struct {
	DBInstanceId          string
	DBInstanceDescription string
	RegionId              string
	ZoneId                string
	Engine                string
	EngineVersion         string
	DBInstanceClass       string
	DBInstanceStorage     int
	DBInstanceStatus      string
	LockMode              string
	ChargeType            string
	NetworkType           string
	CreationTime          string
	ExpireTime            string
	DBInstanceType        string
	LastDowngradeTime     int
	MongosList            DescribeDBInstancesMongosAttributeList
	ShardList             DescribeDBInstancesShardAttributeList
}

type DescribeDBInstancesMongosAttribute struct {
	NodeId          string
	NodeDescription string
	NodeClass       string
	ConnectSting    string
	Port            int
}

type DescribeDBInstancesShardAttribute struct {
	NodeId          string
	NodeDescription string
	NodeClass       string
	NodeStorage     int
}
type DescribeDBInstancesArgs struct {
	ResourceOwnerId      int64
	DBInstanceIds        string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	PageNumber           int
	DBInstanceType       string
	SecurityToken        string
	Engine               string
	PageSize             int
}
type DescribeDBInstancesResponse struct {
	RequestId   string
	PageNumber  int
	PageSize    int
	TotalCount  int
	DBInstances DescribeDBInstancesDBInstanceList
}

type DescribeDBInstancesMongosAttributeList []DescribeDBInstancesMongosAttribute

func (list *DescribeDBInstancesMongosAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesMongosAttribute)
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

type DescribeDBInstancesShardAttributeList []DescribeDBInstancesShardAttribute

func (list *DescribeDBInstancesShardAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesShardAttribute)
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

type DescribeDBInstancesDBInstanceList []DescribeDBInstancesDBInstance

func (list *DescribeDBInstancesDBInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesDBInstance)
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

func (c *DdsClient) ModifyNodeSpec(req *ModifyNodeSpecArgs) (resp *ModifyNodeSpecResponse, err error) {
	resp = &ModifyNodeSpecResponse{}
	err = c.Request("ModifyNodeSpec", req, resp)
	return
}

type ModifyNodeSpecArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	NodeStorage          int
	OwnerAccount         string
	OwnerId              int64
	NodeClass            string
	SecurityToken        string
	DBInstanceId         string
	NodeId               string
}
type ModifyNodeSpecResponse struct {
	RequestId string
	OrderId   string
}

func (c *DdsClient) ModifySecurityIps(req *ModifySecurityIpsArgs) (resp *ModifySecurityIpsResponse, err error) {
	resp = &ModifySecurityIpsResponse{}
	err = c.Request("ModifySecurityIps", req, resp)
	return
}

type ModifySecurityIpsArgs struct {
	ResourceOwnerId          int64
	ModifyMode               string
	ResourceOwnerAccount     string
	OwnerAccount             string
	SecurityIps              string
	OwnerId                  int64
	SecurityIpGroupName      string
	SecurityToken            string
	DBInstanceId             string
	SecurityIpGroupAttribute string
}
type ModifySecurityIpsResponse struct {
	RequestId string
}

func (c *DdsClient) DescribeSecurityIps(req *DescribeSecurityIpsArgs) (resp *DescribeSecurityIpsResponse, err error) {
	resp = &DescribeSecurityIpsResponse{}
	err = c.Request("DescribeSecurityIps", req, resp)
	return
}

type DescribeSecurityIpsSecurityIpGroup struct {
	SecurityIpGroupName      string
	SecurityIpGroupAttribute string
	SecurityIpList           string
}
type DescribeSecurityIpsArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeSecurityIpsResponse struct {
	RequestId        string
	SecurityIps      string
	SecurityIpGroups DescribeSecurityIpsSecurityIpGroupList
}

type DescribeSecurityIpsSecurityIpGroupList []DescribeSecurityIpsSecurityIpGroup

func (list *DescribeSecurityIpsSecurityIpGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityIpsSecurityIpGroup)
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

func (c *DdsClient) DescribeReplicaInitializeProgress(req *DescribeReplicaInitializeProgressArgs) (resp *DescribeReplicaInitializeProgressResponse, err error) {
	resp = &DescribeReplicaInitializeProgressResponse{}
	err = c.Request("DescribeReplicaInitializeProgress", req, resp)
	return
}

type DescribeReplicaInitializeProgressItemsItem struct {
	ReplicaId   string
	Status      string
	Progress    string
	FinishTime  string
	CurrentStep string
}
type DescribeReplicaInitializeProgressArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	ReplicaId            string
	OwnerId              int64
}
type DescribeReplicaInitializeProgressResponse struct {
	RequestId string
	Items     DescribeReplicaInitializeProgressItemsItemList
}

type DescribeReplicaInitializeProgressItemsItemList []DescribeReplicaInitializeProgressItemsItem

func (list *DescribeReplicaInitializeProgressItemsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicaInitializeProgressItemsItem)
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

func (c *DdsClient) ModifyDBInstanceNetworkType(req *ModifyDBInstanceNetworkTypeArgs) (resp *ModifyDBInstanceNetworkTypeResponse, err error) {
	resp = &ModifyDBInstanceNetworkTypeResponse{}
	err = c.Request("ModifyDBInstanceNetworkType", req, resp)
	return
}

type ModifyDBInstanceNetworkTypeArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	NetworkType          string
	OwnerId              int64
	VSwitchId            string
	SecurityToken        string
	RetainClassic        string
	ClassicExpiredDays   int
	VpcId                string
	DBInstanceId         string
}
type ModifyDBInstanceNetworkTypeResponse struct {
	RequestId string
}

func (c *DdsClient) ModifyDBInstanceNetExpireTime(req *ModifyDBInstanceNetExpireTimeArgs) (resp *ModifyDBInstanceNetExpireTimeResponse, err error) {
	resp = &ModifyDBInstanceNetExpireTimeResponse{}
	err = c.Request("ModifyDBInstanceNetExpireTime", req, resp)
	return
}

type ModifyDBInstanceNetExpireTimeArgs struct {
	ResourceOwnerId          int64
	SecurityToken            string
	ResourceOwnerAccount     string
	ConnectionString         string
	OwnerAccount             string
	DBInstanceId             string
	OwnerId                  int64
	ClassicExpendExpiredDays int
}
type ModifyDBInstanceNetExpireTimeResponse struct {
	RequestId string
}

func (c *DdsClient) DescribeAuditRecords(req *DescribeAuditRecordsArgs) (resp *DescribeAuditRecordsResponse, err error) {
	resp = &DescribeAuditRecordsResponse{}
	err = c.Request("DescribeAuditRecords", req, resp)
	return
}

type DescribeAuditRecordsSQLRecord struct {
	DBName              string
	AccountName         string
	HostAddress         string
	Syntax              string
	TotalExecutionTimes int64
	ReturnRowCounts     int64
	ExecuteTime         string
	ThreadID            string
}
type DescribeAuditRecordsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	EndTime              string
	StartTime            string
	OwnerId              int64
	QueryKeywords        string
	PageNumber           int
	Database             string
	Form                 string
	SecurityToken        string
	PageSize             int
	DBInstanceId         string
	NodeId               string
	User                 string
}
type DescribeAuditRecordsResponse struct {
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeAuditRecordsSQLRecordList
}

type DescribeAuditRecordsSQLRecordList []DescribeAuditRecordsSQLRecord

func (list *DescribeAuditRecordsSQLRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAuditRecordsSQLRecord)
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

func (c *DdsClient) DescribeReplicaSetRole(req *DescribeReplicaSetRoleArgs) (resp *DescribeReplicaSetRoleResponse, err error) {
	resp = &DescribeReplicaSetRoleResponse{}
	err = c.Request("DescribeReplicaSetRole", req, resp)
	return
}

type DescribeReplicaSetRoleReplicaSet struct {
	ReplicaSetRole   string
	ConnectionDomain string
	ConnectionPort   string
	ExpiredTime      string
	NetworkType      string
}
type DescribeReplicaSetRoleArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeReplicaSetRoleResponse struct {
	RequestId    string
	DBInstanceId string
	ReplicaSets  DescribeReplicaSetRoleReplicaSetList
}

type DescribeReplicaSetRoleReplicaSetList []DescribeReplicaSetRoleReplicaSet

func (list *DescribeReplicaSetRoleReplicaSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicaSetRoleReplicaSet)
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

func (c *DdsClient) CreateNode(req *CreateNodeArgs) (resp *CreateNodeResponse, err error) {
	resp = &CreateNodeResponse{}
	err = c.Request("CreateNode", req, resp)
	return
}

type CreateNodeArgs struct {
	ResourceOwnerId      int64
	NodeType             string
	ResourceOwnerAccount string
	ClientToken          string
	NodeStorage          int
	OwnerAccount         string
	OwnerId              int64
	NodeClass            string
	SecurityToken        string
	DBInstanceId         string
}
type CreateNodeResponse struct {
	RequestId string
	OrderId   string
}

func (c *DdsClient) ModifyDBInstanceDescription(req *ModifyDBInstanceDescriptionArgs) (resp *ModifyDBInstanceDescriptionResponse, err error) {
	resp = &ModifyDBInstanceDescriptionResponse{}
	err = c.Request("ModifyDBInstanceDescription", req, resp)
	return
}

type ModifyDBInstanceDescriptionArgs struct {
	ResourceOwnerId       int64
	SecurityToken         string
	ResourceOwnerAccount  string
	OwnerAccount          string
	DBInstanceId          string
	DBInstanceDescription string
	OwnerId               int64
	NodeId                string
}
type ModifyDBInstanceDescriptionResponse struct {
	RequestId string
}

func (c *DdsClient) DeleteNode(req *DeleteNodeArgs) (resp *DeleteNodeResponse, err error) {
	resp = &DeleteNodeResponse{}
	err = c.Request("DeleteNode", req, resp)
	return
}

type DeleteNodeArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
	NodeId               string
}
type DeleteNodeResponse struct {
	RequestId string
	TaskId    int
}

func (c *DdsClient) ReleaseReplica(req *ReleaseReplicaArgs) (resp *ReleaseReplicaResponse, err error) {
	resp = &ReleaseReplicaResponse{}
	err = c.Request("ReleaseReplica", req, resp)
	return
}

type ReleaseReplicaArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	ReplicaId            string
	OwnerId              int64
}
type ReleaseReplicaResponse struct {
	RequestId string
}

func (c *DdsClient) DescribeReplicas(req *DescribeReplicasArgs) (resp *DescribeReplicasResponse, err error) {
	resp = &DescribeReplicasResponse{}
	err = c.Request("DescribeReplicas", req, resp)
	return
}

type DescribeReplicasItems struct {
	ReplicaId          string
	ReplicaDescription string
	ReplicaStatus      string
	DBInstances        DescribeReplicasItems1List
}

type DescribeReplicasItems1 struct {
	DBInstanceId string
	Role         string
}
type DescribeReplicasArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	ReplicaId            string
	PageSize             int
	OwnerId              int64
	PageNumber           int
}
type DescribeReplicasResponse struct {
	RequestId        string
	PageNumber       string
	TotalRecordCount int
	PageRecordCount  int
	Replicas         DescribeReplicasItemsList
}

type DescribeReplicasItems1List []DescribeReplicasItems1

func (list *DescribeReplicasItems1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicasItems1)
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

type DescribeReplicasItemsList []DescribeReplicasItems

func (list *DescribeReplicasItemsList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicasItems)
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

func (c *DdsClient) RestoreDBInstance(req *RestoreDBInstanceArgs) (resp *RestoreDBInstanceResponse, err error) {
	resp = &RestoreDBInstanceResponse{}
	err = c.Request("RestoreDBInstance", req, resp)
	return
}

type RestoreDBInstanceArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	BackupId             int
	DBInstanceId         string
	OwnerId              int64
}
type RestoreDBInstanceResponse struct {
	RequestId string
}

func (c *DdsClient) CreateShardingDBInstance(req *CreateShardingDBInstanceArgs) (resp *CreateShardingDBInstanceResponse, err error) {
	resp = &CreateShardingDBInstanceResponse{}
	err = c.Request("CreateShardingDBInstance", req, resp)
	return
}

type CreateShardingDBInstanceReplicaSet struct {
	_class  string
	Storage int
}

type CreateShardingDBInstanceConfigServer struct {
	_class  string
	Storage int
}

type CreateShardingDBInstanceMongos struct {
	_class string
}
type CreateShardingDBInstanceArgs struct {
	ResourceOwnerId       int64
	ClientToken           string
	EngineVersion         string
	NetworkType           string
	ReplicaSets           CreateShardingDBInstanceReplicaSetList
	StorageEngine         string
	SecurityToken         string
	Engine                string
	DBInstanceDescription string
	Period                int
	RestoreTime           string
	ResourceOwnerAccount  string
	SrcDBInstanceId       string
	OwnerAccount          string
	ConfigServers         CreateShardingDBInstanceConfigServerList
	OwnerId               int64
	Mongoss               CreateShardingDBInstanceMongosList
	SecurityIPList        string
	VSwitchId             string
	AccountPassword       string
	VpcId                 string
	ZoneId                string
	ChargeType            string
}
type CreateShardingDBInstanceResponse struct {
	RequestId    string
	OrderId      string
	DBInstanceId string
}

type CreateShardingDBInstanceReplicaSetList []CreateShardingDBInstanceReplicaSet

func (list *CreateShardingDBInstanceReplicaSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateShardingDBInstanceReplicaSet)
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

type CreateShardingDBInstanceConfigServerList []CreateShardingDBInstanceConfigServer

func (list *CreateShardingDBInstanceConfigServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateShardingDBInstanceConfigServer)
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

type CreateShardingDBInstanceMongosList []CreateShardingDBInstanceMongos

func (list *CreateShardingDBInstanceMongosList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateShardingDBInstanceMongos)
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

func (c *DdsClient) ModifyBackupPolicy(req *ModifyBackupPolicyArgs) (resp *ModifyBackupPolicyResponse, err error) {
	resp = &ModifyBackupPolicyResponse{}
	err = c.Request("ModifyBackupPolicy", req, resp)
	return
}

type ModifyBackupPolicyArgs struct {
	PreferredBackupTime   string
	PreferredBackupPeriod string
	ResourceOwnerId       int64
	SecurityToken         string
	ResourceOwnerAccount  string
	OwnerAccount          string
	DBInstanceId          string
	OwnerId               int64
}
type ModifyBackupPolicyResponse struct {
	RequestId string
}

func (c *DdsClient) DescribeAccounts(req *DescribeAccountsArgs) (resp *DescribeAccountsResponse, err error) {
	resp = &DescribeAccountsResponse{}
	err = c.Request("DescribeAccounts", req, resp)
	return
}

type DescribeAccountsAccount struct {
	DBInstanceId       string
	AccountName        string
	AccountStatus      string
	AccountDescription string
}
type DescribeAccountsArgs struct {
	ResourceOwnerId      int64
	AccountName          string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeAccountsResponse struct {
	RequestId string
	Accounts  DescribeAccountsAccountList
}

type DescribeAccountsAccountList []DescribeAccountsAccount

func (list *DescribeAccountsAccountList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccountsAccount)
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

func (c *DdsClient) DescribeAuditFiles(req *DescribeAuditFilesArgs) (resp *DescribeAuditFilesResponse, err error) {
	resp = &DescribeAuditFilesResponse{}
	err = c.Request("DescribeAuditFiles", req, resp)
	return
}

type DescribeAuditFilesLogFile struct {
	FileID         int
	LogStatus      string
	LogStartTime   string
	LogEndTime     string
	LogDownloadURL string
	LogSize        int64
}
type DescribeAuditFilesArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	DBInstanceId         string
	OwnerId              int64
	NodeId               string
	PageNumber           int
}
type DescribeAuditFilesResponse struct {
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	DBInstanceId     string
	Items            DescribeAuditFilesLogFileList
}

type DescribeAuditFilesLogFileList []DescribeAuditFilesLogFile

func (list *DescribeAuditFilesLogFileList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAuditFilesLogFile)
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

func (c *DdsClient) DescribeReplicaPerformance(req *DescribeReplicaPerformanceArgs) (resp *DescribeReplicaPerformanceResponse, err error) {
	resp = &DescribeReplicaPerformanceResponse{}
	err = c.Request("DescribeReplicaPerformance", req, resp)
	return
}

type DescribeReplicaPerformancePerformanceKeys struct {
	PerformanceKey DescribeReplicaPerformancePerformanceKeyItemList
}

type DescribeReplicaPerformancePerformanceKeyItem struct {
	Key               string
	Unit              string
	ValueFormat       string
	PerformanceValues DescribeReplicaPerformancePerformanceValues
}

type DescribeReplicaPerformancePerformanceValues struct {
	PerformanceValue DescribeReplicaPerformancePerformanceValueItemList
}

type DescribeReplicaPerformancePerformanceValueItem struct {
	Value string
	Date  string
}
type DescribeReplicaPerformanceArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	EndTime              string
	StartTime            string
	OwnerId              int64
	SourceDBInstanceId   string
	SecurityToken        string
	ReplicaId            string
	Key                  string
}
type DescribeReplicaPerformanceResponse struct {
	RequestId       string
	StartTime       string
	EndTime         string
	ReplicaId       string
	PerformanceKeys DescribeReplicaPerformancePerformanceKeys
}

type DescribeReplicaPerformancePerformanceKeyItemList []DescribeReplicaPerformancePerformanceKeyItem

func (list *DescribeReplicaPerformancePerformanceKeyItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicaPerformancePerformanceKeyItem)
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

type DescribeReplicaPerformancePerformanceValueItemList []DescribeReplicaPerformancePerformanceValueItem

func (list *DescribeReplicaPerformancePerformanceValueItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicaPerformancePerformanceValueItem)
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

func (c *DdsClient) DescribeRegions(req *DescribeRegionsArgs) (resp *DescribeRegionsResponse, err error) {
	resp = &DescribeRegionsResponse{}
	err = c.Request("DescribeRegions", req, resp)
	return
}

type DescribeRegionsDdsRegion struct {
	RegionId string
	ZoneIds  string
	Zones    DescribeRegionsZoneList
}

type DescribeRegionsZone struct {
	ZoneId     string
	VpcEnabled core.Bool
}
type DescribeRegionsArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeRegionsResponse struct {
	RequestId string
	Regions   DescribeRegionsDdsRegionList
}

type DescribeRegionsZoneList []DescribeRegionsZone

func (list *DescribeRegionsZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsZone)
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

type DescribeRegionsDdsRegionList []DescribeRegionsDdsRegion

func (list *DescribeRegionsDdsRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsDdsRegion)
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

func (c *DdsClient) DescribeDBInstancePerformance(req *DescribeDBInstancePerformanceArgs) (resp *DescribeDBInstancePerformanceResponse, err error) {
	resp = &DescribeDBInstancePerformanceResponse{}
	err = c.Request("DescribeDBInstancePerformance", req, resp)
	return
}

type DescribeDBInstancePerformancePerformanceKey struct {
	Key               string
	Unit              string
	ValueFormat       string
	PerformanceValues DescribeDBInstancePerformancePerformanceValueList
}

type DescribeDBInstancePerformancePerformanceValue struct {
	Value string
	Date  string
}
type DescribeDBInstancePerformanceArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	EndTime              string
	StartTime            string
	OwnerId              int64
	ReplicaSetRole       string
	SecurityToken        string
	DBInstanceId         string
	NodeId               string
	Key                  string
}
type DescribeDBInstancePerformanceResponse struct {
	RequestId       string
	DBInstanceId    string
	Engine          string
	StartTime       string
	EndTime         string
	PerformanceKeys DescribeDBInstancePerformancePerformanceKeyList
}

type DescribeDBInstancePerformancePerformanceValueList []DescribeDBInstancePerformancePerformanceValue

func (list *DescribeDBInstancePerformancePerformanceValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancePerformancePerformanceValue)
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

type DescribeDBInstancePerformancePerformanceKeyList []DescribeDBInstancePerformancePerformanceKey

func (list *DescribeDBInstancePerformancePerformanceKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancePerformancePerformanceKey)
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

func (c *DdsClient) DeleteDBInstance(req *DeleteDBInstanceArgs) (resp *DeleteDBInstanceResponse, err error) {
	resp = &DeleteDBInstanceResponse{}
	err = c.Request("DeleteDBInstance", req, resp)
	return
}

type DeleteDBInstanceArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DeleteDBInstanceResponse struct {
	RequestId string
}

func (c *DdsClient) Sample(req *SampleArgs) (resp *SampleResponse, err error) {
	resp = &SampleResponse{}
	err = c.Request("Sample", req, resp)
	return
}

type SampleSecurityIpGroup struct {
	SecurityIpGroupName      string
	SecurityIpGroupAttribute string
	SecurityIpList           string
}
type SampleArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type SampleResponse struct {
	RequestId        string
	SecurityIps      string
	SecurityIpGroups SampleSecurityIpGroupList
}

type SampleSecurityIpGroupList []SampleSecurityIpGroup

func (list *SampleSecurityIpGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SampleSecurityIpGroup)
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

func (c *DdsClient) DescribeDBInstanceAttribute(req *DescribeDBInstanceAttributeArgs) (resp *DescribeDBInstanceAttributeResponse, err error) {
	resp = &DescribeDBInstanceAttributeResponse{}
	err = c.Request("DescribeDBInstanceAttribute", req, resp)
	return
}

type DescribeDBInstanceAttributeDBInstance struct {
	DBInstanceId          string
	DBInstanceDescription string
	RegionId              string
	ZoneId                string
	Engine                string
	EngineVersion         string
	StorageEngine         string
	DBInstanceClass       string
	DBInstanceStorage     int
	DBInstanceStatus      string
	LockMode              string
	ChargeType            string
	CreationTime          string
	ReplicaSetName        string
	NetworkType           string
	ExpireTime            string
	MaintainStartTime     string
	MaintainEndTime       string
	DBInstanceType        string
	LastDowngradeTime     int
	MongosList            DescribeDBInstanceAttributeMongosAttributeList
	ShardList             DescribeDBInstanceAttributeShardAttributeList
}

type DescribeDBInstanceAttributeMongosAttribute struct {
	NodeId          string
	NodeDescription string
	NodeClass       string
	ConnectSting    string
	Port            int
}

type DescribeDBInstanceAttributeShardAttribute struct {
	NodeId          string
	NodeDescription string
	NodeClass       string
	NodeStorage     int
}
type DescribeDBInstanceAttributeArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	Engine               string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeDBInstanceAttributeResponse struct {
	RequestId   string
	DBInstances DescribeDBInstanceAttributeDBInstanceList
}

type DescribeDBInstanceAttributeMongosAttributeList []DescribeDBInstanceAttributeMongosAttribute

func (list *DescribeDBInstanceAttributeMongosAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceAttributeMongosAttribute)
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

type DescribeDBInstanceAttributeShardAttributeList []DescribeDBInstanceAttributeShardAttribute

func (list *DescribeDBInstanceAttributeShardAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceAttributeShardAttribute)
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

type DescribeDBInstanceAttributeDBInstanceList []DescribeDBInstanceAttributeDBInstance

func (list *DescribeDBInstanceAttributeDBInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceAttributeDBInstance)
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
