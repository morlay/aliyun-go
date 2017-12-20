package polardb

import (
	"encoding/json"

	"github.com/morlay/aliyun-go/core"
)

func (c *PolardbClient) DescribeBackupPolicy(req *DescribeBackupPolicyArgs) (resp *DescribeBackupPolicyResponse, err error) {
	resp = &DescribeBackupPolicyResponse{}
	err = c.Request("DescribeBackupPolicy", req, resp)
	return
}

type DescribeBackupPolicyArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	DBClusterId          string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeBackupPolicyResponse struct {
	RequestId               string
	BackupRetentionPeriod   int
	PreferredNextBackupTime string
	PreferredBackupTime     string
	PreferredBackupPeriod   string
}

func (c *PolardbClient) ModifyAccountDescription(req *ModifyAccountDescriptionArgs) (resp *ModifyAccountDescriptionResponse, err error) {
	resp = &ModifyAccountDescriptionResponse{}
	err = c.Request("ModifyAccountDescription", req, resp)
	return
}

type ModifyAccountDescriptionArgs struct {
	ResourceOwnerId      int64
	AccountName          string
	ResourceOwnerAccount string
	DBClusterId          string
	OwnerAccount         string
	OwnerId              int64
	AccountDescription   string
}
type ModifyAccountDescriptionResponse struct {
	RequestId string
}

func (c *PolardbClient) DescribeTaskInfo(req *DescribeTaskInfoArgs) (resp *DescribeTaskInfoResponse, err error) {
	resp = &DescribeTaskInfoResponse{}
	err = c.Request("DescribeTaskInfo", req, resp)
	return
}

type DescribeTaskInfoArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	TaskId               string
}
type DescribeTaskInfoResponse struct {
	RequestId          string
	TaskId             string
	BeginTime          string
	FinishTime         string
	ExpectedFinishTime string
	TaskAction         string
	Progress           int
	DBName             string
	ProgressInfo       string
}

func (c *PolardbClient) ModifyDBClusterConnectionString(req *ModifyDBClusterConnectionStringArgs) (resp *ModifyDBClusterConnectionStringResponse, err error) {
	resp = &ModifyDBClusterConnectionStringResponse{}
	err = c.Request("ModifyDBClusterConnectionString", req, resp)
	return
}

type ModifyDBClusterConnectionStringArgs struct {
	ResourceOwnerId         int64
	ConnectionStringPrefix  string
	ResourceOwnerAccount    string
	DBClusterId             string
	OwnerAccount            string
	OwnerId                 int64
	CurrentConnectionString string
}
type ModifyDBClusterConnectionStringResponse struct {
	RequestId           string
	DBClusterId         string
	OldConnectionString string
	OldPort             string
	NewConnectionString string
	NewPort             string
	IPType              string
}

func (c *PolardbClient) ModifyDBInstanceMaintainTime(req *ModifyDBInstanceMaintainTimeArgs) (resp *ModifyDBInstanceMaintainTimeResponse, err error) {
	resp = &ModifyDBInstanceMaintainTimeResponse{}
	err = c.Request("ModifyDBInstanceMaintainTime", req, resp)
	return
}

type ModifyDBInstanceMaintainTimeArgs struct {
	MaintainTime         string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type ModifyDBInstanceMaintainTimeResponse struct {
	RequestId string
}

func (c *PolardbClient) DescribeDBClusterIPArrayList(req *DescribeDBClusterIPArrayListArgs) (resp *DescribeDBClusterIPArrayListResponse, err error) {
	resp = &DescribeDBClusterIPArrayListResponse{}
	err = c.Request("DescribeDBClusterIPArrayList", req, resp)
	return
}

type DescribeDBClusterIPArrayListDBClusterIPArray struct {
	DBClusterIPArrayName      string
	DBClusterIPArrayAttribute string
	SecurityIPList            string
}
type DescribeDBClusterIPArrayListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	DBClusterId          string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeDBClusterIPArrayListResponse struct {
	RequestId string
	Items     DescribeDBClusterIPArrayListDBClusterIPArrayList
}

type DescribeDBClusterIPArrayListDBClusterIPArrayList []DescribeDBClusterIPArrayListDBClusterIPArray

func (list *DescribeDBClusterIPArrayListDBClusterIPArrayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBClusterIPArrayListDBClusterIPArray)
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

func (c *PolardbClient) CreateDBCluster(req *CreateDBClusterArgs) (resp *CreateDBClusterResponse, err error) {
	resp = &CreateDBClusterResponse{}
	err = c.Request("CreateDBCluster", req, resp)
	return
}

type CreateDBClusterArgs struct {
	ResourceOwnerId      int64
	DBClusterDescription string
	Period               string
	DBInstanceStorage    string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	OwnerId              int64
	UsedTime             string
	DBInstanceClass      string
	ClusterNetworkType   string
	SecurityIPList       string
	VSwitchId            string
	PrivateIpAddress     string
	Engine               string
	VPCId                string
	DBType               string
	ZoneId               string
	DBVersion            string
	PayType              string
}
type CreateDBClusterResponse struct {
	RequestId        string
	DBClusterId      string
	OrderId          string
	ConnectionString string
	Port             string
}

func (c *PolardbClient) CreateDBInstance(req *CreateDBInstanceArgs) (resp *CreateDBInstanceResponse, err error) {
	resp = &CreateDBInstanceResponse{}
	err = c.Request("CreateDBInstance", req, resp)
	return
}

type CreateDBInstanceArgs struct {
	ResourceOwnerId       int64
	ResourceOwnerAccount  string
	ClientToken           string
	OwnerAccount          string
	OwnerId               int64
	DBInstanceClass       string
	SecurityIPList        string
	VSwitchId             string
	PrivateIpAddress      string
	DBInstanceDescription string
}
type CreateDBInstanceResponse struct {
	RequestId    string
	DBClusterId  string
	DBInstanceId string
	OrderId      string
}

func (c *PolardbClient) ResetAccountPassword(req *ResetAccountPasswordArgs) (resp *ResetAccountPasswordResponse, err error) {
	resp = &ResetAccountPasswordResponse{}
	err = c.Request("ResetAccountPassword", req, resp)
	return
}

type ResetAccountPasswordArgs struct {
	ResourceOwnerId      int64
	AccountPassword      string
	AccountName          string
	ResourceOwnerAccount string
	DBClusterId          string
	OwnerAccount         string
	OwnerId              int64
}
type ResetAccountPasswordResponse struct {
	RequestId string
}

func (c *PolardbClient) ModifyParameter(req *ModifyParameterArgs) (resp *ModifyParameterResponse, err error) {
	resp = &ModifyParameterResponse{}
	err = c.Request("ModifyParameter", req, resp)
	return
}

type ModifyParameterArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	Forcerestart         core.Bool
	OwnerId              int64
	Parameters           string
}
type ModifyParameterResponse struct {
	RequestId string
}

func (c *PolardbClient) CreateBackup(req *CreateBackupArgs) (resp *CreateBackupResponse, err error) {
	resp = &CreateBackupResponse{}
	err = c.Request("CreateBackup", req, resp)
	return
}

type CreateBackupArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	DBClusterId          string
	OwnerAccount         string
	OwnerId              int64
}
type CreateBackupResponse struct {
	RequestId string
}

func (c *PolardbClient) AllocateInstancePublicConnection(req *AllocateInstancePublicConnectionArgs) (resp *AllocateInstancePublicConnectionResponse, err error) {
	resp = &AllocateInstancePublicConnectionResponse{}
	err = c.Request("AllocateInstancePublicConnection", req, resp)
	return
}

type AllocateInstancePublicConnectionArgs struct {
	ResourceOwnerId        int64
	ConnectionStringPrefix string
	ResourceOwnerAccount   string
	Port                   string
	OwnerAccount           string
	DBInstanceId           string
	OwnerId                int64
}
type AllocateInstancePublicConnectionResponse struct {
	RequestId        string
	DBInstanceId     string
	ConnectionString string
	IPType           string
	Port             string
}

func (c *PolardbClient) DescribeBackups(req *DescribeBackupsArgs) (resp *DescribeBackupsResponse, err error) {
	resp = &DescribeBackupsResponse{}
	err = c.Request("DescribeBackups", req, resp)
	return
}

type DescribeBackupsBackup struct {
	BackupId                  string
	DBInstanceId              string
	BackupStatus              string
	BackupStartTime           string
	BackupEndTime             string
	BackupType                string
	BackupMode                string
	BackupMethod              string
	BackupDownloadURL         string
	BackupIntranetDownloadURL string
	BackupLocation            string
	BackupExtractionStatus    string
	BackupScale               string
	BackupDBNames             string
	TotalBackupSize           int64
	BackupSize                int64
	HostInstanceID            string
	StoreStatus               string
}
type DescribeBackupsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	DBClusterId          string
	OwnerAccount         string
	BackupId             string
	EndTime              string
	StartTime            string
	OwnerId              int64
	PageNumber           int
	BackupStatus         string
	PageSize             int
	BackupMode           string
}
type DescribeBackupsResponse struct {
	RequestId        string
	TotalRecordCount string
	PageNumber       string
	PageRecordCount  string
	TotalBackupSize  int64
	Items            DescribeBackupsBackupList
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

func (c *PolardbClient) CreateAccount(req *CreateAccountArgs) (resp *CreateAccountResponse, err error) {
	resp = &CreateAccountResponse{}
	err = c.Request("CreateAccount", req, resp)
	return
}

type CreateAccountArgs struct {
	ResourceOwnerId      int64
	AccountPassword      string
	AccountName          string
	ResourceOwnerAccount string
	DBClusterId          string
	OwnerAccount         string
	DatabaseName         string
	OwnerId              int64
	AccountDescription   string
}
type CreateAccountResponse struct {
	RequestId string
}

func (c *PolardbClient) RestartDBInstance(req *RestartDBInstanceArgs) (resp *RestartDBInstanceResponse, err error) {
	resp = &RestartDBInstanceResponse{}
	err = c.Request("RestartDBInstance", req, resp)
	return
}

type RestartDBInstanceArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type RestartDBInstanceResponse struct {
	RequestId string
}

func (c *PolardbClient) AllocateClusterPublicConnection(req *AllocateClusterPublicConnectionArgs) (resp *AllocateClusterPublicConnectionResponse, err error) {
	resp = &AllocateClusterPublicConnectionResponse{}
	err = c.Request("AllocateClusterPublicConnection", req, resp)
	return
}

type AllocateClusterPublicConnectionArgs struct {
	ResourceOwnerId        int64
	ConnectionStringPrefix string
	ResourceOwnerAccount   string
	Port                   string
	DBClusterId            string
	OwnerAccount           string
	OwnerId                int64
}
type AllocateClusterPublicConnectionResponse struct {
	RequestId        string
	DBInstanceId     string
	ConnectionString string
	IPType           string
	Port             string
}

func (c *PolardbClient) DescribeDBInstances(req *DescribeDBInstancesArgs) (resp *DescribeDBInstancesResponse, err error) {
	resp = &DescribeDBInstancesResponse{}
	err = c.Request("DescribeDBInstances", req, resp)
	return
}

type DescribeDBInstancesDBInstance struct {
	DBInstanceId          string
	DBInstanceDescription string
	PayType               string
	DBInstanceType        string
	DBInstanceClass       string
	InstanceNetworkType   string
	RegionId              string
	ZoneId                string
	DBClusterId           string
	ExpireTime            string
	DBInstanceStatus      string
	Engine                string
	DBType                string
	DBVersion             string
	DBInstanceType1       string
	LockMode              string
	LockReason            string
	CreateTime            string
	VpcId                 string
	VSwitchId             string
}
type DescribeDBInstancesArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	OwnerId              int64
	InstanceNetworkType  string
	PageNumber           int
}
type DescribeDBInstancesResponse struct {
	RequestId        string
	PageNumber       int
	TotalRecordCount int
	PageRecordCount  int
	Items            DescribeDBInstancesDBInstanceList
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

func (c *PolardbClient) DescribeDBClusters(req *DescribeDBClustersArgs) (resp *DescribeDBClustersResponse, err error) {
	resp = &DescribeDBClustersResponse{}
	err = c.Request("DescribeDBClusters", req, resp)
	return
}

type DescribeDBClustersDBCluster struct {
	DBClusterId          string
	DBClusterDescription string
	PayType              string
	DBClusterNetworkType string
	RegionId             string
	ExpireTime           string
	DBClusterStatus      string
	Engine               string
	DBType               string
	DBVersion            string
	LockMode             string
	LockReason           string
	CreateTime           string
	VpcId                string
	VSwitchId            string
}
type DescribeDBClustersArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBType               string
	PageSize             int
	OwnerId              int64
	PageNumber           int
}
type DescribeDBClustersResponse struct {
	RequestId        string
	PageNumber       int
	TotalRecordCount int
	PageRecordCount  int
	Items            DescribeDBClustersDBClusterList
}

type DescribeDBClustersDBClusterList []DescribeDBClustersDBCluster

func (list *DescribeDBClustersDBClusterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBClustersDBCluster)
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

func (c *PolardbClient) ModifySecurityIps(req *ModifySecurityIpsArgs) (resp *ModifySecurityIpsResponse, err error) {
	resp = &ModifySecurityIpsResponse{}
	err = c.Request("ModifySecurityIps", req, resp)
	return
}

type ModifySecurityIpsArgs struct {
	ResourceOwnerId           int64
	ResourceOwnerAccount      string
	DBClusterId               string
	OwnerAccount              string
	SecurityIps               string
	DBClusterIPArrayName      string
	OwnerId                   int64
	DBClusterIPArrayAttribute string
}
type ModifySecurityIpsResponse struct {
	RequestId string
}

func (c *PolardbClient) DescribeParameterTemplates(req *DescribeParameterTemplatesArgs) (resp *DescribeParameterTemplatesResponse, err error) {
	resp = &DescribeParameterTemplatesResponse{}
	err = c.Request("DescribeParameterTemplates", req, resp)
	return
}

type DescribeParameterTemplatesTemplateRecord struct {
	ParameterName        string
	ParameterValue       string
	ForceModify          string
	ForceRestart         string
	CheckingCode         string
	ParameterDescription string
}
type DescribeParameterTemplatesArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeParameterTemplatesResponse struct {
	RequestId      string
	Engine         string
	DBType         string
	DBVersion      string
	ParameterCount string
	Parameters     DescribeParameterTemplatesTemplateRecordList
}

type DescribeParameterTemplatesTemplateRecordList []DescribeParameterTemplatesTemplateRecord

func (list *DescribeParameterTemplatesTemplateRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeParameterTemplatesTemplateRecord)
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

func (c *PolardbClient) ModifyDBInstanceConnectionString(req *ModifyDBInstanceConnectionStringArgs) (resp *ModifyDBInstanceConnectionStringResponse, err error) {
	resp = &ModifyDBInstanceConnectionStringResponse{}
	err = c.Request("ModifyDBInstanceConnectionString", req, resp)
	return
}

type ModifyDBInstanceConnectionStringArgs struct {
	ResourceOwnerId         int64
	ConnectionStringPrefix  string
	ResourceOwnerAccount    string
	OwnerAccount            string
	DBInstanceId            string
	OwnerId                 int64
	CurrentConnectionString string
}
type ModifyDBInstanceConnectionStringResponse struct {
	RequestId           string
	DBInstanceId        string
	OldConnectionString string
	OldPort             string
	NewConnectionString string
	NewPort             string
	IPType              string
}

func (c *PolardbClient) ModifyDBInstanceDescription(req *ModifyDBInstanceDescriptionArgs) (resp *ModifyDBInstanceDescriptionResponse, err error) {
	resp = &ModifyDBInstanceDescriptionResponse{}
	err = c.Request("ModifyDBInstanceDescription", req, resp)
	return
}

type ModifyDBInstanceDescriptionArgs struct {
	ResourceOwnerId       int64
	ResourceOwnerAccount  string
	OwnerAccount          string
	DBInstanceId          string
	DBInstanceDescription string
	OwnerId               int64
}
type ModifyDBInstanceDescriptionResponse struct {
	RequestId string
}

func (c *PolardbClient) DescribeResourceUsage(req *DescribeResourceUsageArgs) (resp *DescribeResourceUsageResponse, err error) {
	resp = &DescribeResourceUsageResponse{}
	err = c.Request("DescribeResourceUsage", req, resp)
	return
}

type DescribeResourceUsageArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeResourceUsageResponse struct {
	RequestId    string
	DBInstanceId string
	Engine       string
	DBType       string
	DBVersion    string
	DiskUsed     int64
	DataSize     int64
	LogSize      int64
	BackupSize   int64
}

func (c *PolardbClient) DeleteDBCluster(req *DeleteDBClusterArgs) (resp *DeleteDBClusterResponse, err error) {
	resp = &DeleteDBClusterResponse{}
	err = c.Request("DeleteDBCluster", req, resp)
	return
}

type DeleteDBClusterArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	DBClusterId          string
	OwnerAccount         string
	OwnerId              int64
}
type DeleteDBClusterResponse struct {
	RequestId string
}

func (c *PolardbClient) ReleaseInstancePublicConnection(req *ReleaseInstancePublicConnectionArgs) (resp *ReleaseInstancePublicConnectionResponse, err error) {
	resp = &ReleaseInstancePublicConnectionResponse{}
	err = c.Request("ReleaseInstancePublicConnection", req, resp)
	return
}

type ReleaseInstancePublicConnectionArgs struct {
	ResourceOwnerId         int64
	ResourceOwnerAccount    string
	OwnerAccount            string
	DBInstanceId            string
	OwnerId                 int64
	CurrentConnectionString string
}
type ReleaseInstancePublicConnectionResponse struct {
	RequestId string
}

func (c *PolardbClient) DescribeTasks(req *DescribeTasksArgs) (resp *DescribeTasksResponse, err error) {
	resp = &DescribeTasksResponse{}
	err = c.Request("DescribeTasks", req, resp)
	return
}

type DescribeTasksTask struct {
	TaskId             string
	BeginTime          string
	FinishTime         string
	ExpectedFinishTime string
	TaskAction         string
	Progress           int
	DBName             string
	ProgressInfo       string
	TaskErrorCode      string
	TaskErrorMessage   string
}
type DescribeTasksArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	EndTime              string
	DBInstanceId         string
	StartTime            string
	OwnerId              int64
	PageNumber           int
	Status               string
}
type DescribeTasksResponse struct {
	RequestId        string
	StartTime        string
	EndTime          string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Tasks            DescribeTasksTaskList
}

type DescribeTasksTaskList []DescribeTasksTask

func (list *DescribeTasksTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTasksTask)
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

func (c *PolardbClient) ReleaseClusterPublicConnection(req *ReleaseClusterPublicConnectionArgs) (resp *ReleaseClusterPublicConnectionResponse, err error) {
	resp = &ReleaseClusterPublicConnectionResponse{}
	err = c.Request("ReleaseClusterPublicConnection", req, resp)
	return
}

type ReleaseClusterPublicConnectionArgs struct {
	ResourceOwnerId         int64
	ResourceOwnerAccount    string
	DBClusterId             string
	OwnerAccount            string
	OwnerId                 int64
	CurrentConnectionString string
}
type ReleaseClusterPublicConnectionResponse struct {
	RequestId string
}

func (c *PolardbClient) DescribeDBInstanceNetInfo(req *DescribeDBInstanceNetInfoArgs) (resp *DescribeDBInstanceNetInfoResponse, err error) {
	resp = &DescribeDBInstanceNetInfoResponse{}
	err = c.Request("DescribeDBInstanceNetInfo", req, resp)
	return
}

type DescribeDBInstanceNetInfoDBInstanceNetInfo struct {
	ConnectionString string
	IPAddress        string
	IPType           string
	Port             string
	VPCId            string
	VSwitchId        string
}
type DescribeDBInstanceNetInfoArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeDBInstanceNetInfoResponse struct {
	RequestId           string
	InstanceNetworkType string
	DBInstanceNetInfos  DescribeDBInstanceNetInfoDBInstanceNetInfoList
}

type DescribeDBInstanceNetInfoDBInstanceNetInfoList []DescribeDBInstanceNetInfoDBInstanceNetInfo

func (list *DescribeDBInstanceNetInfoDBInstanceNetInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetInfoDBInstanceNetInfo)
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

func (c *PolardbClient) FailoverDBCluster(req *FailoverDBClusterArgs) (resp *FailoverDBClusterResponse, err error) {
	resp = &FailoverDBClusterResponse{}
	err = c.Request("FailoverDBCluster", req, resp)
	return
}

type FailoverDBClusterArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	DBClusterId          string
	OwnerAccount         string
	TargetDBInstanceId   string
	OwnerId              int64
}
type FailoverDBClusterResponse struct {
	RequestId string
}

func (c *PolardbClient) ModifyBackupPolicy(req *ModifyBackupPolicyArgs) (resp *ModifyBackupPolicyResponse, err error) {
	resp = &ModifyBackupPolicyResponse{}
	err = c.Request("ModifyBackupPolicy", req, resp)
	return
}

type ModifyBackupPolicyArgs struct {
	PreferredBackupTime   string
	PreferredBackupPeriod string
	BackupRetentionPeriod string
	ResourceOwnerId       int64
	ResourceOwnerAccount  string
	DBClusterId           string
	OwnerAccount          string
	OwnerId               int64
}
type ModifyBackupPolicyResponse struct {
	RequestId string
}

func (c *PolardbClient) DescribeAccounts(req *DescribeAccountsArgs) (resp *DescribeAccountsResponse, err error) {
	resp = &DescribeAccountsResponse{}
	err = c.Request("DescribeAccounts", req, resp)
	return
}

type DescribeAccountsDBInstanceAccount struct {
	DBClusterId        string
	AccountName        string
	AccountStatus      string
	AccountDescription string
	AccountType        string
}
type DescribeAccountsArgs struct {
	ResourceOwnerId      int64
	AccountName          string
	ResourceOwnerAccount string
	DBClusterId          string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeAccountsResponse struct {
	RequestId string
	Accounts  DescribeAccountsDBInstanceAccountList
}

type DescribeAccountsDBInstanceAccountList []DescribeAccountsDBInstanceAccount

func (list *DescribeAccountsDBInstanceAccountList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccountsDBInstanceAccount)
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

func (c *PolardbClient) DescribeRegions(req *DescribeRegionsArgs) (resp *DescribeRegionsResponse, err error) {
	resp = &DescribeRegionsResponse{}
	err = c.Request("DescribeRegions", req, resp)
	return
}

type DescribeRegionsRegion struct {
	RegionId string
	Zones    DescribeRegionsZoneList
}

type DescribeRegionsZone struct {
	ZoneId     string
	VpcEnabled core.Bool
}
type DescribeRegionsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeRegionsResponse struct {
	RequestId string
	Regions   DescribeRegionsRegionList
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

func (c *PolardbClient) DescribeClusterNetInfo(req *DescribeClusterNetInfoArgs) (resp *DescribeClusterNetInfoResponse, err error) {
	resp = &DescribeClusterNetInfoResponse{}
	err = c.Request("DescribeClusterNetInfo", req, resp)
	return
}

type DescribeClusterNetInfoDBInstanceNetInfo struct {
	ConnectionString string
	IPAddress        string
	IPType           string
	Port             string
	VPCId            string
	VSwitchId        string
}
type DescribeClusterNetInfoArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeClusterNetInfoResponse struct {
	RequestId            string
	DBClusterNetworkType string
	DBInstanceNetInfos   DescribeClusterNetInfoDBInstanceNetInfoList
}

type DescribeClusterNetInfoDBInstanceNetInfoList []DescribeClusterNetInfoDBInstanceNetInfo

func (list *DescribeClusterNetInfoDBInstanceNetInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterNetInfoDBInstanceNetInfo)
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

func (c *PolardbClient) DescribeParameters(req *DescribeParametersArgs) (resp *DescribeParametersResponse, err error) {
	resp = &DescribeParametersResponse{}
	err = c.Request("DescribeParameters", req, resp)
	return
}

type DescribeParametersDBInstanceParameter struct {
	ParameterName        string
	ParameterValue       string
	ParameterDescription string
}
type DescribeParametersArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeParametersResponse struct {
	RequestId         string
	Engine            string
	DBType            string
	DBVersion         string
	ConfigParameters  DescribeParametersDBInstanceParameterList
	RunningParameters DescribeParametersDBInstanceParameterList
}

type DescribeParametersDBInstanceParameterList []DescribeParametersDBInstanceParameter

func (list *DescribeParametersDBInstanceParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeParametersDBInstanceParameter)
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

func (c *PolardbClient) DescribeDBInstancePerformance(req *DescribeDBInstancePerformanceArgs) (resp *DescribeDBInstancePerformanceResponse, err error) {
	resp = &DescribeDBInstancePerformanceResponse{}
	err = c.Request("DescribeDBInstancePerformance", req, resp)
	return
}

type DescribeDBInstancePerformancePerformanceItem struct {
	MetricName  string
	Measurement string
	Points      DescribeDBInstancePerformancePerformanceItemValueList
}

type DescribeDBInstancePerformancePerformanceItemValue struct {
	Value     float32
	Timestamp int64
}
type DescribeDBInstancePerformanceArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	EndTime              string
	DBInstanceId         string
	StartTime            string
	OwnerId              int64
	Key                  string
}
type DescribeDBInstancePerformanceResponse struct {
	RequestId       string
	DBInstanceId    string
	Engine          string
	DBType          string
	DBVersion       string
	StartTime       string
	EndTime         string
	PerformanceKeys DescribeDBInstancePerformancePerformanceItemList
}

type DescribeDBInstancePerformancePerformanceItemValueList []DescribeDBInstancePerformancePerformanceItemValue

func (list *DescribeDBInstancePerformancePerformanceItemValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancePerformancePerformanceItemValue)
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

type DescribeDBInstancePerformancePerformanceItemList []DescribeDBInstancePerformancePerformanceItem

func (list *DescribeDBInstancePerformancePerformanceItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancePerformancePerformanceItem)
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

func (c *PolardbClient) ModifyDBClusterDescription(req *ModifyDBClusterDescriptionArgs) (resp *ModifyDBClusterDescriptionResponse, err error) {
	resp = &ModifyDBClusterDescriptionResponse{}
	err = c.Request("ModifyDBClusterDescription", req, resp)
	return
}

type ModifyDBClusterDescriptionArgs struct {
	ResourceOwnerId      int64
	DBClusterDescription string
	ResourceOwnerAccount string
	DBClusterId          string
	OwnerAccount         string
	OwnerId              int64
}
type ModifyDBClusterDescriptionResponse struct {
	RequestId string
}

func (c *PolardbClient) DescribeDBClusterNetInfo(req *DescribeDBClusterNetInfoArgs) (resp *DescribeDBClusterNetInfoResponse, err error) {
	resp = &DescribeDBClusterNetInfoResponse{}
	err = c.Request("DescribeDBClusterNetInfo", req, resp)
	return
}

type DescribeDBClusterNetInfoDBClusterNetInfo struct {
	ConnectionString string
	IPAddress        string
	IPType           string
	Port             string
	VPCId            string
	VSwitchId        string
}
type DescribeDBClusterNetInfoArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	DBClusterId          string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeDBClusterNetInfoResponse struct {
	RequestId          string
	ClusterNetworkType string
	DBClusterNetInfos  DescribeDBClusterNetInfoDBClusterNetInfoList
}

type DescribeDBClusterNetInfoDBClusterNetInfoList []DescribeDBClusterNetInfoDBClusterNetInfo

func (list *DescribeDBClusterNetInfoDBClusterNetInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBClusterNetInfoDBClusterNetInfo)
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

func (c *PolardbClient) DescribeDBInstanceAttribute(req *DescribeDBInstanceAttributeArgs) (resp *DescribeDBInstanceAttributeResponse, err error) {
	resp = &DescribeDBInstanceAttributeResponse{}
	err = c.Request("DescribeDBInstanceAttribute", req, resp)
	return
}

type DescribeDBInstanceAttributeDBInstanceAttribute struct {
	DBInstanceId          string
	DBClusterDescription  string
	DBClusterId           string
	PayType               string
	DBInstanceType        string
	RegionId              string
	ZoneId                string
	Engine                string
	DBType                string
	DBVersion             string
	DBInstanceClass       string
	DBInstanceStorage     int64
	DBInstanceStatus      string
	DBInstanceDescription string
	ConnectionString      int64
	Port                  int64
	DBInstanceNetType     string
	LockMode              string
	LockReason            string
	CreationTime          string
	ExpireTime            string
	MaintainStartTime     string
	MaintainEndTime       string
	MaxConnections        int
	MaxIOPS               int
	SecurityIPList        string
	InstanceNetworkType   string
	VpcId                 string
	VSwitchId             string
	DBInstanceType1       string
}
type DescribeDBInstanceAttributeArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeDBInstanceAttributeResponse struct {
	RequestId string
	Items     DescribeDBInstanceAttributeDBInstanceAttributeList
}

type DescribeDBInstanceAttributeDBInstanceAttributeList []DescribeDBInstanceAttributeDBInstanceAttribute

func (list *DescribeDBInstanceAttributeDBInstanceAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceAttributeDBInstanceAttribute)
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

func (c *PolardbClient) DescribeDBClusterAttribute(req *DescribeDBClusterAttributeArgs) (resp *DescribeDBClusterAttributeResponse, err error) {
	resp = &DescribeDBClusterAttributeResponse{}
	err = c.Request("DescribeDBClusterAttribute", req, resp)
	return
}

type DescribeDBClusterAttributeDBClusterAttribute struct {
	RegionId             string
	DBClusterNetworkType string
	VPCId                string
	VSwitchId            string
	PayType              string
	DBClusterId          string
	DBClusterStatus      string
	DBClusterDescription string
	Engine               string
	DBType               string
	DBVersion            string
	Storage              int64
	ConnectionString     int64
	Port                 int64
	DBClusterNetType     string
	LockMode             string
	LockReason           string
	CreationTime         string
	ExpireTime           string
	DbInstances          DescribeDBClusterAttributeDbInstanceList
}

type DescribeDBClusterAttributeDbInstance struct {
	DBInstanceId          string
	DBInstanceStatus      string
	DBInstanceDescription string
	Engine                string
	DBType                string
	DBVersion             string
	DBInstanceStorage     string
	LockMode              string
	LockReason            string
	MaintainStartTime     string
	MaintainEndTime       string
	CreationTime          string
	DBInstanceClass       string
	SecurityIPList        string
	DBInstanceType        string
}
type DescribeDBClusterAttributeArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	DBClusterId          string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeDBClusterAttributeResponse struct {
	RequestId string
	Items     DescribeDBClusterAttributeDBClusterAttributeList
}

type DescribeDBClusterAttributeDbInstanceList []DescribeDBClusterAttributeDbInstance

func (list *DescribeDBClusterAttributeDbInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBClusterAttributeDbInstance)
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

type DescribeDBClusterAttributeDBClusterAttributeList []DescribeDBClusterAttributeDBClusterAttribute

func (list *DescribeDBClusterAttributeDBClusterAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBClusterAttributeDBClusterAttribute)
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
