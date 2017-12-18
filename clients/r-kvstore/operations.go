package r_kvstore

import "encoding/json"

func (c *RKvstoreClient) ModifyInstanceSpec(req *ModifyInstanceSpecArgs) (resp *ModifyInstanceSpecResponse, err error) {
	resp = &ModifyInstanceSpecResponse{}
	err = c.Request("ModifyInstanceSpec", req, resp)
	return
}

type ModifyInstanceSpecArgs struct {
	ResourceOwnerId      int64
	FromApp              string
	ResourceOwnerAccount string
	OwnerAccount         string
	CouponNo             string
	OwnerId              int64
	InstanceClass        string
	InstanceId           string
	SecurityToken        string
	ForceUpgrade         bool
	BusinessInfo         string
}
type ModifyInstanceSpecResponse struct {
	RequestId string
	OrderId   string
}

func (c *RKvstoreClient) DescribeBackupPolicy(req *DescribeBackupPolicyArgs) (resp *DescribeBackupPolicyResponse, err error) {
	resp = &DescribeBackupPolicyResponse{}
	err = c.Request("DescribeBackupPolicy", req, resp)
	return
}

type DescribeBackupPolicyArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeBackupPolicyResponse struct {
	RequestId               string
	BackupRetentionPeriod   string
	PreferredBackupTime     string
	PreferredBackupPeriod   string
	PreferredNextBackupTime string
}

func (c *RKvstoreClient) DeleteSnapshotSettings(req *DeleteSnapshotSettingsArgs) (resp *DeleteSnapshotSettingsResponse, err error) {
	resp = &DeleteSnapshotSettingsResponse{}
	err = c.Request("DeleteSnapshotSettings", req, resp)
	return
}

type DeleteSnapshotSettingsArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DeleteSnapshotSettingsResponse struct {
	RequestId string
}

func (c *RKvstoreClient) DescribeReplicaUsage(req *DescribeReplicaUsageArgs) (resp *DescribeReplicaUsageResponse, err error) {
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

func (c *RKvstoreClient) ModifyReplicaDescription(req *ModifyReplicaDescriptionArgs) (resp *ModifyReplicaDescriptionResponse, err error) {
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

func (c *RKvstoreClient) ModifyInstanceAttribute(req *ModifyInstanceAttributeArgs) (resp *ModifyInstanceAttributeResponse, err error) {
	resp = &ModifyInstanceAttributeResponse{}
	err = c.Request("ModifyInstanceAttribute", req, resp)
	return
}

type ModifyInstanceAttributeArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	InstanceName         string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	NewPassword          string
}
type ModifyInstanceAttributeResponse struct {
	RequestId string
}

func (c *RKvstoreClient) ModifyInstanceConfig(req *ModifyInstanceConfigArgs) (resp *ModifyInstanceConfigResponse, err error) {
	resp = &ModifyInstanceConfigResponse{}
	err = c.Request("ModifyInstanceConfig", req, resp)
	return
}

type ModifyInstanceConfigArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Config               string
}
type ModifyInstanceConfigResponse struct {
	RequestId string
}

func (c *RKvstoreClient) DescribeInstanceAttribute(req *DescribeInstanceAttributeArgs) (resp *DescribeInstanceAttributeResponse, err error) {
	resp = &DescribeInstanceAttributeResponse{}
	err = c.Request("DescribeInstanceAttribute", req, resp)
	return
}

type DescribeInstanceAttributeDBInstanceAttribute struct {
	InstanceId          string
	InstanceName        string
	ConnectionDomain    string
	Port                int64
	InstanceStatus      string
	RegionId            string
	Capacity            int64
	InstanceClass       string
	QPS                 int64
	Bandwidth           int64
	Connections         int64
	ZoneId              string
	Config              string
	ChargeType          string
	NodeType            string
	NetworkType         string
	VpcId               string
	VSwitchId           string
	PrivateIp           string
	CreateTime          string
	EndTime             string
	HasRenewChangeOrder string
	IsRds               bool
	Engine              string
	EngineVersion       string
	MaintainStartTime   string
	MaintainEndTime     string
	AvailabilityValue   string
	SecurityIPList      string
	InstanceType        string
	ArchitectureType    string
	NodeType1           string
	PackageType         string
}
type DescribeInstanceAttributeArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeInstanceAttributeResponse struct {
	RequestId string
	Instances DescribeInstanceAttributeDBInstanceAttributeList
}

type DescribeInstanceAttributeDBInstanceAttributeList []DescribeInstanceAttributeDBInstanceAttribute

func (list *DescribeInstanceAttributeDBInstanceAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceAttributeDBInstanceAttribute)
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

func (c *RKvstoreClient) CreateSnapshot(req *CreateSnapshotArgs) (resp *CreateSnapshotResponse, err error) {
	resp = &CreateSnapshotResponse{}
	err = c.Request("CreateSnapshot", req, resp)
	return
}

type CreateSnapshotArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	SnapshotName         string
	OwnerId              int64
}
type CreateSnapshotResponse struct {
	RequestId  string
	SnapshotId string
	Status     string
}

func (c *RKvstoreClient) RenewInstance(req *RenewInstanceArgs) (resp *RenewInstanceResponse, err error) {
	resp = &RenewInstanceResponse{}
	err = c.Request("RenewInstance", req, resp)
	return
}

type RenewInstanceArgs struct {
	ResourceOwnerId      int64
	Period               int64
	AutoPay              bool
	FromApp              string
	ResourceOwnerAccount string
	OwnerAccount         string
	CouponNo             string
	OwnerId              int64
	InstanceClass        string
	Capacity             string
	InstanceId           string
	SecurityToken        string
	ForceUpgrade         bool
	BusinessInfo         string
}
type RenewInstanceResponse struct {
	RequestId string
	OrderId   string
	EndTime   string
}

func (c *RKvstoreClient) ModifyInstanceMinorVersion(req *ModifyInstanceMinorVersionArgs) (resp *ModifyInstanceMinorVersionResponse, err error) {
	resp = &ModifyInstanceMinorVersionResponse{}
	err = c.Request("ModifyInstanceMinorVersion", req, resp)
	return
}

type ModifyInstanceMinorVersionArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	SecurityToken        string
	ResourceOwnerAccount string
	Minorversion         string
	OwnerAccount         string
	OwnerId              int64
}
type ModifyInstanceMinorVersionResponse struct {
	RequestId string
}

func (c *RKvstoreClient) CreateBackup(req *CreateBackupArgs) (resp *CreateBackupResponse, err error) {
	resp = &CreateBackupResponse{}
	err = c.Request("CreateBackup", req, resp)
	return
}

type CreateBackupArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type CreateBackupResponse struct {
	RequestId   string
	BackupJobID string
}

func (c *RKvstoreClient) RestoreSnapshot(req *RestoreSnapshotArgs) (resp *RestoreSnapshotResponse, err error) {
	resp = &RestoreSnapshotResponse{}
	err = c.Request("RestoreSnapshot", req, resp)
	return
}

type RestoreSnapshotArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	SnapshotId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type RestoreSnapshotResponse struct {
	RequestId string
}

func (c *RKvstoreClient) DescribeCertification(req *DescribeCertificationArgs) (resp *DescribeCertificationResponse, err error) {
	resp = &DescribeCertificationResponse{}
	err = c.Request("DescribeCertification", req, resp)
	return
}

type DescribeCertificationArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Parameters           string
}
type DescribeCertificationResponse struct {
	RequestId       string
	NoCertification bool
}

func (c *RKvstoreClient) DescribeBackups(req *DescribeBackupsArgs) (resp *DescribeBackupsResponse, err error) {
	resp = &DescribeBackupsResponse{}
	err = c.Request("DescribeBackups", req, resp)
	return
}

type DescribeBackupsBackup struct {
	BackupId          int
	BackupDBNames     string
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
	InstanceId           string
	SecurityToken        string
	PageSize             int
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

func (c *RKvstoreClient) DeleteSnapshot(req *DeleteSnapshotArgs) (resp *DeleteSnapshotResponse, err error) {
	resp = &DeleteSnapshotResponse{}
	err = c.Request("DeleteSnapshot", req, resp)
	return
}

type DeleteSnapshotArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	SnapshotId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DeleteSnapshotResponse struct {
	RequestId string
}

func (c *RKvstoreClient) ModifySecurityIps(req *ModifySecurityIpsArgs) (resp *ModifySecurityIpsResponse, err error) {
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
	InstanceId               string
	SecurityToken            string
	SecurityIpGroupAttribute string
}
type ModifySecurityIpsResponse struct {
	RequestId string
}

func (c *RKvstoreClient) DescribeSecurityIps(req *DescribeSecurityIpsArgs) (resp *DescribeSecurityIpsResponse, err error) {
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
	InstanceId           string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeSecurityIpsResponse struct {
	RequestId        string
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

func (c *RKvstoreClient) DescribeReplicaInitializeProgress(req *DescribeReplicaInitializeProgressArgs) (resp *DescribeReplicaInitializeProgressResponse, err error) {
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

func (c *RKvstoreClient) DescribeMonitorItems(req *DescribeMonitorItemsArgs) (resp *DescribeMonitorItemsResponse, err error) {
	resp = &DescribeMonitorItemsResponse{}
	err = c.Request("DescribeMonitorItems", req, resp)
	return
}

type DescribeMonitorItemsKVStoreMonitorItem struct {
	MonitorKey string
	Unit       string
}
type DescribeMonitorItemsArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeMonitorItemsResponse struct {
	RequestId    string
	MonitorItems DescribeMonitorItemsKVStoreMonitorItemList
}

type DescribeMonitorItemsKVStoreMonitorItemList []DescribeMonitorItemsKVStoreMonitorItem

func (list *DescribeMonitorItemsKVStoreMonitorItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMonitorItemsKVStoreMonitorItem)
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

func (c *RKvstoreClient) DescribeInstances(req *DescribeInstancesArgs) (resp *DescribeInstancesResponse, err error) {
	resp = &DescribeInstancesResponse{}
	err = c.Request("DescribeInstances", req, resp)
	return
}

type DescribeInstancesKVStoreInstance struct {
	InstanceId          string
	InstanceName        string
	ConnectionDomain    string
	Port                int64
	UserName            string
	InstanceStatus      string
	RegionId            string
	Capacity            int64
	InstanceClass       string
	QPS                 int64
	Bandwidth           int64
	Connections         int64
	ZoneId              string
	Config              string
	ChargeType          string
	NetworkType         string
	VpcId               string
	VSwitchId           string
	PrivateIp           string
	CreateTime          string
	EndTime             string
	HasRenewChangeOrder string
	IsRds               bool
	InstanceType        string
	ArchitectureType    string
	NodeType            string
	PackageType         string
}
type DescribeInstancesArgs struct {
	ResourceOwnerId      int64
	InstanceStatus       string
	ResourceOwnerAccount string
	OwnerAccount         string
	NetworkType          string
	OwnerId              int64
	PageNumber           int
	VSwitchId            string
	SecurityToken        string
	InstanceIds          string
	VpcId                string
	PageSize             int
	InstanceType         string
	ChargeType           string
}
type DescribeInstancesResponse struct {
	RequestId  string
	PageNumber int
	PageSize   int
	TotalCount int
	Instances  DescribeInstancesKVStoreInstanceList
}

type DescribeInstancesKVStoreInstanceList []DescribeInstancesKVStoreInstance

func (list *DescribeInstancesKVStoreInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesKVStoreInstance)
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

func (c *RKvstoreClient) ModifyInstanceNetExpireTime(req *ModifyInstanceNetExpireTimeArgs) (resp *ModifyInstanceNetExpireTimeResponse, err error) {
	resp = &ModifyInstanceNetExpireTimeResponse{}
	err = c.Request("ModifyInstanceNetExpireTime", req, resp)
	return
}

type ModifyInstanceNetExpireTimeNetInfoItem struct {
	DBInstanceNetType string
	Port              string
	ExpiredTime       string
	ConnectionString  string
	IPAddress         string
}
type ModifyInstanceNetExpireTimeArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	SecurityToken        string
	ResourceOwnerAccount string
	ConnectionString     string
	ClassicExpiredDays   int
	OwnerAccount         string
	OwnerId              int64
}
type ModifyInstanceNetExpireTimeResponse struct {
	RequestId    string
	InstanceId   string
	NetInfoItems ModifyInstanceNetExpireTimeNetInfoItemList
}

type ModifyInstanceNetExpireTimeNetInfoItemList []ModifyInstanceNetExpireTimeNetInfoItem

func (list *ModifyInstanceNetExpireTimeNetInfoItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyInstanceNetExpireTimeNetInfoItem)
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

func (c *RKvstoreClient) DescribeTempInstance(req *DescribeTempInstanceArgs) (resp *DescribeTempInstanceResponse, err error) {
	resp = &DescribeTempInstanceResponse{}
	err = c.Request("DescribeTempInstance", req, resp)
	return
}

type DescribeTempInstanceTempInstance struct {
	InstanceId     string
	TempInstanceId string
	SnapshotId     string
	CreateTime     string
	Domain         string
	Status         string
	Memory         int64
	ExpireTime     string
}
type DescribeTempInstanceArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeTempInstanceResponse struct {
	RequestId     string
	TempInstances DescribeTempInstanceTempInstanceList
}

type DescribeTempInstanceTempInstanceList []DescribeTempInstanceTempInstance

func (list *DescribeTempInstanceTempInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTempInstanceTempInstance)
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

func (c *RKvstoreClient) QueryTask(req *QueryTaskArgs) (resp *QueryTaskResponse, err error) {
	resp = &QueryTaskResponse{}
	err = c.Request("QueryTask", req, resp)
	return
}

type QueryTaskArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type QueryTaskResponse struct {
	RequestId string
	Action    string
	Progress  int
}

func (c *RKvstoreClient) DescribeMonthlyServiceStatus(req *DescribeMonthlyServiceStatusArgs) (resp *DescribeMonthlyServiceStatusResponse, err error) {
	resp = &DescribeMonthlyServiceStatusResponse{}
	err = c.Request("DescribeMonthlyServiceStatus", req, resp)
	return
}

type DescribeMonthlyServiceStatusInstanceSLAInfo struct {
	InstanceId string
	UptimePct  float32
}
type DescribeMonthlyServiceStatusArgs struct {
	ResourceOwnerId      int64
	Month                string
	SecurityToken        string
	ResourceOwnerAccount string
	InstanceIds          string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeMonthlyServiceStatusResponse struct {
	RequestId        string
	TotalCount       int64
	InstanceSLAInfos DescribeMonthlyServiceStatusInstanceSLAInfoList
}

type DescribeMonthlyServiceStatusInstanceSLAInfoList []DescribeMonthlyServiceStatusInstanceSLAInfo

func (list *DescribeMonthlyServiceStatusInstanceSLAInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMonthlyServiceStatusInstanceSLAInfo)
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

func (c *RKvstoreClient) TransformToPrePaid(req *TransformToPrePaidArgs) (resp *TransformToPrePaidResponse, err error) {
	resp = &TransformToPrePaidResponse{}
	err = c.Request("TransformToPrePaid", req, resp)
	return
}

type TransformToPrePaidArgs struct {
	ResourceOwnerId      int64
	Period               int64
	InstanceId           string
	AutoPay              bool
	FromApp              string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type TransformToPrePaidResponse struct {
	RequestId string
	OrderId   string
	EndTime   string
}

func (c *RKvstoreClient) DeleteTempInstance(req *DeleteTempInstanceArgs) (resp *DeleteTempInstanceResponse, err error) {
	resp = &DeleteTempInstanceResponse{}
	err = c.Request("DeleteTempInstance", req, resp)
	return
}

type DeleteTempInstanceArgs struct {
	ResourceOwnerId      int64
	TempInstanceId       string
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DeleteTempInstanceResponse struct {
	RequestId string
}

func (c *RKvstoreClient) FlushInstance(req *FlushInstanceArgs) (resp *FlushInstanceResponse, err error) {
	resp = &FlushInstanceResponse{}
	err = c.Request("FlushInstance", req, resp)
	return
}

type FlushInstanceArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type FlushInstanceResponse struct {
	RequestId string
}

func (c *RKvstoreClient) ReleaseReplica(req *ReleaseReplicaArgs) (resp *ReleaseReplicaResponse, err error) {
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

func (c *RKvstoreClient) SetSnapshotSettings(req *SetSnapshotSettingsArgs) (resp *SetSnapshotSettingsResponse, err error) {
	resp = &SetSnapshotSettingsResponse{}
	err = c.Request("SetSnapshotSettings", req, resp)
	return
}

type SetSnapshotSettingsArgs struct {
	ResourceOwnerId      int64
	EndHour              int
	ResourceOwnerAccount string
	OwnerAccount         string
	DayList              int
	OwnerId              int64
	InstanceId           string
	RetentionDay         int
	MaxManualSnapshots   int
	MaxAutoSnapshots     int
	BeginHour            int
}
type SetSnapshotSettingsResponse struct {
	RequestId string
}

func (c *RKvstoreClient) DescribeReplicas(req *DescribeReplicasArgs) (resp *DescribeReplicasResponse, err error) {
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

func (c *RKvstoreClient) DescribeDBInstanceNetInfo(req *DescribeDBInstanceNetInfoArgs) (resp *DescribeDBInstanceNetInfoResponse, err error) {
	resp = &DescribeDBInstanceNetInfoResponse{}
	err = c.Request("DescribeDBInstanceNetInfo", req, resp)
	return
}

type DescribeDBInstanceNetInfoInstanceNetInfo struct {
	ConnectionString  string
	IPAddress         string
	Port              string
	VPCId             string
	VSwitchId         string
	DBInstanceNetType string
	ExpiredTime       string
	Upgradeable       string
}
type DescribeDBInstanceNetInfoArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeDBInstanceNetInfoResponse struct {
	RequestId           string
	InstanceNetworkType string
	NetInfoItems        DescribeDBInstanceNetInfoInstanceNetInfoList
}

type DescribeDBInstanceNetInfoInstanceNetInfoList []DescribeDBInstanceNetInfoInstanceNetInfo

func (list *DescribeDBInstanceNetInfoInstanceNetInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetInfoInstanceNetInfo)
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

func (c *RKvstoreClient) GetSnapshotSettings(req *GetSnapshotSettingsArgs) (resp *GetSnapshotSettingsResponse, err error) {
	resp = &GetSnapshotSettingsResponse{}
	err = c.Request("GetSnapshotSettings", req, resp)
	return
}

type GetSnapshotSettingsArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type GetSnapshotSettingsResponse struct {
	RequestId          string
	InstanceId         string
	BeginHour          int
	EndHour            int
	RetentionDay       int
	MaxAutoSnapshots   int
	MaxManualSnapshots int
	DayList            int
	NextTime           string
}

func (c *RKvstoreClient) CreateInstance(req *CreateInstanceArgs) (resp *CreateInstanceResponse, err error) {
	resp = &CreateInstanceResponse{}
	err = c.Request("CreateInstance", req, resp)
	return
}

type CreateInstanceArgs struct {
	ResourceOwnerId      int64
	NodeType             string
	CouponNo             string
	NetworkType          string
	InstanceClass        string
	Capacity             int64
	Password             string
	SecurityToken        string
	InstanceType         string
	BusinessInfo         string
	Period               string
	ResourceOwnerAccount string
	SrcDBInstanceId      string
	OwnerAccount         string
	BackupId             string
	OwnerId              int64
	Token                string
	VSwitchId            string
	InstanceName         string
	VpcId                string
	ZoneId               string
	ChargeType           string
	Config               string
}
type CreateInstanceResponse struct {
	RequestId        string
	InstanceId       string
	InstanceName     string
	ConnectionDomain string
	Port             int
	UserName         string
	InstanceStatus   string
	RegionId         string
	Capacity         int64
	QPS              int64
	Bandwidth        int64
	Connections      int64
	ZoneId           string
	Config           string
	ChargeType       string
	EndTime          string
	NodeType         string
	NetworkType      string
	VpcId            string
	VSwitchId        string
	PrivateIpAddr    string
}

func (c *RKvstoreClient) CreateTempInstance(req *CreateTempInstanceArgs) (resp *CreateTempInstanceResponse, err error) {
	resp = &CreateTempInstanceResponse{}
	err = c.Request("CreateTempInstance", req, resp)
	return
}

type CreateTempInstanceArgs struct {
	ResourceOwnerId      int64
	SnapshotId           string
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type CreateTempInstanceResponse struct {
	RequestId      string
	InstanceId     string
	SnapshotId     string
	TempInstanceId string
	Status         string
}

func (c *RKvstoreClient) RestoreInstance(req *RestoreInstanceArgs) (resp *RestoreInstanceResponse, err error) {
	resp = &RestoreInstanceResponse{}
	err = c.Request("RestoreInstance", req, resp)
	return
}

type RestoreInstanceArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	BackupId             string
	OwnerId              int64
}
type RestoreInstanceResponse struct {
	RequestId string
}

func (c *RKvstoreClient) ModifyInstanceMaintainTime(req *ModifyInstanceMaintainTimeArgs) (resp *ModifyInstanceMaintainTimeResponse, err error) {
	resp = &ModifyInstanceMaintainTimeResponse{}
	err = c.Request("ModifyInstanceMaintainTime", req, resp)
	return
}

type ModifyInstanceMaintainTimeArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	SecurityToken        string
	ResourceOwnerAccount string
	MaintainStartTime    string
	OwnerAccount         string
	OwnerId              int64
	MaintainEndTime      string
}
type ModifyInstanceMaintainTimeResponse struct {
	RequestId string
}

func (c *RKvstoreClient) DeleteInstance(req *DeleteInstanceArgs) (resp *DeleteInstanceResponse, err error) {
	resp = &DeleteInstanceResponse{}
	err = c.Request("DeleteInstance", req, resp)
	return
}

type DeleteInstanceArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DeleteInstanceResponse struct {
	RequestId string
}

func (c *RKvstoreClient) ModifyBackupPolicy(req *ModifyBackupPolicyArgs) (resp *ModifyBackupPolicyResponse, err error) {
	resp = &ModifyBackupPolicyResponse{}
	err = c.Request("ModifyBackupPolicy", req, resp)
	return
}

type ModifyBackupPolicyArgs struct {
	PreferredBackupTime   string
	PreferredBackupPeriod string
	ResourceOwnerId       int64
	InstanceId            string
	SecurityToken         string
	ResourceOwnerAccount  string
	OwnerAccount          string
	OwnerId               int64
}
type ModifyBackupPolicyResponse struct {
	RequestId string
}

func (c *RKvstoreClient) SwitchNetwork(req *SwitchNetworkArgs) (resp *SwitchNetworkResponse, err error) {
	resp = &SwitchNetworkResponse{}
	err = c.Request("SwitchNetwork", req, resp)
	return
}

type SwitchNetworkArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	VSwitchId            string
	InstanceId           string
	SecurityToken        string
	TargetNetworkType    string
	RetainClassic        string
	ClassicExpiredDays   string
	VpcId                string
}
type SwitchNetworkResponse struct {
	RequestId string
	TaskId    string
}

func (c *RKvstoreClient) ModifyCertification(req *ModifyCertificationArgs) (resp *ModifyCertificationResponse, err error) {
	resp = &ModifyCertificationResponse{}
	err = c.Request("ModifyCertification", req, resp)
	return
}

type ModifyCertificationArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	NoCertification      bool
}
type ModifyCertificationResponse struct {
	RequestId string
}

func (c *RKvstoreClient) DescribeHistoryMonitorValues(req *DescribeHistoryMonitorValuesArgs) (resp *DescribeHistoryMonitorValuesResponse, err error) {
	resp = &DescribeHistoryMonitorValuesResponse{}
	err = c.Request("DescribeHistoryMonitorValues", req, resp)
	return
}

type DescribeHistoryMonitorValuesArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	EndTime              string
	StartTime            string
	OwnerId              int64
	InstanceId           string
	SecurityToken        string
	IntervalForHistory   string
	MonitorKeys          string
}
type DescribeHistoryMonitorValuesResponse struct {
	RequestId      string
	MonitorHistory string
}

func (c *RKvstoreClient) DescribeInstanceConfig(req *DescribeInstanceConfigArgs) (resp *DescribeInstanceConfigResponse, err error) {
	resp = &DescribeInstanceConfigResponse{}
	err = c.Request("DescribeInstanceConfig", req, resp)
	return
}

type DescribeInstanceConfigArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeInstanceConfigResponse struct {
	RequestId string
	Config    string
}

func (c *RKvstoreClient) SwitchTempInstance(req *SwitchTempInstanceArgs) (resp *SwitchTempInstanceResponse, err error) {
	resp = &SwitchTempInstanceResponse{}
	err = c.Request("SwitchTempInstance", req, resp)
	return
}

type SwitchTempInstanceArgs struct {
	ResourceOwnerId      int64
	TempInstanceId       string
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type SwitchTempInstanceResponse struct {
	RequestId string
}

func (c *RKvstoreClient) ModifyInstanceSpecPreCheck(req *ModifyInstanceSpecPreCheckArgs) (resp *ModifyInstanceSpecPreCheckResponse, err error) {
	resp = &ModifyInstanceSpecPreCheckResponse{}
	err = c.Request("ModifyInstanceSpecPreCheck", req, resp)
	return
}

type ModifyInstanceSpecPreCheckArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	TargetInstanceClass  string
}
type ModifyInstanceSpecPreCheckResponse struct {
	RequestId       string
	IsAllowModify   bool
	DisableCommands string
}

func (c *RKvstoreClient) DescribeReplicaPerformance(req *DescribeReplicaPerformanceArgs) (resp *DescribeReplicaPerformanceResponse, err error) {
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

func (c *RKvstoreClient) RenewMultiInstance(req *RenewMultiInstanceArgs) (resp *RenewMultiInstanceResponse, err error) {
	resp = &RenewMultiInstanceResponse{}
	err = c.Request("RenewMultiInstance", req, resp)
	return
}

type RenewMultiInstanceArgs struct {
	ResourceOwnerId      int64
	Period               int64
	AutoPay              bool
	FromApp              string
	ResourceOwnerAccount string
	OwnerAccount         string
	CouponNo             string
	OwnerId              int64
	SecurityToken        string
	InstanceIds          string
	BusinessInfo         string
}
type RenewMultiInstanceResponse struct {
	RequestId string
	OrderId   string
}

func (c *RKvstoreClient) DescribeSnapshots(req *DescribeSnapshotsArgs) (resp *DescribeSnapshotsResponse, err error) {
	resp = &DescribeSnapshotsResponse{}
	err = c.Request("DescribeSnapshots", req, resp)
	return
}

type DescribeSnapshotsSnapshot struct {
	SnapshotId         string
	SnapshotName       string
	InstanceId         string
	CreateTime         string
	Memory             int64
	RdbSize            int64
	Status             string
	Type               string
	OssDownloadInPath  string
	OssDownloadOutPath string
}
type DescribeSnapshotsArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	SnapshotIds          string
	EndTime              string
	BeginTime            string
	OwnerId              int64
}
type DescribeSnapshotsResponse struct {
	RequestId string
	Snapshots DescribeSnapshotsSnapshotList
}

type DescribeSnapshotsSnapshotList []DescribeSnapshotsSnapshot

func (list *DescribeSnapshotsSnapshotList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSnapshotsSnapshot)
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

func (c *RKvstoreClient) DescribeRegions(req *DescribeRegionsArgs) (resp *DescribeRegionsResponse, err error) {
	resp = &DescribeRegionsResponse{}
	err = c.Request("DescribeRegions", req, resp)
	return
}

type DescribeRegionsKVStoreRegion struct {
	RegionId  string
	ZoneIds   string
	LocalName string
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
	RegionIds DescribeRegionsKVStoreRegionList
}

type DescribeRegionsKVStoreRegionList []DescribeRegionsKVStoreRegion

func (list *DescribeRegionsKVStoreRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsKVStoreRegion)
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

func (c *RKvstoreClient) VerifyPassword(req *VerifyPasswordArgs) (resp *VerifyPasswordResponse, err error) {
	resp = &VerifyPasswordResponse{}
	err = c.Request("VerifyPassword", req, resp)
	return
}

type VerifyPasswordArgs struct {
	ResourceOwnerId      int64
	Password             string
	InstanceId           string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type VerifyPasswordResponse struct {
	RequestId string
}

func (c *RKvstoreClient) DescribeMonthlyServiceStatusDetail(req *DescribeMonthlyServiceStatusDetailArgs) (resp *DescribeMonthlyServiceStatusDetailResponse, err error) {
	resp = &DescribeMonthlyServiceStatusDetailResponse{}
	err = c.Request("DescribeMonthlyServiceStatusDetail", req, resp)
	return
}

type DescribeMonthlyServiceStatusDetailAffectedInfo struct {
	StartTime   string
	EndTime     string
	Description string
}
type DescribeMonthlyServiceStatusDetailArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	Month                string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeMonthlyServiceStatusDetailResponse struct {
	RequestId     string
	InstanceId    string
	UptimePct     float32
	AffectedInfos DescribeMonthlyServiceStatusDetailAffectedInfoList
}

type DescribeMonthlyServiceStatusDetailAffectedInfoList []DescribeMonthlyServiceStatusDetailAffectedInfo

func (list *DescribeMonthlyServiceStatusDetailAffectedInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMonthlyServiceStatusDetailAffectedInfo)
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
