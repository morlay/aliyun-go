package rds

import (
	"encoding/json"

	"github.com/morlay/aliyun-go/core"
)

func (c *RdsClient) ModifyDBInstanceHAConfig(req *ModifyDBInstanceHAConfigArgs) (resp *ModifyDBInstanceHAConfigResponse, err error) {
	resp = &ModifyDBInstanceHAConfigResponse{}
	err = c.Request("ModifyDBInstanceHAConfig", req, resp)
	return
}

type ModifyDBInstanceHAConfigArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	SyncMode             string
	DbInstanceId         string
	OwnerId              int64
	HAMode               string
}
type ModifyDBInstanceHAConfigResponse struct {
	RequestId string
}

func (c *RdsClient) CompensateInstanceForChannel(req *CompensateInstanceForChannelArgs) (resp *CompensateInstanceForChannelResponse, err error) {
	resp = &CompensateInstanceForChannelResponse{}
	err = c.Request("CompensateInstanceForChannel", req, resp)
	return
}

type CompensateInstanceForChannelArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ZoneId               string
	SubDomain            string
	DBInstanceId         string
	OwnerId              int64
}
type CompensateInstanceForChannelResponse struct {
	RequestId string
}

func (c *RdsClient) ModifyDampPolicy(req *ModifyDampPolicyArgs) (resp *ModifyDampPolicyResponse, err error) {
	resp = &ModifyDampPolicyResponse{}
	err = c.Request("ModifyDampPolicy", req, resp)
	return
}

type ModifyDampPolicyArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	TimeRules            string
	ActionRules          string
	SecurityToken        string
	Handlers             string
	DBInstanceId         string
	PolicyName           string
	SourceRules          string
}
type ModifyDampPolicyResponse struct {
	RequestId  string
	PolicyId   string
	PolicyName string
}

func (c *RdsClient) AllocateReadWriteSplittingConnection(req *AllocateReadWriteSplittingConnectionArgs) (resp *AllocateReadWriteSplittingConnectionResponse, err error) {
	resp = &AllocateReadWriteSplittingConnectionResponse{}
	err = c.Request("AllocateReadWriteSplittingConnection", req, resp)
	return
}

type AllocateReadWriteSplittingConnectionArgs struct {
	ResourceOwnerId        int64
	ConnectionStringPrefix string
	ResourceOwnerAccount   string
	OwnerAccount           string
	Weight                 string
	OwnerId                int64
	IPType                 string
	Port                   string
	DistributionType       string
	DBInstanceId           string
	MaxDelayTime           string
}
type AllocateReadWriteSplittingConnectionResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeBackupPolicy(req *DescribeBackupPolicyArgs) (resp *DescribeBackupPolicyResponse, err error) {
	resp = &DescribeBackupPolicyResponse{}
	err = c.Request("DescribeBackupPolicy", req, resp)
	return
}

type DescribeBackupPolicyArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeBackupPolicyResponse struct {
	RequestId                string
	BackupRetentionPeriod    int
	PreferredNextBackupTime  string
	PreferredBackupTime      string
	PreferredBackupPeriod    string
	BackupLog                string
	LogBackupRetentionPeriod int
}

func (c *RdsClient) DescribeReplicaUsage(req *DescribeReplicaUsageArgs) (resp *DescribeReplicaUsageResponse, err error) {
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

func (c *RdsClient) DescribeSQLDiagnosisList(req *DescribeSQLDiagnosisListArgs) (resp *DescribeSQLDiagnosisListResponse, err error) {
	resp = &DescribeSQLDiagnosisListResponse{}
	err = c.Request("DescribeSQLDiagnosisList", req, resp)
	return
}

type DescribeSQLDiagnosisListSQLDiag struct {
	SQLDiagId string
	StartTime string
	EndTime   string
	Status    int
	Progress  int
}
type DescribeSQLDiagnosisListArgs struct {
	DBInstanceId string
}
type DescribeSQLDiagnosisListResponse struct {
	RequestId   string
	SQLDiagList DescribeSQLDiagnosisListSQLDiagList
}

type DescribeSQLDiagnosisListSQLDiagList []DescribeSQLDiagnosisListSQLDiag

func (list *DescribeSQLDiagnosisListSQLDiagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLDiagnosisListSQLDiag)
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

func (c *RdsClient) ModifyReplicaDescription(req *ModifyReplicaDescriptionArgs) (resp *ModifyReplicaDescriptionResponse, err error) {
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

func (c *RdsClient) ModifyAccountDescription(req *ModifyAccountDescriptionArgs) (resp *ModifyAccountDescriptionResponse, err error) {
	resp = &ModifyAccountDescriptionResponse{}
	err = c.Request("ModifyAccountDescription", req, resp)
	return
}

type ModifyAccountDescriptionArgs struct {
	ResourceOwnerId      int64
	AccountName          string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
	AccountDescription   string
}
type ModifyAccountDescriptionResponse struct {
	RequestId string
}

func (c *RdsClient) CheckAccountNameAvailable(req *CheckAccountNameAvailableArgs) (resp *CheckAccountNameAvailableResponse, err error) {
	resp = &CheckAccountNameAvailableResponse{}
	err = c.Request("CheckAccountNameAvailable", req, resp)
	return
}

type CheckAccountNameAvailableArgs struct {
	ResourceOwnerId      int64
	AccountName          string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type CheckAccountNameAvailableResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeTags(req *DescribeTagsArgs) (resp *DescribeTagsResponse, err error) {
	resp = &DescribeTagsResponse{}
	err = c.Request("DescribeTags", req, resp)
	return
}

type DescribeTagsTagInfos struct {
	TagKey        string
	TagValue      string
	DBInstanceIds DescribeTagsDBInstanceIdList
}
type DescribeTagsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
	ProxyId              string
	Tags                 string
}
type DescribeTagsResponse struct {
	RequestId string
	Items     DescribeTagsTagInfosList
}

type DescribeTagsDBInstanceIdList []string

func (list *DescribeTagsDBInstanceIdList) UnmarshalJSON(data []byte) error {
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

type DescribeTagsTagInfosList []DescribeTagsTagInfos

func (list *DescribeTagsTagInfosList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTagsTagInfos)
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

func (c *RdsClient) CreateDiagnosticReport(req *CreateDiagnosticReportArgs) (resp *CreateDiagnosticReportResponse, err error) {
	resp = &CreateDiagnosticReportResponse{}
	err = c.Request("CreateDiagnosticReport", req, resp)
	return
}

type CreateDiagnosticReportArgs struct {
	EndTime      string
	DBInstanceId string
	StartTime    string
}
type CreateDiagnosticReportResponse struct {
	RequestId string
	ReportId  string
}

func (c *RdsClient) CloneDBInstanceForSecurity(req *CloneDBInstanceForSecurityArgs) (resp *CloneDBInstanceForSecurityResponse, err error) {
	resp = &CloneDBInstanceForSecurityResponse{}
	err = c.Request("CloneDBInstanceForSecurity", req, resp)
	return
}

type CloneDBInstanceForSecurityArgs struct {
	ResourceOwnerId      int64
	DBInstanceStorage    int
	ResourceOwnerAccount string
	ClientToken          string
	TargetAliBid         string
	BackupId             string
	OwnerAccount         string
	OwnerId              int64
	DBInstanceClass      string
	ResourceGroupId      string
	TargetAliUid         string
	DBInstanceId         string
	PayType              string
}
type CloneDBInstanceForSecurityResponse struct {
	RequestId        string
	DBInstanceId     string
	OrderId          string
	ConnectionString string
	Port             string
}

func (c *RdsClient) CheckRecoveryConditions(req *CheckRecoveryConditionsArgs) (resp *CheckRecoveryConditionsResponse, err error) {
	resp = &CheckRecoveryConditionsResponse{}
	err = c.Request("CheckRecoveryConditions", req, resp)
	return
}

type CheckRecoveryConditionsArgs struct {
	ResourceOwnerId      int64
	RestoreTime          string
	ResourceOwnerAccount string
	BackupFile           string
	BackupId             string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type CheckRecoveryConditionsResponse struct {
	RequestId      string
	DBInstanceId   string
	RecoveryStatus string
}

func (c *RdsClient) DescribeSlowLogs(req *DescribeSlowLogsArgs) (resp *DescribeSlowLogsResponse, err error) {
	resp = &DescribeSlowLogsResponse{}
	err = c.Request("DescribeSlowLogs", req, resp)
	return
}

type DescribeSlowLogsSQLSlowLog struct {
	SlowLogId                     int64
	SQLId                         int64
	DBName                        string
	SQLText                       string
	MySQLTotalExecutionCounts     int64
	MySQLTotalExecutionTimes      int64
	TotalLockTimes                int64
	MaxLockTime                   int64
	ParseTotalRowCounts           int64
	ParseMaxRowCount              int64
	ReturnTotalRowCounts          int64
	ReturnMaxRowCount             int64
	CreateTime                    string
	SQLServerTotalExecutionCounts int64
	SQLServerTotalExecutionTimes  int64
	TotalLogicalReadCounts        int64
	TotalPhysicalReadCounts       int64
	ReportTime                    string
	MaxExecutionTime              int64
	AvgExecutionTime              int64
}
type DescribeSlowLogsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	EndTime              string
	StartTime            string
	OwnerId              int64
	PageNumber           int
	SortKey              string
	DBName               string
	PageSize             int
	DBInstanceId         string
}
type DescribeSlowLogsResponse struct {
	RequestId        string
	Engine           string
	StartTime        string
	EndTime          string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeSlowLogsSQLSlowLogList
}

type DescribeSlowLogsSQLSlowLogList []DescribeSlowLogsSQLSlowLog

func (list *DescribeSlowLogsSQLSlowLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSlowLogsSQLSlowLog)
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

func (c *RdsClient) UpgradeDBInstanceEngineVersion(req *UpgradeDBInstanceEngineVersionArgs) (resp *UpgradeDBInstanceEngineVersionResponse, err error) {
	resp = &UpgradeDBInstanceEngineVersionResponse{}
	err = c.Request("UpgradeDBInstanceEngineVersion", req, resp)
	return
}

type UpgradeDBInstanceEngineVersionArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	EngineVersion        string
	OwnerId              int64
}
type UpgradeDBInstanceEngineVersionResponse struct {
	RequestId string
	TaskId    string
}

func (c *RdsClient) DescribeTaskInfo(req *DescribeTaskInfoArgs) (resp *DescribeTaskInfoResponse, err error) {
	resp = &DescribeTaskInfoResponse{}
	err = c.Request("DescribeTaskInfo", req, resp)
	return
}

type DescribeTaskInfoArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
	TaskId               int64
}
type DescribeTaskInfoResponse struct {
	RequestId          string
	TaskId             string
	BeginTime          string
	FinishTime         string
	CreateTime         string
	TaskAction         string
	DBName             string
	TaskErrorCode      string
	Progress           string
	ExpectedFinishTime string
	TaskErrorMessage   string
	ProgressInfo       string
	Status             string
}

func (c *RdsClient) RenewInstance(req *RenewInstanceArgs) (resp *RenewInstanceResponse, err error) {
	resp = &RenewInstanceResponse{}
	err = c.Request("RenewInstance", req, resp)
	return
}

type RenewInstanceArgs struct {
	ResourceOwnerId      int64
	Period               string
	AutoPay              string
	ResourceOwnerAccount string
	ClientToken          string
	DBInstanceId         string
	OwnerId              int64
	BusinessInfo         string
}
type RenewInstanceResponse struct {
	RequestId string
	OrderId   int64
}

func (c *RdsClient) ModifyDBInstanceSpec(req *ModifyDBInstanceSpecArgs) (resp *ModifyDBInstanceSpecResponse, err error) {
	resp = &ModifyDBInstanceSpecResponse{}
	err = c.Request("ModifyDBInstanceSpec", req, resp)
	return
}

type ModifyDBInstanceSpecArgs struct {
	ResourceOwnerId      int64
	DBInstanceStorage    int
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
	PayType              string
	DBInstanceClass      string
}
type ModifyDBInstanceSpecResponse struct {
	RequestId string
}

func (c *RdsClient) ModifyDBInstanceMaintainTime(req *ModifyDBInstanceMaintainTimeArgs) (resp *ModifyDBInstanceMaintainTimeResponse, err error) {
	resp = &ModifyDBInstanceMaintainTimeResponse{}
	err = c.Request("ModifyDBInstanceMaintainTime", req, resp)
	return
}

type ModifyDBInstanceMaintainTimeArgs struct {
	MaintainTime         string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type ModifyDBInstanceMaintainTimeResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeRealtimeDiagnoses(req *DescribeRealtimeDiagnosesArgs) (resp *DescribeRealtimeDiagnosesResponse, err error) {
	resp = &DescribeRealtimeDiagnosesResponse{}
	err = c.Request("DescribeRealtimeDiagnoses", req, resp)
	return
}

type DescribeRealtimeDiagnosesRealtimeDiagnoseTasks struct {
	CreateTime  string
	TaskId      string
	HealthScore string
	Status      string
}
type DescribeRealtimeDiagnosesArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	EndTime              string
	DBInstanceId         string
	StartTime            string
	OwnerId              int64
	PageNumber           int
}
type DescribeRealtimeDiagnosesResponse struct {
	RequestId        string
	Engine           string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Tasks            DescribeRealtimeDiagnosesRealtimeDiagnoseTasksList
}

type DescribeRealtimeDiagnosesRealtimeDiagnoseTasksList []DescribeRealtimeDiagnosesRealtimeDiagnoseTasks

func (list *DescribeRealtimeDiagnosesRealtimeDiagnoseTasksList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRealtimeDiagnosesRealtimeDiagnoseTasks)
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

func (c *RdsClient) CancelImport(req *CancelImportArgs) (resp *CancelImportResponse, err error) {
	resp = &CancelImportResponse{}
	err = c.Request("CancelImport", req, resp)
	return
}

type CancelImportArgs struct {
	ResourceOwnerId      int64
	ImportId             int
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type CancelImportResponse struct {
	RequestId string
}

func (c *RdsClient) ResetAccountForPG(req *ResetAccountForPGArgs) (resp *ResetAccountForPGResponse, err error) {
	resp = &ResetAccountForPGResponse{}
	err = c.Request("ResetAccountForPG", req, resp)
	return
}

type ResetAccountForPGArgs struct {
	ResourceOwnerId      int64
	AccountPassword      string
	AccountName          string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type ResetAccountForPGResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeOptimizeAdviceOnStorage(req *DescribeOptimizeAdviceOnStorageArgs) (resp *DescribeOptimizeAdviceOnStorageResponse, err error) {
	resp = &DescribeOptimizeAdviceOnStorageResponse{}
	err = c.Request("DescribeOptimizeAdviceOnStorage", req, resp)
	return
}

type DescribeOptimizeAdviceOnStorageAdviceOnStorage struct {
	DBName        string
	TableName     string
	CurrentEngine string
	AdviseEngine  string
}
type DescribeOptimizeAdviceOnStorageArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	DBInstanceId         string
	OwnerId              int64
	PageNumber           int
}
type DescribeOptimizeAdviceOnStorageResponse struct {
	RequestId         string
	DBInstanceId      string
	TotalRecordsCount int
	PageNumber        int
	PageRecordCount   int
	Items             DescribeOptimizeAdviceOnStorageAdviceOnStorageList
}

type DescribeOptimizeAdviceOnStorageAdviceOnStorageList []DescribeOptimizeAdviceOnStorageAdviceOnStorage

func (list *DescribeOptimizeAdviceOnStorageAdviceOnStorageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOptimizeAdviceOnStorageAdviceOnStorage)
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

func (c *RdsClient) CreateDBInstance(req *CreateDBInstanceArgs) (resp *CreateDBInstanceResponse, err error) {
	resp = &CreateDBInstanceResponse{}
	err = c.Request("CreateDBInstance", req, resp)
	return
}

type CreateDBInstanceArgs struct {
	ConnectionMode        string
	ResourceOwnerId       int64
	DBInstanceStorage     int
	SystemDBCharset       string
	ClientToken           string
	EngineVersion         string
	ResourceGroupId       string
	Engine                string
	DBInstanceDescription string
	DBInstanceNetType     string
	Period                string
	ResourceOwnerAccount  string
	OwnerAccount          string
	OwnerId               int64
	UsedTime              string
	DBInstanceClass       string
	SecurityIPList        string
	VSwitchId             string
	PrivateIpAddress      string
	VPCId                 string
	ZoneId                string
	PayType               string
	InstanceNetworkType   string
}
type CreateDBInstanceResponse struct {
	RequestId        string
	DBInstanceId     string
	OrderId          string
	ConnectionString string
	Port             string
}

func (c *RdsClient) ModifyInstanceAutoRenewalAttribute(req *ModifyInstanceAutoRenewalAttributeArgs) (resp *ModifyInstanceAutoRenewalAttributeResponse, err error) {
	resp = &ModifyInstanceAutoRenewalAttributeResponse{}
	err = c.Request("ModifyInstanceAutoRenewalAttribute", req, resp)
	return
}

type ModifyInstanceAutoRenewalAttributeArgs struct {
	Duration             string
	ResourceOwnerId      int64
	AutoRenew            string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type ModifyInstanceAutoRenewalAttributeResponse struct {
	RequestId string
}

func (c *RdsClient) ResetAccountPassword(req *ResetAccountPasswordArgs) (resp *ResetAccountPasswordResponse, err error) {
	resp = &ResetAccountPasswordResponse{}
	err = c.Request("ResetAccountPassword", req, resp)
	return
}

type ResetAccountPasswordArgs struct {
	ResourceOwnerId      int64
	AccountPassword      string
	AccountName          string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type ResetAccountPasswordResponse struct {
	RequestId string
}

func (c *RdsClient) LoginDBInstancefromCloudDBA(req *LoginDBInstancefromCloudDBAArgs) (resp *LoginDBInstancefromCloudDBAResponse, err error) {
	resp = &LoginDBInstancefromCloudDBAResponse{}
	err = c.Request("LoginDBInstancefromCloudDBA", req, resp)
	return
}

type LoginDBInstancefromCloudDBAArgs struct {
	ServiceRequestParam string
	DBInstanceId        string
	ServiceRequestType  string
}
type LoginDBInstancefromCloudDBAResponse struct {
	RequestId string
	ListData  string
	AttrData  string
}

func (c *RdsClient) ModifyParameter(req *ModifyParameterArgs) (resp *ModifyParameterResponse, err error) {
	resp = &ModifyParameterResponse{}
	err = c.Request("ModifyParameter", req, resp)
	return
}

type ModifyParameterArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	Forcerestart         core.Bool
	OwnerId              int64
	Parameters           string
}
type ModifyParameterResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeOperatorPermission(req *DescribeOperatorPermissionArgs) (resp *DescribeOperatorPermissionResponse, err error) {
	resp = &DescribeOperatorPermissionResponse{}
	err = c.Request("DescribeOperatorPermission", req, resp)
	return
}

type DescribeOperatorPermissionArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeOperatorPermissionResponse struct {
	RequestId   string
	Privileges  string
	CreatedTime string
	ExpiredTime string
}

func (c *RdsClient) DeleteAccount(req *DeleteAccountArgs) (resp *DeleteAccountResponse, err error) {
	resp = &DeleteAccountResponse{}
	err = c.Request("DeleteAccount", req, resp)
	return
}

type DeleteAccountArgs struct {
	ResourceOwnerId      int64
	AccountName          string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DeleteAccountResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeSQLLogFiles(req *DescribeSQLLogFilesArgs) (resp *DescribeSQLLogFilesResponse, err error) {
	resp = &DescribeSQLLogFilesResponse{}
	err = c.Request("DescribeSQLLogFiles", req, resp)
	return
}

type DescribeSQLLogFilesLogFile struct {
	FileID         string
	LogStatus      string
	LogDownloadURL string
	LogSize        string
	LogStartTime   string
	LogEndTime     string
}
type DescribeSQLLogFilesArgs struct {
	ResourceOwnerId      int64
	FileName             string
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	DBInstanceId         string
	OwnerId              int64
	PageNumber           int
}
type DescribeSQLLogFilesResponse struct {
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeSQLLogFilesLogFileList
}

type DescribeSQLLogFilesLogFileList []DescribeSQLLogFilesLogFile

func (list *DescribeSQLLogFilesLogFileList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogFilesLogFile)
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

func (c *RdsClient) DescribeDBInstanceUser(req *DescribeDBInstanceUserArgs) (resp *DescribeDBInstanceUserResponse, err error) {
	resp = &DescribeDBInstanceUserResponse{}
	err = c.Request("DescribeDBInstanceUser", req, resp)
	return
}

type DescribeDBInstanceUserArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ConnectionString     string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeDBInstanceUserResponse struct {
	RequestId      string
	DBInstanceName string
	InternalDBFlag string
}

func (c *RdsClient) DescribeSQLCollectorPolicy(req *DescribeSQLCollectorPolicyArgs) (resp *DescribeSQLCollectorPolicyResponse, err error) {
	resp = &DescribeSQLCollectorPolicyResponse{}
	err = c.Request("DescribeSQLCollectorPolicy", req, resp)
	return
}

type DescribeSQLCollectorPolicyArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeSQLCollectorPolicyResponse struct {
	RequestId          string
	SQLCollectorStatus string
}

func (c *RdsClient) SwitchDBInstanceHA(req *SwitchDBInstanceHAArgs) (resp *SwitchDBInstanceHAResponse, err error) {
	resp = &SwitchDBInstanceHAResponse{}
	err = c.Request("SwitchDBInstanceHA", req, resp)
	return
}

type SwitchDBInstanceHAArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	Force                string
	OwnerId              int64
	NodeId               string
	Operation            string
}
type SwitchDBInstanceHAResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeResourceDiagnosis(req *DescribeResourceDiagnosisArgs) (resp *DescribeResourceDiagnosisResponse, err error) {
	resp = &DescribeResourceDiagnosisResponse{}
	err = c.Request("DescribeResourceDiagnosis", req, resp)
	return
}

type DescribeResourceDiagnosisArgs struct {
	EndTime      string
	DBInstanceId string
	StartTime    string
}
type DescribeResourceDiagnosisResponse struct {
	RequestId  string
	StartTime  string
	EndTime    string
	CPU        DescribeResourceDiagnosisCPUList
	Memory     DescribeResourceDiagnosisMemoryList
	Storage    DescribeResourceDiagnosisStorageList
	IOPS       DescribeResourceDiagnosisIOPList
	Connection DescribeResourceDiagnosisConnectionList
}

type DescribeResourceDiagnosisCPUList []string

func (list *DescribeResourceDiagnosisCPUList) UnmarshalJSON(data []byte) error {
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

type DescribeResourceDiagnosisMemoryList []string

func (list *DescribeResourceDiagnosisMemoryList) UnmarshalJSON(data []byte) error {
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

type DescribeResourceDiagnosisStorageList []string

func (list *DescribeResourceDiagnosisStorageList) UnmarshalJSON(data []byte) error {
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

type DescribeResourceDiagnosisIOPList []string

func (list *DescribeResourceDiagnosisIOPList) UnmarshalJSON(data []byte) error {
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

type DescribeResourceDiagnosisConnectionList []string

func (list *DescribeResourceDiagnosisConnectionList) UnmarshalJSON(data []byte) error {
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

func (c *RdsClient) DescribeDiagnosticReportList(req *DescribeDiagnosticReportListArgs) (resp *DescribeDiagnosticReportListResponse, err error) {
	resp = &DescribeDiagnosticReportListResponse{}
	err = c.Request("DescribeDiagnosticReportList", req, resp)
	return
}

type DescribeDiagnosticReportListReport struct {
	DiagnosticTime string
	Score          int
	StartTime      string
	EndTime        string
	DownloadURL    string
}
type DescribeDiagnosticReportListArgs struct {
	DBInstanceId string
}
type DescribeDiagnosticReportListResponse struct {
	RequestId  string
	ReportList DescribeDiagnosticReportListReportList
}

type DescribeDiagnosticReportListReportList []DescribeDiagnosticReportListReport

func (list *DescribeDiagnosticReportListReportList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDiagnosticReportListReport)
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

func (c *RdsClient) CreateBackup(req *CreateBackupArgs) (resp *CreateBackupResponse, err error) {
	resp = &CreateBackupResponse{}
	err = c.Request("CreateBackup", req, resp)
	return
}

type CreateBackupArgs struct {
	BackupMethod         string
	ResourceOwnerId      int64
	DBName               string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
	BackupType           string
}
type CreateBackupResponse struct {
	RequestId string
}

func (c *RdsClient) CreateDampPolicy(req *CreateDampPolicyArgs) (resp *CreateDampPolicyResponse, err error) {
	resp = &CreateDampPolicyResponse{}
	err = c.Request("CreateDampPolicy", req, resp)
	return
}

type CreateDampPolicyArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Priority             int
	TimeRules            string
	ActionRules          string
	SecurityToken        string
	Handlers             string
	DBInstanceId         string
	PolicyName           string
	SourceRules          string
}
type CreateDampPolicyResponse struct {
	RequestId  string
	PolicyId   string
	PolicyName string
}

func (c *RdsClient) AllocateInstancePublicConnection(req *AllocateInstancePublicConnectionArgs) (resp *AllocateInstancePublicConnectionResponse, err error) {
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
	RequestId string
}

func (c *RdsClient) ModifyDBInstanceConnectionMode(req *ModifyDBInstanceConnectionModeArgs) (resp *ModifyDBInstanceConnectionModeResponse, err error) {
	resp = &ModifyDBInstanceConnectionModeResponse{}
	err = c.Request("ModifyDBInstanceConnectionMode", req, resp)
	return
}

type ModifyDBInstanceConnectionModeArgs struct {
	ConnectionMode       string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type ModifyDBInstanceConnectionModeResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeBackups(req *DescribeBackupsArgs) (resp *DescribeBackupsResponse, err error) {
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
	BackupId             string
	OwnerAccount         string
	EndTime              string
	StartTime            string
	OwnerId              int64
	PageNumber           int
	BackupStatus         string
	BackupLocation       string
	PageSize             int
	DBInstanceId         string
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

func (c *RdsClient) CreateAccount(req *CreateAccountArgs) (resp *CreateAccountResponse, err error) {
	resp = &CreateAccountResponse{}
	err = c.Request("CreateAccount", req, resp)
	return
}

type CreateAccountArgs struct {
	ResourceOwnerId      int64
	AccountPassword      string
	AccountName          string
	ResourceOwnerAccount string
	OwnerAccount         string
	AccountType          string
	DBInstanceId         string
	OwnerId              int64
	AccountDescription   string
}
type CreateAccountResponse struct {
	RequestId string
}

func (c *RdsClient) ImportDatabaseBetweenInstances(req *ImportDatabaseBetweenInstancesArgs) (resp *ImportDatabaseBetweenInstancesResponse, err error) {
	resp = &ImportDatabaseBetweenInstancesResponse{}
	err = c.Request("ImportDatabaseBetweenInstances", req, resp)
	return
}

type ImportDatabaseBetweenInstancesArgs struct {
	ResourceOwnerId      int64
	SourceDBInstanceId   string
	ResourceOwnerAccount string
	DBInfo               string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type ImportDatabaseBetweenInstancesResponse struct {
	RequestId string
	ImportId  string
}

func (c *RdsClient) DescribeFilesForSQLServer(req *DescribeFilesForSQLServerArgs) (resp *DescribeFilesForSQLServerResponse, err error) {
	resp = &DescribeFilesForSQLServerResponse{}
	err = c.Request("DescribeFilesForSQLServer", req, resp)
	return
}

type DescribeFilesForSQLServerSQLServerUploadFile struct {
	DBName            string
	FileName          string
	FileSize          int64
	InternetFtpServer string
	InternetPort      int
	IntranetFtpserver string
	Intranetport      int
	UserName          string
	Password          string
	FileStatus        string
	Description       string
	CreationTime      string
}
type DescribeFilesForSQLServerArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	EndTime              string
	DBInstanceId         string
	StartTime            string
	OwnerId              int64
	PageNumber           int
}
type DescribeFilesForSQLServerResponse struct {
	RequestId        string
	DBInstanceId     string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeFilesForSQLServerSQLServerUploadFileList
}

type DescribeFilesForSQLServerSQLServerUploadFileList []DescribeFilesForSQLServerSQLServerUploadFile

func (list *DescribeFilesForSQLServerSQLServerUploadFileList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeFilesForSQLServerSQLServerUploadFile)
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

func (c *RdsClient) DescribeDBInstanceTDE(req *DescribeDBInstanceTDEArgs) (resp *DescribeDBInstanceTDEResponse, err error) {
	resp = &DescribeDBInstanceTDEResponse{}
	err = c.Request("DescribeDBInstanceTDE", req, resp)
	return
}

type DescribeDBInstanceTDEDatabase struct {
	DBName    string
	TDEStatus string
}
type DescribeDBInstanceTDEArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeDBInstanceTDEResponse struct {
	RequestId string
	TDEStatus string
	Databases DescribeDBInstanceTDEDatabaseList
}

type DescribeDBInstanceTDEDatabaseList []DescribeDBInstanceTDEDatabase

func (list *DescribeDBInstanceTDEDatabaseList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceTDEDatabase)
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

func (c *RdsClient) DescribeErrorLogs(req *DescribeErrorLogsArgs) (resp *DescribeErrorLogsResponse, err error) {
	resp = &DescribeErrorLogsResponse{}
	err = c.Request("DescribeErrorLogs", req, resp)
	return
}

type DescribeErrorLogsErrorLog struct {
	ErrorInfo  string
	CreateTime string
}
type DescribeErrorLogsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	EndTime              string
	DBInstanceId         string
	StartTime            string
	OwnerId              int64
	PageNumber           int
}
type DescribeErrorLogsResponse struct {
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeErrorLogsErrorLogList
}

type DescribeErrorLogsErrorLogList []DescribeErrorLogsErrorLog

func (list *DescribeErrorLogsErrorLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeErrorLogsErrorLog)
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

func (c *RdsClient) RestartDBInstance(req *RestartDBInstanceArgs) (resp *RestartDBInstanceResponse, err error) {
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

func (c *RdsClient) DescribeDBInstances(req *DescribeDBInstancesArgs) (resp *DescribeDBInstancesResponse, err error) {
	resp = &DescribeDBInstancesResponse{}
	err = c.Request("DescribeDBInstances", req, resp)
	return
}

type DescribeDBInstancesDBInstance struct {
	InsId                 int
	DBInstanceId          string
	DBInstanceDescription string
	PayType               string
	DBInstanceType        string
	RegionId              string
	ExpireTime            string
	DBInstanceStatus      string
	Engine                string
	DBInstanceNetType     string
	ConnectionMode        string
	LockMode              string
	DBInstanceClass       string
	InstanceNetworkType   string
	LockReason            string
	ZoneId                string
	MutriORsignle         core.Bool
	CreateTime            string
	EngineVersion         string
	GuardDBInstanceId     string
	TempDBInstanceId      string
	MasterInstanceId      string
	VpcId                 string
	VSwitchId             string
	ResourceGroupId       string
	ReadOnlyDBInstanceIds DescribeDBInstancesReadOnlyDBInstanceIdList
}

type DescribeDBInstancesReadOnlyDBInstanceId struct {
	DBInstanceId string
}
type DescribeDBInstancesArgs struct {
	ConnectionMode       string
	Tag4value            string
	ResourceOwnerId      int64
	Tag2key              string
	ClientToken          string
	SearchKey            string
	Tag3key              string
	PageNumber           int
	Tag1value            string
	Engine               string
	PageSize             int
	DBInstanceStatus     string
	DBInstanceId         string
	Tag3value            string
	ProxyId              string
	Tag5key              string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tag5value            string
	DBInstanceType       string
	Tags                 string
	VSwitchId            string
	Tag1key              string
	VpcId                string
	Tag2value            string
	Tag4key              string
	InstanceNetworkType  string
}
type DescribeDBInstancesResponse struct {
	RequestId        string
	PageNumber       int
	TotalRecordCount int
	PageRecordCount  int
	Items            DescribeDBInstancesDBInstanceList
}

type DescribeDBInstancesReadOnlyDBInstanceIdList []DescribeDBInstancesReadOnlyDBInstanceId

func (list *DescribeDBInstancesReadOnlyDBInstanceIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesReadOnlyDBInstanceId)
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

func (c *RdsClient) DescribeSlowLogRecords(req *DescribeSlowLogRecordsArgs) (resp *DescribeSlowLogRecordsResponse, err error) {
	resp = &DescribeSlowLogRecordsResponse{}
	err = c.Request("DescribeSlowLogRecords", req, resp)
	return
}

type DescribeSlowLogRecordsSQLSlowRecord struct {
	HostAddress        string
	DBName             string
	SQLText            string
	QueryTimes         int64
	LockTimes          int64
	ParseRowCounts     int64
	ReturnRowCounts    int64
	ExecutionStartTime string
}
type DescribeSlowLogRecordsArgs struct {
	SQLId                int64
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	EndTime              string
	StartTime            string
	OwnerId              int64
	PageNumber           int
	DBName               string
	PageSize             int
	DBInstanceId         string
}
type DescribeSlowLogRecordsResponse struct {
	RequestId        string
	Engine           string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeSlowLogRecordsSQLSlowRecordList
}

type DescribeSlowLogRecordsSQLSlowRecordList []DescribeSlowLogRecordsSQLSlowRecord

func (list *DescribeSlowLogRecordsSQLSlowRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSlowLogRecordsSQLSlowRecord)
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

func (c *RdsClient) AddTagsToResource(req *AddTagsToResourceArgs) (resp *AddTagsToResourceResponse, err error) {
	resp = &AddTagsToResourceResponse{}
	err = c.Request("AddTagsToResource", req, resp)
	return
}

type AddTagsToResourceArgs struct {
	Tag4value            string
	ResourceOwnerId      int64
	Tag2key              string
	Tag5key              string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	Tag3key              string
	OwnerId              int64
	Tag5value            string
	Tags                 string
	Tag1key              string
	Tag1value            string
	Tag2value            string
	Tag4key              string
	DBInstanceId         string
	Tag3value            string
	ProxyId              string
}
type AddTagsToResourceResponse struct {
	RequestId string
}

func (c *RdsClient) RemoveTagsFromResource(req *RemoveTagsFromResourceArgs) (resp *RemoveTagsFromResourceResponse, err error) {
	resp = &RemoveTagsFromResourceResponse{}
	err = c.Request("RemoveTagsFromResource", req, resp)
	return
}

type RemoveTagsFromResourceArgs struct {
	Tag4value            string
	ResourceOwnerId      int64
	Tag2key              string
	Tag5key              string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	Tag3key              string
	OwnerId              int64
	Tag5value            string
	Tags                 string
	Tag1key              string
	Tag1value            string
	Tag2value            string
	Tag4key              string
	DBInstanceId         string
	Tag3value            string
	ProxyId              string
}
type RemoveTagsFromResourceResponse struct {
	RequestId string
}

func (c *RdsClient) ModifySecurityIps(req *ModifySecurityIpsArgs) (resp *ModifySecurityIpsResponse, err error) {
	resp = &ModifySecurityIpsResponse{}
	err = c.Request("ModifySecurityIps", req, resp)
	return
}

type ModifySecurityIpsArgs struct {
	DBInstanceIPArrayName      string
	ResourceOwnerId            int64
	ModifyMode                 string
	ResourceOwnerAccount       string
	ClientToken                string
	OwnerAccount               string
	SecurityIps                string
	DBInstanceIPArrayAttribute string
	DBInstanceId               string
	OwnerId                    int64
}
type ModifySecurityIpsResponse struct {
	RequestId string
	TaskId    string
}

func (c *RdsClient) DescribeDBInstanceNetwork(req *DescribeDBInstanceNetworkArgs) (resp *DescribeDBInstanceNetworkResponse, err error) {
	resp = &DescribeDBInstanceNetworkResponse{}
	err = c.Request("DescribeDBInstanceNetwork", req, resp)
	return
}

type DescribeDBInstanceNetworkTopologyItem struct {
	StartPoint        string
	EndPoint          string
	NetworkTrafficIn  string
	NetworkTrafficOut string
	NetworkLatency    string
	BackendLatency    string
	NetworkErrors     string
}
type DescribeDBInstanceNetworkArgs struct {
	EndTime      string
	DBInstanceId string
	StartTime    string
}
type DescribeDBInstanceNetworkResponse struct {
	RequestId    string
	DBInstanceId string
	StartTime    string
	EndTime      string
	Topology     DescribeDBInstanceNetworkTopologyItemList
}

type DescribeDBInstanceNetworkTopologyItemList []DescribeDBInstanceNetworkTopologyItem

func (list *DescribeDBInstanceNetworkTopologyItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetworkTopologyItem)
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

func (c *RdsClient) DescribePreCheckResults(req *DescribePreCheckResultsArgs) (resp *DescribePreCheckResultsResponse, err error) {
	resp = &DescribePreCheckResultsResponse{}
	err = c.Request("DescribePreCheckResults", req, resp)
	return
}

type DescribePreCheckResultsPreCheckResult struct {
	PreCheckName   string
	PreCheckResult string
	FailReasion    string
	RepairMethod   string
}
type DescribePreCheckResultsArgs struct {
	PreCheckId           string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribePreCheckResultsResponse struct {
	RequestId    string
	DBInstanceId string
	Items        DescribePreCheckResultsPreCheckResultList
}

type DescribePreCheckResultsPreCheckResultList []DescribePreCheckResultsPreCheckResult

func (list *DescribePreCheckResultsPreCheckResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePreCheckResultsPreCheckResult)
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

func (c *RdsClient) DescribeDBInstanceHAConfig(req *DescribeDBInstanceHAConfigArgs) (resp *DescribeDBInstanceHAConfigResponse, err error) {
	resp = &DescribeDBInstanceHAConfigResponse{}
	err = c.Request("DescribeDBInstanceHAConfig", req, resp)
	return
}

type DescribeDBInstanceHAConfigNodeInfo struct {
	NodeId       string
	RegionId     string
	LogSyncTime  string
	DataSyncTime string
	NodeType     string
	ZoneId       string
	SyncStatus   string
}
type DescribeDBInstanceHAConfigArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeDBInstanceHAConfigResponse struct {
	RequestId         string
	DBInstanceId      string
	SyncMode          string
	HAMode            string
	HostInstanceInfos DescribeDBInstanceHAConfigNodeInfoList
}

type DescribeDBInstanceHAConfigNodeInfoList []DescribeDBInstanceHAConfigNodeInfo

func (list *DescribeDBInstanceHAConfigNodeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceHAConfigNodeInfo)
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

func (c *RdsClient) DescribeDBInstanceMonitor(req *DescribeDBInstanceMonitorArgs) (resp *DescribeDBInstanceMonitorResponse, err error) {
	resp = &DescribeDBInstanceMonitorResponse{}
	err = c.Request("DescribeDBInstanceMonitor", req, resp)
	return
}

type DescribeDBInstanceMonitorArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeDBInstanceMonitorResponse struct {
	RequestId string
	Period    string
}

func (c *RdsClient) DescribeOptimizeAdviceOnBigTable(req *DescribeOptimizeAdviceOnBigTableArgs) (resp *DescribeOptimizeAdviceOnBigTableResponse, err error) {
	resp = &DescribeOptimizeAdviceOnBigTableResponse{}
	err = c.Request("DescribeOptimizeAdviceOnBigTable", req, resp)
	return
}

type DescribeOptimizeAdviceOnBigTableAdviceOnBigTable struct {
	DBName    string
	TableName string
	TableSize int64
	DataSize  int64
	IndexSize int64
}
type DescribeOptimizeAdviceOnBigTableArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	DBInstanceId         string
	OwnerId              int64
	PageNumber           int
}
type DescribeOptimizeAdviceOnBigTableResponse struct {
	RequestId         string
	TotalRecordsCount int
	PageNumber        int
	PageRecordCount   int
	Items             DescribeOptimizeAdviceOnBigTableAdviceOnBigTableList
}

type DescribeOptimizeAdviceOnBigTableAdviceOnBigTableList []DescribeOptimizeAdviceOnBigTableAdviceOnBigTable

func (list *DescribeOptimizeAdviceOnBigTableAdviceOnBigTableList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOptimizeAdviceOnBigTableAdviceOnBigTable)
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

func (c *RdsClient) DescribeReplicaInitializeProgress(req *DescribeReplicaInitializeProgressArgs) (resp *DescribeReplicaInitializeProgressResponse, err error) {
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

func (c *RdsClient) ReleaseReadWriteSplittingConnection(req *ReleaseReadWriteSplittingConnectionArgs) (resp *ReleaseReadWriteSplittingConnectionResponse, err error) {
	resp = &ReleaseReadWriteSplittingConnectionResponse{}
	err = c.Request("ReleaseReadWriteSplittingConnection", req, resp)
	return
}

type ReleaseReadWriteSplittingConnectionArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type ReleaseReadWriteSplittingConnectionResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeSQLInjectionInfos(req *DescribeSQLInjectionInfosArgs) (resp *DescribeSQLInjectionInfosResponse, err error) {
	resp = &DescribeSQLInjectionInfosResponse{}
	err = c.Request("DescribeSQLInjectionInfos", req, resp)
	return
}

type DescribeSQLInjectionInfosSQLInjectionInfo struct {
	DBName         string
	SQLText        string
	LatencyTime    string
	HostAddress    string
	ExecuteTime    string
	AccountName    string
	EffectRowCount string
}
type DescribeSQLInjectionInfosArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	EndTime              string
	DBInstanceId         string
	StartTime            string
	OwnerId              int64
	PageNumber           int
}
type DescribeSQLInjectionInfosResponse struct {
	RequestId        string
	Engine           string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeSQLInjectionInfosSQLInjectionInfoList
}

type DescribeSQLInjectionInfosSQLInjectionInfoList []DescribeSQLInjectionInfosSQLInjectionInfo

func (list *DescribeSQLInjectionInfosSQLInjectionInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLInjectionInfosSQLInjectionInfo)
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

func (c *RdsClient) StartDBInstanceDiagnose(req *StartDBInstanceDiagnoseArgs) (resp *StartDBInstanceDiagnoseResponse, err error) {
	resp = &StartDBInstanceDiagnoseResponse{}
	err = c.Request("StartDBInstanceDiagnose", req, resp)
	return
}

type StartDBInstanceDiagnoseArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
	ProxyId              string
}
type StartDBInstanceDiagnoseResponse struct {
	RequestId      string
	DBInstanceName string
	DBInstanceId   string
}

func (c *RdsClient) CloneDBInstance(req *CloneDBInstanceArgs) (resp *CloneDBInstanceResponse, err error) {
	resp = &CloneDBInstanceResponse{}
	err = c.Request("CloneDBInstance", req, resp)
	return
}

type CloneDBInstanceArgs struct {
	ResourceOwnerId       int64
	RestoreTime           string
	Period                string
	DBInstanceStorage     int
	ResourceOwnerAccount  string
	ClientToken           string
	BackupId              string
	OwnerAccount          string
	OwnerId               int64
	UsedTime              string
	DBInstanceClass       string
	VSwitchId             string
	PrivateIpAddress      string
	ResourceGroupId       string
	VPCId                 string
	DBInstanceDescription string
	DBInstanceId          string
	PayType               string
	InstanceNetworkType   string
}
type CloneDBInstanceResponse struct {
	RequestId        string
	DBInstanceId     string
	OrderId          string
	ConnectionString string
	Port             string
}

func (c *RdsClient) ModifySecurityIpsForChannel(req *ModifySecurityIpsForChannelArgs) (resp *ModifySecurityIpsForChannelResponse, err error) {
	resp = &ModifySecurityIpsForChannelResponse{}
	err = c.Request("ModifySecurityIpsForChannel", req, resp)
	return
}

type ModifySecurityIpsForChannelArgs struct {
	DBInstanceIPArrayName      string
	ResourceOwnerId            int64
	ModifyMode                 string
	ResourceOwnerAccount       string
	ClientToken                string
	OwnerAccount               string
	SecurityIps                string
	DBInstanceIPArrayAttribute string
	DBInstanceId               string
	OwnerId                    int64
}
type ModifySecurityIpsForChannelResponse struct {
	RequestId string
}

func (c *RdsClient) ImportDataFromDatabase(req *ImportDataFromDatabaseArgs) (resp *ImportDataFromDatabaseResponse, err error) {
	resp = &ImportDataFromDatabaseResponse{}
	err = c.Request("ImportDataFromDatabase", req, resp)
	return
}

type ImportDataFromDatabaseArgs struct {
	ResourceOwnerId        int64
	ImportDataType         string
	ResourceOwnerAccount   string
	IsLockTable            string
	OwnerAccount           string
	SourceDatabaseDBNames  string
	SourceDatabaseIp       string
	OwnerId                int64
	SourceDatabasePassword string
	SourceDatabasePort     string
	SourceDatabaseUserName string
	DBInstanceId           string
}
type ImportDataFromDatabaseResponse struct {
	RequestId string
	ImportId  int
}

func (c *RdsClient) ModifyDBInstanceNetworkType(req *ModifyDBInstanceNetworkTypeArgs) (resp *ModifyDBInstanceNetworkTypeResponse, err error) {
	resp = &ModifyDBInstanceNetworkTypeResponse{}
	err = c.Request("ModifyDBInstanceNetworkType", req, resp)
	return
}

type ModifyDBInstanceNetworkTypeArgs struct {
	ResourceOwnerId                      int64
	ResourceOwnerAccount                 string
	OwnerAccount                         string
	OwnerId                              int64
	VSwitchId                            string
	PrivateIpAddress                     string
	RetainClassic                        string
	ClassicExpiredDays                   string
	VPCId                                string
	DBInstanceId                         string
	ReadWriteSplittingPrivateIpAddress   string
	InstanceNetworkType                  string
	ReadWriteSplittingClassicExpiredDays int
}
type ModifyDBInstanceNetworkTypeResponse struct {
	RequestId string
	TaskId    string
}

func (c *RdsClient) DescribeInstanceAutoRenewAttribute(req *DescribeInstanceAutoRenewAttributeArgs) (resp *DescribeInstanceAutoRenewAttributeResponse, err error) {
	resp = &DescribeInstanceAutoRenewAttributeResponse{}
	err = c.Request("DescribeInstanceAutoRenewAttribute", req, resp)
	return
}

type DescribeInstanceAutoRenewAttributeItem struct {
	DBInstanceId string
	RegionId     string
	Duration     int
	Status       string
	AutoRenew    string
}
type DescribeInstanceAutoRenewAttributeArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	PageSize             int
	DBInstanceId         string
	OwnerId              int64
	PageNumber           int
	ProxyId              string
}
type DescribeInstanceAutoRenewAttributeResponse struct {
	RequestId        string
	PageNumber       int
	TotalRecordCount int
	PageRecordCount  int
	Items            DescribeInstanceAutoRenewAttributeItemList
}

type DescribeInstanceAutoRenewAttributeItemList []DescribeInstanceAutoRenewAttributeItem

func (list *DescribeInstanceAutoRenewAttributeItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceAutoRenewAttributeItem)
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

func (c *RdsClient) CreateSQLDiagnosis(req *CreateSQLDiagnosisArgs) (resp *CreateSQLDiagnosisResponse, err error) {
	resp = &CreateSQLDiagnosisResponse{}
	err = c.Request("CreateSQLDiagnosis", req, resp)
	return
}

type CreateSQLDiagnosisArgs struct {
	EndTime      string
	DBInstanceId string
	StartTime    string
}
type CreateSQLDiagnosisResponse struct {
	RequestId string
	SQLDiagId string
}

func (c *RdsClient) DescribeParameterTemplates(req *DescribeParameterTemplatesArgs) (resp *DescribeParameterTemplatesResponse, err error) {
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
	ClientToken          string
	Engine               string
	OwnerAccount         string
	EngineVersion        string
	OwnerId              int64
}
type DescribeParameterTemplatesResponse struct {
	RequestId      string
	Engine         string
	EngineVersion  string
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

func (c *RdsClient) ModifyDBInstanceConnectionString(req *ModifyDBInstanceConnectionStringArgs) (resp *ModifyDBInstanceConnectionStringResponse, err error) {
	resp = &ModifyDBInstanceConnectionStringResponse{}
	err = c.Request("ModifyDBInstanceConnectionString", req, resp)
	return
}

type ModifyDBInstanceConnectionStringArgs struct {
	ResourceOwnerId         int64
	ConnectionStringPrefix  string
	ResourceOwnerAccount    string
	Port                    string
	OwnerAccount            string
	DBInstanceId            string
	OwnerId                 int64
	CurrentConnectionString string
}
type ModifyDBInstanceConnectionStringResponse struct {
	RequestId string
}

func (c *RdsClient) ModifyDBInstanceMonitor(req *ModifyDBInstanceMonitorArgs) (resp *ModifyDBInstanceMonitorResponse, err error) {
	resp = &ModifyDBInstanceMonitorResponse{}
	err = c.Request("ModifyDBInstanceMonitor", req, resp)
	return
}

type ModifyDBInstanceMonitorArgs struct {
	ResourceOwnerId      int64
	Period               string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type ModifyDBInstanceMonitorResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeAbnormalDBInstances(req *DescribeAbnormalDBInstancesArgs) (resp *DescribeAbnormalDBInstancesResponse, err error) {
	resp = &DescribeAbnormalDBInstancesResponse{}
	err = c.Request("DescribeAbnormalDBInstances", req, resp)
	return
}

type DescribeAbnormalDBInstancesInstanceResult struct {
	DBInstanceDescription string
	DBInstanceId          string
	AbnormalItems         DescribeAbnormalDBInstancesAbnormalItemList
}

type DescribeAbnormalDBInstancesAbnormalItem struct {
	CheckTime      string
	CheckItem      string
	AbnormalReason string
	AbnormalValue  string
	AbnormalDetail string
	AdviceKey      string
	AdviseValue    DescribeAbnormalDBInstancesAdviseValueList
}
type DescribeAbnormalDBInstancesArgs struct {
	Tag4value            string
	ResourceOwnerId      int64
	Tag2key              string
	Tag5key              string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	Tag3key              string
	OwnerId              int64
	Tag5value            string
	PageNumber           int
	Tags                 string
	Tag1key              string
	Tag1value            string
	Tag2value            string
	PageSize             int
	Tag4key              string
	DBInstanceId         string
	Tag3value            string
	ProxyId              string
}
type DescribeAbnormalDBInstancesResponse struct {
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeAbnormalDBInstancesInstanceResultList
}

type DescribeAbnormalDBInstancesAbnormalItemList []DescribeAbnormalDBInstancesAbnormalItem

func (list *DescribeAbnormalDBInstancesAbnormalItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAbnormalDBInstancesAbnormalItem)
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

type DescribeAbnormalDBInstancesAdviseValueList []string

func (list *DescribeAbnormalDBInstancesAdviseValueList) UnmarshalJSON(data []byte) error {
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

type DescribeAbnormalDBInstancesInstanceResultList []DescribeAbnormalDBInstancesInstanceResult

func (list *DescribeAbnormalDBInstancesInstanceResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAbnormalDBInstancesInstanceResult)
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

func (c *RdsClient) ModifyDBInstanceNetExpireTime(req *ModifyDBInstanceNetExpireTimeArgs) (resp *ModifyDBInstanceNetExpireTimeResponse, err error) {
	resp = &ModifyDBInstanceNetExpireTimeResponse{}
	err = c.Request("ModifyDBInstanceNetExpireTime", req, resp)
	return
}

type ModifyDBInstanceNetExpireTimeArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ConnectionString     string
	ClassicExpiredDays   int
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type ModifyDBInstanceNetExpireTimeResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeSQLLogReports(req *DescribeSQLLogReportsArgs) (resp *DescribeSQLLogReportsResponse, err error) {
	resp = &DescribeSQLLogReportsResponse{}
	err = c.Request("DescribeSQLLogReports", req, resp)
	return
}

type DescribeSQLLogReportsItem struct {
	ReportTime       string
	LatencyTopNItems DescribeSQLLogReportsLatencyTopNItemList
	QPSTopNItems     DescribeSQLLogReportsQPSTopNItemList
}

type DescribeSQLLogReportsLatencyTopNItem struct {
	SQLText         string
	AvgLatency      int64
	SQLExecuteTimes int64
}

type DescribeSQLLogReportsQPSTopNItem struct {
	SQLText         string
	SQLExecuteTimes int64
}
type DescribeSQLLogReportsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	EndTime              string
	DBInstanceId         string
	StartTime            string
	OwnerId              int64
	PageNumber           int
}
type DescribeSQLLogReportsResponse struct {
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeSQLLogReportsItemList
}

type DescribeSQLLogReportsLatencyTopNItemList []DescribeSQLLogReportsLatencyTopNItem

func (list *DescribeSQLLogReportsLatencyTopNItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogReportsLatencyTopNItem)
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

type DescribeSQLLogReportsQPSTopNItemList []DescribeSQLLogReportsQPSTopNItem

func (list *DescribeSQLLogReportsQPSTopNItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogReportsQPSTopNItem)
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

type DescribeSQLLogReportsItemList []DescribeSQLLogReportsItem

func (list *DescribeSQLLogReportsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogReportsItem)
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

func (c *RdsClient) GrantAccountPrivilege(req *GrantAccountPrivilegeArgs) (resp *GrantAccountPrivilegeResponse, err error) {
	resp = &GrantAccountPrivilegeResponse{}
	err = c.Request("GrantAccountPrivilege", req, resp)
	return
}

type GrantAccountPrivilegeArgs struct {
	ResourceOwnerId      int64
	AccountName          string
	DBName               string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
	AccountPrivilege     string
}
type GrantAccountPrivilegeResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeSQLLogReportList(req *DescribeSQLLogReportListArgs) (resp *DescribeSQLLogReportListResponse, err error) {
	resp = &DescribeSQLLogReportListResponse{}
	err = c.Request("DescribeSQLLogReportList", req, resp)
	return
}

type DescribeSQLLogReportListItem struct {
	ReportTime       string
	LatencyTopNItems DescribeSQLLogReportListLatencyTopNItemList
	QPSTopNItems     DescribeSQLLogReportListQPSTopNItemList
}

type DescribeSQLLogReportListLatencyTopNItem struct {
	SQLText         string
	AvgLatency      int64
	SQLExecuteTimes int64
}

type DescribeSQLLogReportListQPSTopNItem struct {
	SQLText         string
	SQLExecuteTimes int64
}
type DescribeSQLLogReportListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	EndTime              string
	DBInstanceId         string
	StartTime            string
	OwnerId              int64
	PageNumber           int
}
type DescribeSQLLogReportListResponse struct {
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeSQLLogReportListItemList
}

type DescribeSQLLogReportListLatencyTopNItemList []DescribeSQLLogReportListLatencyTopNItem

func (list *DescribeSQLLogReportListLatencyTopNItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogReportListLatencyTopNItem)
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

type DescribeSQLLogReportListQPSTopNItemList []DescribeSQLLogReportListQPSTopNItem

func (list *DescribeSQLLogReportListQPSTopNItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogReportListQPSTopNItem)
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

type DescribeSQLLogReportListItemList []DescribeSQLLogReportListItem

func (list *DescribeSQLLogReportListItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogReportListItem)
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

func (c *RdsClient) StartArchiveSQLLog(req *StartArchiveSQLLogArgs) (resp *StartArchiveSQLLogResponse, err error) {
	resp = &StartArchiveSQLLogResponse{}
	err = c.Request("StartArchiveSQLLog", req, resp)
	return
}

type StartArchiveSQLLogArgs struct {
	ResourceOwnerId      int64
	Database             string
	ResourceOwnerAccount string
	OwnerAccount         string
	EndTime              string
	DBInstanceId         string
	StartTime            string
	OwnerId              int64
	User                 string
	QueryKeywords        string
}
type StartArchiveSQLLogResponse struct {
	RequestId string
}

func (c *RdsClient) ModifyDBInstanceDescription(req *ModifyDBInstanceDescriptionArgs) (resp *ModifyDBInstanceDescriptionResponse, err error) {
	resp = &ModifyDBInstanceDescriptionResponse{}
	err = c.Request("ModifyDBInstanceDescription", req, resp)
	return
}

type ModifyDBInstanceDescriptionArgs struct {
	ResourceOwnerId       int64
	ResourceOwnerAccount  string
	ClientToken           string
	OwnerAccount          string
	DBInstanceId          string
	DBInstanceDescription string
	OwnerId               int64
}
type ModifyDBInstanceDescriptionResponse struct {
	RequestId string
}

func (c *RdsClient) ModifySQLCollectorPolicy(req *ModifySQLCollectorPolicyArgs) (resp *ModifySQLCollectorPolicyResponse, err error) {
	resp = &ModifySQLCollectorPolicyResponse{}
	err = c.Request("ModifySQLCollectorPolicy", req, resp)
	return
}

type ModifySQLCollectorPolicyArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	SQLCollectorStatus   string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type ModifySQLCollectorPolicyResponse struct {
	RequestId string
}

func (c *RdsClient) PreCheckBeforeImportData(req *PreCheckBeforeImportDataArgs) (resp *PreCheckBeforeImportDataResponse, err error) {
	resp = &PreCheckBeforeImportDataResponse{}
	err = c.Request("PreCheckBeforeImportData", req, resp)
	return
}

type PreCheckBeforeImportDataArgs struct {
	ResourceOwnerId        int64
	ImportDataType         string
	ResourceOwnerAccount   string
	ClientToken            string
	OwnerAccount           string
	SourceDatabaseDBNames  string
	SourceDatabaseIp       string
	OwnerId                int64
	SourceDatabasePassword string
	SourceDatabasePort     string
	SourceDatabaseUserName string
	DBInstanceId           string
}
type PreCheckBeforeImportDataResponse struct {
	RequestId  string
	PreCheckId string
}

func (c *RdsClient) DescribeResourceUsage(req *DescribeResourceUsageArgs) (resp *DescribeResourceUsageResponse, err error) {
	resp = &DescribeResourceUsageResponse{}
	err = c.Request("DescribeResourceUsage", req, resp)
	return
}

type DescribeResourceUsageArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeResourceUsageResponse struct {
	RequestId      string
	DBInstanceId   string
	Engine         string
	DiskUsed       int64
	DataSize       int64
	LogSize        int64
	BackupSize     int64
	SQLSize        int64
	ColdBackupSize int64
}

func (c *RdsClient) DescribeDBInstanceIPArrayList(req *DescribeDBInstanceIPArrayListArgs) (resp *DescribeDBInstanceIPArrayListResponse, err error) {
	resp = &DescribeDBInstanceIPArrayListResponse{}
	err = c.Request("DescribeDBInstanceIPArrayList", req, resp)
	return
}

type DescribeDBInstanceIPArrayListDBInstanceIPArray struct {
	DBInstanceIPArrayName      string
	DBInstanceIPArrayAttribute string
	SecurityIPList             string
}
type DescribeDBInstanceIPArrayListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeDBInstanceIPArrayListResponse struct {
	RequestId string
	Items     DescribeDBInstanceIPArrayListDBInstanceIPArrayList
}

type DescribeDBInstanceIPArrayListDBInstanceIPArrayList []DescribeDBInstanceIPArrayListDBInstanceIPArray

func (list *DescribeDBInstanceIPArrayListDBInstanceIPArrayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceIPArrayListDBInstanceIPArray)
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

func (c *RdsClient) SwitchDBInstanceChargeType(req *SwitchDBInstanceChargeTypeArgs) (resp *SwitchDBInstanceChargeTypeResponse, err error) {
	resp = &SwitchDBInstanceChargeTypeResponse{}
	err = c.Request("SwitchDBInstanceChargeType", req, resp)
	return
}

type SwitchDBInstanceChargeTypeArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type SwitchDBInstanceChargeTypeResponse struct {
	RequestId string
}

func (c *RdsClient) ModifyInstanceAutoRenewAttribute(req *ModifyInstanceAutoRenewAttributeArgs) (resp *ModifyInstanceAutoRenewAttributeResponse, err error) {
	resp = &ModifyInstanceAutoRenewAttributeResponse{}
	err = c.Request("ModifyInstanceAutoRenewAttribute", req, resp)
	return
}

type ModifyInstanceAutoRenewAttributeArgs struct {
	Duration             string
	ResourceOwnerId      int64
	AutoRenew            string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type ModifyInstanceAutoRenewAttributeResponse struct {
	RequestId string
}

func (c *RdsClient) ModifyDBInstanceSSL(req *ModifyDBInstanceSSLArgs) (resp *ModifyDBInstanceSSLResponse, err error) {
	resp = &ModifyDBInstanceSSLResponse{}
	err = c.Request("ModifyDBInstanceSSL", req, resp)
	return
}

type ModifyDBInstanceSSLArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ConnectionString     string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type ModifyDBInstanceSSLResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeModifyParameterLog(req *DescribeModifyParameterLogArgs) (resp *DescribeModifyParameterLogResponse, err error) {
	resp = &DescribeModifyParameterLogResponse{}
	err = c.Request("DescribeModifyParameterLog", req, resp)
	return
}

type DescribeModifyParameterLogParameterChangeLog struct {
	ModifyTime        string
	OldParameterValue string
	NewParameterValue string
	ParameterName     string
	Status            string
}
type DescribeModifyParameterLogArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	EndTime              string
	DBInstanceId         string
	StartTime            string
	OwnerId              int64
	PageNumber           int
}
type DescribeModifyParameterLogResponse struct {
	RequestId        string
	Engine           string
	DBInstanceId     string
	EngineVersion    string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeModifyParameterLogParameterChangeLogList
}

type DescribeModifyParameterLogParameterChangeLogList []DescribeModifyParameterLogParameterChangeLog

func (list *DescribeModifyParameterLogParameterChangeLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeModifyParameterLogParameterChangeLog)
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

func (c *RdsClient) ReleaseInstancePublicConnection(req *ReleaseInstancePublicConnectionArgs) (resp *ReleaseInstancePublicConnectionResponse, err error) {
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

func (c *RdsClient) ReleaseReplica(req *ReleaseReplicaArgs) (resp *ReleaseReplicaResponse, err error) {
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

func (c *RdsClient) DescribeDampPoliciesByCid(req *DescribeDampPoliciesByCidArgs) (resp *DescribeDampPoliciesByCidResponse, err error) {
	resp = &DescribeDampPoliciesByCidResponse{}
	err = c.Request("DescribeDampPoliciesByCid", req, resp)
	return
}

type DescribeDampPoliciesByCidPolicy struct {
	PolicyName string
	Comment    string
}
type DescribeDampPoliciesByCidArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeDampPoliciesByCidResponse struct {
	RequestId string
	Policies  DescribeDampPoliciesByCidPolicyList
}

type DescribeDampPoliciesByCidPolicyList []DescribeDampPoliciesByCidPolicy

func (list *DescribeDampPoliciesByCidPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDampPoliciesByCidPolicy)
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

func (c *RdsClient) DeleteDatabase(req *DeleteDatabaseArgs) (resp *DeleteDatabaseResponse, err error) {
	resp = &DeleteDatabaseResponse{}
	err = c.Request("DeleteDatabase", req, resp)
	return
}

type DeleteDatabaseArgs struct {
	ResourceOwnerId      int64
	DBName               string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DeleteDatabaseResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeTasks(req *DescribeTasksArgs) (resp *DescribeTasksResponse, err error) {
	resp = &DescribeTasksResponse{}
	err = c.Request("DescribeTasks", req, resp)
	return
}

type DescribeTasksTaskProgressInfo struct {
	DBName             string
	BeginTime          string
	ProgressInfo       string
	FinishTime         string
	TaskAction         string
	TaskId             string
	Progress           string
	ExpectedFinishTime string
	Status             string
	TaskErrorCode      string
	TaskErrorMessage   string
}
type DescribeTasksArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	EndTime              string
	StartTime            string
	OwnerId              int64
	PageNumber           int
	PageSize             int
	DBInstanceId         string
	TaskAction           string
	Status               string
}
type DescribeTasksResponse struct {
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeTasksTaskProgressInfoList
}

type DescribeTasksTaskProgressInfoList []DescribeTasksTaskProgressInfo

func (list *DescribeTasksTaskProgressInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTasksTaskProgressInfo)
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

func (c *RdsClient) CalculateDBInstanceWeight(req *CalculateDBInstanceWeightArgs) (resp *CalculateDBInstanceWeightResponse, err error) {
	resp = &CalculateDBInstanceWeightResponse{}
	err = c.Request("CalculateDBInstanceWeight", req, resp)
	return
}

type CalculateDBInstanceWeightDBInstanceWeight struct {
	DBInstanceId   string
	DBInstanceType string
	Availability   string
	Weight         string
}
type CalculateDBInstanceWeightArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type CalculateDBInstanceWeightResponse struct {
	RequestId string
	Items     CalculateDBInstanceWeightDBInstanceWeightList
}

type CalculateDBInstanceWeightDBInstanceWeightList []CalculateDBInstanceWeightDBInstanceWeight

func (list *CalculateDBInstanceWeightDBInstanceWeightList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CalculateDBInstanceWeightDBInstanceWeight)
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

func (c *RdsClient) DescribeDBInstancesByExpireTime(req *DescribeDBInstancesByExpireTimeArgs) (resp *DescribeDBInstancesByExpireTimeResponse, err error) {
	resp = &DescribeDBInstancesByExpireTimeResponse{}
	err = c.Request("DescribeDBInstancesByExpireTime", req, resp)
	return
}

type DescribeDBInstancesByExpireTimeDBInstanceExpireTime struct {
	DBInstanceId          string
	DBInstanceDescription string
	ExpireTime            string
	DBInstanceStatus      string
	LockMode              string
}
type DescribeDBInstancesByExpireTimeArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	PageNumber           int
	Tags                 string
	Expired              core.Bool
	PageSize             int
	ExpirePeriod         int
	ProxyId              string
}
type DescribeDBInstancesByExpireTimeResponse struct {
	RequestId        string
	PageNumber       int
	TotalRecordCount int
	PageRecordCount  int
	Items            DescribeDBInstancesByExpireTimeDBInstanceExpireTimeList
}

type DescribeDBInstancesByExpireTimeDBInstanceExpireTimeList []DescribeDBInstancesByExpireTimeDBInstanceExpireTime

func (list *DescribeDBInstancesByExpireTimeDBInstanceExpireTimeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesByExpireTimeDBInstanceExpireTime)
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

func (c *RdsClient) DescribeReplicas(req *DescribeReplicasArgs) (resp *DescribeReplicasResponse, err error) {
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

func (c *RdsClient) RestoreDBInstance(req *RestoreDBInstanceArgs) (resp *RestoreDBInstanceResponse, err error) {
	resp = &RestoreDBInstanceResponse{}
	err = c.Request("RestoreDBInstance", req, resp)
	return
}

type RestoreDBInstanceArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	BackupId             string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type RestoreDBInstanceResponse struct {
	RequestId string
}

func (c *RdsClient) DeleteBackup(req *DeleteBackupArgs) (resp *DeleteBackupResponse, err error) {
	resp = &DeleteBackupResponse{}
	err = c.Request("DeleteBackup", req, resp)
	return
}

type DeleteBackupArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	BackupId             string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DeleteBackupResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeSQLDiagnosis(req *DescribeSQLDiagnosisArgs) (resp *DescribeSQLDiagnosisResponse, err error) {
	resp = &DescribeSQLDiagnosisResponse{}
	err = c.Request("DescribeSQLDiagnosis", req, resp)
	return
}

type DescribeSQLDiagnosisArgs struct {
	SQLDiagId    string
	DBInstanceId string
}
type DescribeSQLDiagnosisResponse struct {
	RequestId string
	SQLList   DescribeSQLDiagnosisSQLListList
}

type DescribeSQLDiagnosisSQLListList []string

func (list *DescribeSQLDiagnosisSQLListList) UnmarshalJSON(data []byte) error {
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

func (c *RdsClient) DescribeVpcZoneNos(req *DescribeVpcZoneNosArgs) (resp *DescribeVpcZoneNosResponse, err error) {
	resp = &DescribeVpcZoneNosResponse{}
	err = c.Request("DescribeVpcZoneNos", req, resp)
	return
}

type DescribeVpcZoneNosVpcZoneId struct {
	ZoneId    string
	Region    string
	SubDomain string
}
type DescribeVpcZoneNosArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	ZoneId               string
	OwnerId              int64
	Region               string
}
type DescribeVpcZoneNosResponse struct {
	RequestId string
	Items     DescribeVpcZoneNosVpcZoneIdList
}

type DescribeVpcZoneNosVpcZoneIdList []DescribeVpcZoneNosVpcZoneId

func (list *DescribeVpcZoneNosVpcZoneIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVpcZoneNosVpcZoneId)
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

func (c *RdsClient) DescribeDBInstanceNetInfo(req *DescribeDBInstanceNetInfoArgs) (resp *DescribeDBInstanceNetInfoResponse, err error) {
	resp = &DescribeDBInstanceNetInfoResponse{}
	err = c.Request("DescribeDBInstanceNetInfo", req, resp)
	return
}

type DescribeDBInstanceNetInfoDBInstanceNetInfo struct {
	Upgradeable          string
	ExpiredTime          string
	ConnectionString     string
	IPAddress            string
	IPType               string
	Port                 string
	VPCId                string
	VSwitchId            string
	ConnectionStringType string
	MaxDelayTime         string
	DistributionType     string
	SecurityIPGroups     DescribeDBInstanceNetInfoSecurityIPGroupList
	DBInstanceWeights    DescribeDBInstanceNetInfoDBInstanceWeightList
}

type DescribeDBInstanceNetInfoSecurityIPGroup struct {
	SecurityIPGroupName string
	SecurityIPs         string
}

type DescribeDBInstanceNetInfoDBInstanceWeight struct {
	DBInstanceId   string
	DBInstanceType string
	Availability   string
	Weight         string
}
type DescribeDBInstanceNetInfoArgs struct {
	ResourceOwnerId          int64
	Flag                     string
	DBInstanceNetRWSplitType string
	ResourceOwnerAccount     string
	ClientToken              string
	OwnerAccount             string
	DBInstanceId             string
	OwnerId                  int64
}
type DescribeDBInstanceNetInfoResponse struct {
	RequestId           string
	InstanceNetworkType string
	DBInstanceNetInfos  DescribeDBInstanceNetInfoDBInstanceNetInfoList
}

type DescribeDBInstanceNetInfoSecurityIPGroupList []DescribeDBInstanceNetInfoSecurityIPGroup

func (list *DescribeDBInstanceNetInfoSecurityIPGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetInfoSecurityIPGroup)
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

type DescribeDBInstanceNetInfoDBInstanceWeightList []DescribeDBInstanceNetInfoDBInstanceWeight

func (list *DescribeDBInstanceNetInfoDBInstanceWeightList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetInfoDBInstanceWeight)
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

func (c *RdsClient) DescribeDBInstanceSSL(req *DescribeDBInstanceSSLArgs) (resp *DescribeDBInstanceSSLResponse, err error) {
	resp = &DescribeDBInstanceSSLResponse{}
	err = c.Request("DescribeDBInstanceSSL", req, resp)
	return
}

type DescribeDBInstanceSSLArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeDBInstanceSSLResponse struct {
	RequestId           string
	ConnectionString    string
	SSLExpireTime       string
	RequireUpdate       string
	RequireUpdateReason string
}

func (c *RdsClient) ModifyDBDescription(req *ModifyDBDescriptionArgs) (resp *ModifyDBDescriptionResponse, err error) {
	resp = &ModifyDBDescriptionResponse{}
	err = c.Request("ModifyDBDescription", req, resp)
	return
}

type ModifyDBDescriptionArgs struct {
	ResourceOwnerId      int64
	DBName               string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	DBDescription        string
	OwnerId              int64
}
type ModifyDBDescriptionResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeRenewalPrice(req *DescribeRenewalPriceArgs) (resp *DescribeRenewalPriceResponse, err error) {
	resp = &DescribeRenewalPriceResponse{}
	err = c.Request("DescribeRenewalPrice", req, resp)
	return
}

type DescribeRenewalPriceRule struct {
	RuleId      int64
	Name        string
	Description string
}

type DescribeRenewalPricePriceInfo struct {
	Currency      string
	OriginalPrice float32
	TradePrice    float32
	DiscountPrice float32
	Coupons       DescribeRenewalPriceCouponList
	RuleIds       DescribeRenewalPriceRuleIdList
	ActivityInfo  DescribeRenewalPriceActivityInfo
}

type DescribeRenewalPriceCoupon struct {
	CouponNo    string
	Name        string
	Description string
	IsSelected  string
}

type DescribeRenewalPriceActivityInfo struct {
	CheckErrMsg string
	ErrorCode   string
	Success     string
}
type DescribeRenewalPriceArgs struct {
	ResourceOwnerId      int64
	Quantity             int
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	CommodityCode        string
	OwnerId              int64
	UsedTime             string
	DBInstanceClass      string
	PromotionCode        string
	DBInstanceId         string
	TimeType             string
	PayType              string
	BusinessInfo         string
	OrderType            string
}
type DescribeRenewalPriceResponse struct {
	RequestId string
	Rules     DescribeRenewalPriceRuleList
	PriceInfo DescribeRenewalPricePriceInfo
}

type DescribeRenewalPriceCouponList []DescribeRenewalPriceCoupon

func (list *DescribeRenewalPriceCouponList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRenewalPriceCoupon)
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

type DescribeRenewalPriceRuleIdList []string

func (list *DescribeRenewalPriceRuleIdList) UnmarshalJSON(data []byte) error {
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

type DescribeRenewalPriceRuleList []DescribeRenewalPriceRule

func (list *DescribeRenewalPriceRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRenewalPriceRule)
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

func (c *RdsClient) SwitchDBInstanceNetType(req *SwitchDBInstanceNetTypeArgs) (resp *SwitchDBInstanceNetTypeResponse, err error) {
	resp = &SwitchDBInstanceNetTypeResponse{}
	err = c.Request("SwitchDBInstanceNetType", req, resp)
	return
}

type SwitchDBInstanceNetTypeArgs struct {
	ResourceOwnerId        int64
	ConnectionStringPrefix string
	ConnectionStringType   string
	ResourceOwnerAccount   string
	ClientToken            string
	Port                   string
	OwnerAccount           string
	DBInstanceId           string
	OwnerId                int64
}
type SwitchDBInstanceNetTypeResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeBackupSetsForSecurity(req *DescribeBackupSetsForSecurityArgs) (resp *DescribeBackupSetsForSecurityResponse, err error) {
	resp = &DescribeBackupSetsForSecurityResponse{}
	err = c.Request("DescribeBackupSetsForSecurity", req, resp)
	return
}

type DescribeBackupSetsForSecurityBackup struct {
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
type DescribeBackupSetsForSecurityArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	TargetAliBid         string
	BackupId             string
	OwnerAccount         string
	EndTime              string
	StartTime            string
	OwnerId              int64
	PageNumber           int
	BackupStatus         string
	BackupLocation       string
	TargetAliUid         string
	PageSize             int
	DBInstanceId         string
	BackupMode           string
}
type DescribeBackupSetsForSecurityResponse struct {
	RequestId        string
	TotalRecordCount string
	PageNumber       string
	PageRecordCount  string
	TotalBackupSize  int64
	Items            DescribeBackupSetsForSecurityBackupList
}

type DescribeBackupSetsForSecurityBackupList []DescribeBackupSetsForSecurityBackup

func (list *DescribeBackupSetsForSecurityBackupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBackupSetsForSecurityBackup)
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

func (c *RdsClient) ModifyReadWriteSplittingConnection(req *ModifyReadWriteSplittingConnectionArgs) (resp *ModifyReadWriteSplittingConnectionResponse, err error) {
	resp = &ModifyReadWriteSplittingConnectionResponse{}
	err = c.Request("ModifyReadWriteSplittingConnection", req, resp)
	return
}

type ModifyReadWriteSplittingConnectionArgs struct {
	ResourceOwnerId        int64
	ConnectionStringPrefix string
	ResourceOwnerAccount   string
	Port                   string
	DistributionType       string
	OwnerAccount           string
	Weight                 string
	DBInstanceId           string
	OwnerId                int64
	MaxDelayTime           string
}
type ModifyReadWriteSplittingConnectionResponse struct {
	RequestId string
}

func (c *RdsClient) AddBuDBInstanceRelation(req *AddBuDBInstanceRelationArgs) (resp *AddBuDBInstanceRelationResponse, err error) {
	resp = &AddBuDBInstanceRelationResponse{}
	err = c.Request("AddBuDBInstanceRelation", req, resp)
	return
}

type AddBuDBInstanceRelationArgs struct {
	BusinessUnit         string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	DBInstanceId         string
	OwnerId              int64
}
type AddBuDBInstanceRelationResponse struct {
	RequestId      string
	BusinessUnit   string
	DBInstanceName string
}

func (c *RdsClient) UpgradeDBInstanceNetWorkInfo(req *UpgradeDBInstanceNetWorkInfoArgs) (resp *UpgradeDBInstanceNetWorkInfoResponse, err error) {
	resp = &UpgradeDBInstanceNetWorkInfoResponse{}
	err = c.Request("UpgradeDBInstanceNetWorkInfo", req, resp)
	return
}

type UpgradeDBInstanceNetWorkInfoArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ConnectionString     string
	DBInstanceId         string
	OwnerId              int64
}
type UpgradeDBInstanceNetWorkInfoResponse struct {
	RequestId string
}

func (c *RdsClient) ModifyDBInstancePayType(req *ModifyDBInstancePayTypeArgs) (resp *ModifyDBInstancePayTypeResponse, err error) {
	resp = &ModifyDBInstancePayTypeResponse{}
	err = c.Request("ModifyDBInstancePayType", req, resp)
	return
}

type ModifyDBInstancePayTypeArgs struct {
	ResourceOwnerId      int64
	Period               string
	AgentId              string
	AutoPay              string
	ResourceOwnerAccount string
	ClientToken          string
	Resource             string
	OwnerAccount         string
	OwnerId              int64
	UsedTime             int
	DBInstanceId         string
	PayType              string
	BusinessInfo         string
}
type ModifyDBInstancePayTypeResponse struct {
	RequestId    string
	DBInstanceId string
	OrderId      int64
}

func (c *RdsClient) CheckDBNameAvailable(req *CheckDBNameAvailableArgs) (resp *CheckDBNameAvailableResponse, err error) {
	resp = &CheckDBNameAvailableResponse{}
	err = c.Request("CheckDBNameAvailable", req, resp)
	return
}

type CheckDBNameAvailableArgs struct {
	ResourceOwnerId      int64
	DBName               string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type CheckDBNameAvailableResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeDampPolicyByPolicyName(req *DescribeDampPolicyByPolicyNameArgs) (resp *DescribeDampPolicyByPolicyNameResponse, err error) {
	resp = &DescribeDampPolicyByPolicyNameResponse{}
	err = c.Request("DescribeDampPolicyByPolicyName", req, resp)
	return
}

type DescribeDampPolicyByPolicyNameArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	PolicyName           string
	OwnerId              int64
}
type DescribeDampPolicyByPolicyNameResponse struct {
	RequestId   string
	Policy      string
	TimeRules   string
	ActionRules string
	SourceRules string
	Handler     string
}

func (c *RdsClient) ModifyPostpaidDBInstanceSpec(req *ModifyPostpaidDBInstanceSpecArgs) (resp *ModifyPostpaidDBInstanceSpecResponse, err error) {
	resp = &ModifyPostpaidDBInstanceSpecResponse{}
	err = c.Request("ModifyPostpaidDBInstanceSpec", req, resp)
	return
}

type ModifyPostpaidDBInstanceSpecArgs struct {
	ResourceOwnerId      int64
	DBInstanceStorage    int
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
	DBInstanceClass      string
}
type ModifyPostpaidDBInstanceSpecResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeDatabases(req *DescribeDatabasesArgs) (resp *DescribeDatabasesResponse, err error) {
	resp = &DescribeDatabasesResponse{}
	err = c.Request("DescribeDatabases", req, resp)
	return
}

type DescribeDatabasesDatabase struct {
	DBName           string
	DBInstanceId     string
	Engine           string
	DBStatus         string
	CharacterSetName string
	DBDescription    string
	Accounts         DescribeDatabasesAccountPrivilegeInfoList
}

type DescribeDatabasesAccountPrivilegeInfo struct {
	Account          string
	AccountPrivilege string
}
type DescribeDatabasesArgs struct {
	ResourceOwnerId      int64
	DBName               string
	ResourceOwnerAccount string
	DBStatus             string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeDatabasesResponse struct {
	RequestId string
	Databases DescribeDatabasesDatabaseList
}

type DescribeDatabasesAccountPrivilegeInfoList []DescribeDatabasesAccountPrivilegeInfo

func (list *DescribeDatabasesAccountPrivilegeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDatabasesAccountPrivilegeInfo)
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

type DescribeDatabasesDatabaseList []DescribeDatabasesDatabase

func (list *DescribeDatabasesDatabaseList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDatabasesDatabase)
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

func (c *RdsClient) DescribeBackupTasks(req *DescribeBackupTasksArgs) (resp *DescribeBackupTasksResponse, err error) {
	resp = &DescribeBackupTasksResponse{}
	err = c.Request("DescribeBackupTasks", req, resp)
	return
}

type DescribeBackupTasksBackupJob struct {
	BackupProgressStatus string
	JobMode              string
	Process              string
	TaskAction           string
	BackupjobId          string
}
type DescribeBackupTasksArgs struct {
	BackupJobId          string
	ResourceOwnerId      int64
	Flag                 string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	BackupMode           string
	OwnerId              int64
	BackupJobStatus      string
}
type DescribeBackupTasksResponse struct {
	RequestId string
	Items     DescribeBackupTasksBackupJobList
}

type DescribeBackupTasksBackupJobList []DescribeBackupTasksBackupJob

func (list *DescribeBackupTasksBackupJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBackupTasksBackupJob)
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

func (c *RdsClient) DegradeDBInstanceSpec(req *DegradeDBInstanceSpecArgs) (resp *DegradeDBInstanceSpecResponse, err error) {
	resp = &DegradeDBInstanceSpecResponse{}
	err = c.Request("DegradeDBInstanceSpec", req, resp)
	return
}

type DegradeDBInstanceSpecArgs struct {
	ResourceOwnerId      int64
	DBInstanceStorage    int
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
	DBInstanceClass      string
}
type DegradeDBInstanceSpecResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeDBInstancesAsCsv(req *DescribeDBInstancesAsCsvArgs) (resp *DescribeDBInstancesAsCsvResponse, err error) {
	resp = &DescribeDBInstancesAsCsvResponse{}
	err = c.Request("DescribeDBInstancesAsCsv", req, resp)
	return
}

type DescribeDBInstancesAsCsvDBInstanceAttribute struct {
	InsId                       int
	DBInstanceId                string
	DBInstanceName              string
	PayType                     string
	DBInstanceClassType         string
	DBInstanceType              string
	RegionId                    string
	ConnectionString            string
	Port                        string
	Engine                      string
	EngineVersion               string
	DBInstanceClass             string
	DBInstanceMemory            int64
	DBInstanceStorage           int
	DBInstanceNetType           string
	DBInstanceStatus            string
	DBInstanceDescription       string
	LockMode                    string
	LockReason                  string
	ReadDelayTime               string
	DBMaxQuantity               int
	AccountMaxQuantity          int
	CreationTime                string
	ExpireTime                  string
	MaintainTime                string
	AvailabilityValue           string
	MaxIOPS                     int
	MaxConnections              int
	MasterInstanceId            string
	DBInstanceCPU               string
	IncrementSourceDBInstanceId string
	GuardDBInstanceId           string
	TempDBInstanceId            string
	SecurityIPList              string
	ZoneId                      string
	InstanceNetworkType         string
	Category                    string
	AccountType                 string
	SupportUpgradeAccountType   string
	VpcId                       string
	VSwitchId                   string
	ConnectionMode              string
	Tags                        string
}
type DescribeDBInstancesAsCsvArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeDBInstancesAsCsvResponse struct {
	RequestId string
	Items     DescribeDBInstancesAsCsvDBInstanceAttributeList
}

type DescribeDBInstancesAsCsvDBInstanceAttributeList []DescribeDBInstancesAsCsvDBInstanceAttribute

func (list *DescribeDBInstancesAsCsvDBInstanceAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesAsCsvDBInstanceAttribute)
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

func (c *RdsClient) DescribeImportsForSQLServer(req *DescribeImportsForSQLServerArgs) (resp *DescribeImportsForSQLServerResponse, err error) {
	resp = &DescribeImportsForSQLServerResponse{}
	err = c.Request("DescribeImportsForSQLServer", req, resp)
	return
}

type DescribeImportsForSQLServerSQLServerImport struct {
	ImportId     int
	FileName     string
	DBName       string
	ImportStatus string
	StartTime    string
}
type DescribeImportsForSQLServerArgs struct {
	ResourceOwnerId      int64
	ImportId             int
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	EndTime              string
	DBInstanceId         string
	StartTime            string
	OwnerId              int64
	PageNumber           int
}
type DescribeImportsForSQLServerResponse struct {
	RequestId         string
	TotalRecordCounts int
	PageNumber        int
	SQLItemsCounts    int
	Items             DescribeImportsForSQLServerSQLServerImportList
}

type DescribeImportsForSQLServerSQLServerImportList []DescribeImportsForSQLServerSQLServerImport

func (list *DescribeImportsForSQLServerSQLServerImportList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImportsForSQLServerSQLServerImport)
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

func (c *RdsClient) DescribeOptimizeAdviceOnMissPK(req *DescribeOptimizeAdviceOnMissPKArgs) (resp *DescribeOptimizeAdviceOnMissPKResponse, err error) {
	resp = &DescribeOptimizeAdviceOnMissPKResponse{}
	err = c.Request("DescribeOptimizeAdviceOnMissPK", req, resp)
	return
}

type DescribeOptimizeAdviceOnMissPKAdviceOnMissPK struct {
	DBName    string
	TableName string
}
type DescribeOptimizeAdviceOnMissPKArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	DBInstanceId         string
	OwnerId              int64
	PageNumber           int
}
type DescribeOptimizeAdviceOnMissPKResponse struct {
	RequestId         string
	TotalRecordsCount int
	PageNumber        int
	PageRecordCount   int
	Items             DescribeOptimizeAdviceOnMissPKAdviceOnMissPKList
}

type DescribeOptimizeAdviceOnMissPKAdviceOnMissPKList []DescribeOptimizeAdviceOnMissPKAdviceOnMissPK

func (list *DescribeOptimizeAdviceOnMissPKAdviceOnMissPKList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOptimizeAdviceOnMissPKAdviceOnMissPK)
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

func (c *RdsClient) CreateDatabase(req *CreateDatabaseArgs) (resp *CreateDatabaseResponse, err error) {
	resp = &CreateDatabaseResponse{}
	err = c.Request("CreateDatabase", req, resp)
	return
}

type CreateDatabaseArgs struct {
	ResourceOwnerId      int64
	DBName               string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	DBDescription        string
	OwnerId              int64
	CharacterSetName     string
}
type CreateDatabaseResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeDBInstanceNetworkDetail(req *DescribeDBInstanceNetworkDetailArgs) (resp *DescribeDBInstanceNetworkDetailResponse, err error) {
	resp = &DescribeDBInstanceNetworkDetailResponse{}
	err = c.Request("DescribeDBInstanceNetworkDetail", req, resp)
	return
}

type DescribeDBInstanceNetworkDetailNetworkKey struct {
	Key    string
	Unit   string
	Values DescribeDBInstanceNetworkDetailNetworkValueList
}

type DescribeDBInstanceNetworkDetailNetworkValue struct {
	Value    string
	DateTime string
}
type DescribeDBInstanceNetworkDetailArgs struct {
	EndPoint     string
	StartPoint   string
	EndTime      string
	DBInstanceId string
	StartTime    string
}
type DescribeDBInstanceNetworkDetailResponse struct {
	RequestId         string
	DBInstanceId      string
	StartTime         string
	EndTime           string
	NewConnection     string
	ActiveConnection  string
	AbortedConnection string
	NetworkRequest    string
	NetworkTrafficIn  string
	NetworkTrafficOut string
	NetworkLatency    string
	BackendLatency    string
	NetworkErrors     string
	NetworkDetail     DescribeDBInstanceNetworkDetailNetworkKeyList
}

type DescribeDBInstanceNetworkDetailNetworkValueList []DescribeDBInstanceNetworkDetailNetworkValue

func (list *DescribeDBInstanceNetworkDetailNetworkValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetworkDetailNetworkValue)
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

type DescribeDBInstanceNetworkDetailNetworkKeyList []DescribeDBInstanceNetworkDetailNetworkKey

func (list *DescribeDBInstanceNetworkDetailNetworkKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetworkDetailNetworkKey)
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

func (c *RdsClient) CheckResource(req *CheckResourceArgs) (resp *CheckResourceResponse, err error) {
	resp = &CheckResourceResponse{}
	err = c.Request("CheckResource", req, resp)
	return
}

type CheckResourceResource struct {
	DBInstanceAvailable string
	Engine              string
	EngineVersion       string
}
type CheckResourceArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	SpecifyCount         string
	EngineVersion        string
	OwnerId              int64
	DBInstanceClass      string
	Engine               string
	ZoneId               string
	DBInstanceUseType    string
	DBInstanceId         string
}
type CheckResourceResponse struct {
	RequestId    string
	SpecifyCount string
	Resources    CheckResourceResourceList
}

type CheckResourceResourceList []CheckResourceResource

func (list *CheckResourceResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CheckResourceResource)
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

func (c *RdsClient) DescribeCharacterSetName(req *DescribeCharacterSetNameArgs) (resp *DescribeCharacterSetNameResponse, err error) {
	resp = &DescribeCharacterSetNameResponse{}
	err = c.Request("DescribeCharacterSetName", req, resp)
	return
}

type DescribeCharacterSetNameArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	Engine               string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeCharacterSetNameResponse struct {
	RequestId             string
	Engine                string
	CharacterSetNameItems DescribeCharacterSetNameCharacterSetNameItemList
}

type DescribeCharacterSetNameCharacterSetNameItemList []string

func (list *DescribeCharacterSetNameCharacterSetNameItemList) UnmarshalJSON(data []byte) error {
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

func (c *RdsClient) CreateDBInstanceReplica(req *CreateDBInstanceReplicaArgs) (resp *CreateDBInstanceReplicaResponse, err error) {
	resp = &CreateDBInstanceReplicaResponse{}
	err = c.Request("CreateDBInstanceReplica", req, resp)
	return
}

type CreateDBInstanceReplicaArgs struct {
	ConnectionMode        string
	ReplicaDescription    string
	ResourceOwnerId       int64
	DBInstanceStorage     int
	SystemDBCharset       string
	ClientToken           string
	EngineVersion         string
	Engine                string
	DBInstanceDescription string
	DBInstanceNetType     string
	Period                string
	ResourceOwnerAccount  string
	OwnerAccount          string
	OwnerId               int64
	UsedTime              string
	DBInstanceClass       string
	SecurityIPList        string
	VSwitchId             string
	PrivateIpAddress      string
	SourceDBInstanceId    string
	VPCId                 string
	ZoneId                string
	PayType               string
	InstanceNetworkType   string
}
type CreateDBInstanceReplicaResponse struct {
	RequestId    string
	DBInstanceId string
	OrderId      int64
	ReplicaId    string
	WorkflowId   string
}

func (c *RdsClient) MigrateToOtherZone(req *MigrateToOtherZoneArgs) (resp *MigrateToOtherZoneResponse, err error) {
	resp = &MigrateToOtherZoneResponse{}
	err = c.Request("MigrateToOtherZone", req, resp)
	return
}

type MigrateToOtherZoneArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	ZoneId               string
	DBInstanceId         string
	OwnerId              int64
}
type MigrateToOtherZoneResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeSQLReports(req *DescribeSQLReportsArgs) (resp *DescribeSQLReportsResponse, err error) {
	resp = &DescribeSQLReportsResponse{}
	err = c.Request("DescribeSQLReports", req, resp)
	return
}

type DescribeSQLReportsItem struct {
	ReportTime       string
	LatencyTopNItems DescribeSQLReportsLatencyTopNItemList
	QPSTopNItems     DescribeSQLReportsQPSTopNItemList
}

type DescribeSQLReportsLatencyTopNItem struct {
	SQLText         string
	AvgLatency      int64
	SQLExecuteTimes int64
}

type DescribeSQLReportsQPSTopNItem struct {
	SQLText         string
	SQLExecuteTimes int64
}
type DescribeSQLReportsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	EndTime              string
	DBInstanceId         string
	StartTime            string
	OwnerId              int64
	PageNumber           int
}
type DescribeSQLReportsResponse struct {
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeSQLReportsItemList
}

type DescribeSQLReportsLatencyTopNItemList []DescribeSQLReportsLatencyTopNItem

func (list *DescribeSQLReportsLatencyTopNItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLReportsLatencyTopNItem)
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

type DescribeSQLReportsQPSTopNItemList []DescribeSQLReportsQPSTopNItem

func (list *DescribeSQLReportsQPSTopNItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLReportsQPSTopNItem)
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

type DescribeSQLReportsItemList []DescribeSQLReportsItem

func (list *DescribeSQLReportsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLReportsItem)
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

func (c *RdsClient) ModifyBackupPolicy(req *ModifyBackupPolicyArgs) (resp *ModifyBackupPolicyResponse, err error) {
	resp = &ModifyBackupPolicyResponse{}
	err = c.Request("ModifyBackupPolicy", req, resp)
	return
}

type ModifyBackupPolicyArgs struct {
	PreferredBackupTime      string
	PreferredBackupPeriod    string
	BackupRetentionPeriod    string
	ResourceOwnerId          int64
	ResourceOwnerAccount     string
	OwnerAccount             string
	DBInstanceId             string
	BackupLog                string
	OwnerId                  int64
	LogBackupRetentionPeriod string
}
type ModifyBackupPolicyResponse struct {
	RequestId string
}

func (c *RdsClient) StopSyncing(req *StopSyncingArgs) (resp *StopSyncingResponse, err error) {
	resp = &StopSyncingResponse{}
	err = c.Request("StopSyncing", req, resp)
	return
}

type StopSyncingArgs struct {
	ResourceOwnerId      int64
	ImportId             int
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type StopSyncingResponse struct {
	RequestId string
}

func (c *RdsClient) AllocateInstancePrivateConnection(req *AllocateInstancePrivateConnectionArgs) (resp *AllocateInstancePrivateConnectionResponse, err error) {
	resp = &AllocateInstancePrivateConnectionResponse{}
	err = c.Request("AllocateInstancePrivateConnection", req, resp)
	return
}

type AllocateInstancePrivateConnectionArgs struct {
	ResourceOwnerId        int64
	ConnectionStringPrefix string
	ResourceOwnerAccount   string
	Port                   string
	OwnerAccount           string
	DBInstanceId           string
	OwnerId                int64
}
type AllocateInstancePrivateConnectionResponse struct {
	RequestId string
}

func (c *RdsClient) ImportDataForSQLServer(req *ImportDataForSQLServerArgs) (resp *ImportDataForSQLServerResponse, err error) {
	resp = &ImportDataForSQLServerResponse{}
	err = c.Request("ImportDataForSQLServer", req, resp)
	return
}

type ImportDataForSQLServerArgs struct {
	ResourceOwnerId      int64
	FileName             string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type ImportDataForSQLServerResponse struct {
	RequestId string
	ImportID  int
}

func (c *RdsClient) DescribeSQLLogRecords(req *DescribeSQLLogRecordsArgs) (resp *DescribeSQLLogRecordsResponse, err error) {
	resp = &DescribeSQLLogRecordsResponse{}
	err = c.Request("DescribeSQLLogRecords", req, resp)
	return
}

type DescribeSQLLogRecordsSQLRecord struct {
	DBName              string
	AccountName         string
	HostAddress         string
	SQLText             string
	TotalExecutionTimes int64
	ReturnRowCounts     int64
	ExecuteTime         string
	ThreadID            string
}
type DescribeSQLLogRecordsArgs struct {
	SQLId                int64
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	EndTime              string
	StartTime            string
	OwnerId              int64
	QueryKeywords        string
	PageNumber           int
	Database             string
	Form                 string
	PageSize             int
	DBInstanceId         string
	User                 string
}
type DescribeSQLLogRecordsResponse struct {
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeSQLLogRecordsSQLRecordList
}

type DescribeSQLLogRecordsSQLRecordList []DescribeSQLLogRecordsSQLRecord

func (list *DescribeSQLLogRecordsSQLRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogRecordsSQLRecord)
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

func (c *RdsClient) RevokeAccountPrivilege(req *RevokeAccountPrivilegeArgs) (resp *RevokeAccountPrivilegeResponse, err error) {
	resp = &RevokeAccountPrivilegeResponse{}
	err = c.Request("RevokeAccountPrivilege", req, resp)
	return
}

type RevokeAccountPrivilegeArgs struct {
	ResourceOwnerId      int64
	AccountName          string
	DBName               string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type RevokeAccountPrivilegeResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeDBInstanceByTags(req *DescribeDBInstanceByTagsArgs) (resp *DescribeDBInstanceByTagsResponse, err error) {
	resp = &DescribeDBInstanceByTagsResponse{}
	err = c.Request("DescribeDBInstanceByTags", req, resp)
	return
}

type DescribeDBInstanceByTagsDBInstanceTag struct {
	DBInstanceId string
	Tags         DescribeDBInstanceByTagsTagList
}

type DescribeDBInstanceByTagsTag struct {
	TagKey   string
	TagValue string
}
type DescribeDBInstanceByTagsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	PageSize             int
	DBInstanceId         string
	OwnerId              int64
	PageNumber           int
	ProxyId              string
}
type DescribeDBInstanceByTagsResponse struct {
	RequestId        string
	PageNumber       int
	PageRecordCount  int
	TotalRecordCount int
	Items            DescribeDBInstanceByTagsDBInstanceTagList
}

type DescribeDBInstanceByTagsTagList []DescribeDBInstanceByTagsTag

func (list *DescribeDBInstanceByTagsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceByTagsTag)
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

type DescribeDBInstanceByTagsDBInstanceTagList []DescribeDBInstanceByTagsDBInstanceTag

func (list *DescribeDBInstanceByTagsDBInstanceTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceByTagsDBInstanceTag)
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

func (c *RdsClient) DescribeDBInstancesByPerformance(req *DescribeDBInstancesByPerformanceArgs) (resp *DescribeDBInstancesByPerformanceResponse, err error) {
	resp = &DescribeDBInstancesByPerformanceResponse{}
	err = c.Request("DescribeDBInstancesByPerformance", req, resp)
	return
}

type DescribeDBInstancesByPerformanceDBInstancePerformance struct {
	CPUUsage              string
	IOPSUsage             string
	DiskUsage             string
	SessionUsage          string
	DBInstanceId          string
	DBInstanceDescription string
}
type DescribeDBInstancesByPerformanceArgs struct {
	Tag4value            string
	ResourceOwnerId      int64
	Tag2key              string
	Tag5key              string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	Tag3key              string
	OwnerId              int64
	Tag5value            string
	PageNumber           int
	Tags                 string
	Tag1key              string
	Tag1value            string
	SortKey              string
	SortMethod           string
	Tag2value            string
	PageSize             int
	Tag4key              string
	DBInstanceId         string
	Tag3value            string
	ProxyId              string
}
type DescribeDBInstancesByPerformanceResponse struct {
	RequestId        string
	PageNumber       int
	TotalRecordCount int
	PageRecordCount  int
	Items            DescribeDBInstancesByPerformanceDBInstancePerformanceList
}

type DescribeDBInstancesByPerformanceDBInstancePerformanceList []DescribeDBInstancesByPerformanceDBInstancePerformance

func (list *DescribeDBInstancesByPerformanceDBInstancePerformanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesByPerformanceDBInstancePerformance)
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

func (c *RdsClient) DescribeAccounts(req *DescribeAccountsArgs) (resp *DescribeAccountsResponse, err error) {
	resp = &DescribeAccountsResponse{}
	err = c.Request("DescribeAccounts", req, resp)
	return
}

type DescribeAccountsDBInstanceAccount struct {
	DBInstanceId       string
	AccountName        string
	AccountStatus      string
	AccountType        string
	AccountDescription string
	DatabasePrivileges DescribeAccountsDatabasePrivilegeList
}

type DescribeAccountsDatabasePrivilege struct {
	DBName           string
	AccountPrivilege string
}
type DescribeAccountsArgs struct {
	ResourceOwnerId      int64
	AccountName          string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeAccountsResponse struct {
	RequestId string
	Accounts  DescribeAccountsDBInstanceAccountList
}

type DescribeAccountsDatabasePrivilegeList []DescribeAccountsDatabasePrivilege

func (list *DescribeAccountsDatabasePrivilegeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccountsDatabasePrivilege)
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

func (c *RdsClient) CreateUploadPathForSQLServer(req *CreateUploadPathForSQLServerArgs) (resp *CreateUploadPathForSQLServerResponse, err error) {
	resp = &CreateUploadPathForSQLServerResponse{}
	err = c.Request("CreateUploadPathForSQLServer", req, resp)
	return
}

type CreateUploadPathForSQLServerArgs struct {
	ResourceOwnerId      int64
	DBName               string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type CreateUploadPathForSQLServerResponse struct {
	RequestId         string
	InternetFtpServer string
	InternetPort      int
	IntranetFtpserver string
	Intranetport      int
	UserName          string
	Password          string
	FileName          string
}

func (c *RdsClient) CreateTempDBInstance(req *CreateTempDBInstanceArgs) (resp *CreateTempDBInstanceResponse, err error) {
	resp = &CreateTempDBInstanceResponse{}
	err = c.Request("CreateTempDBInstance", req, resp)
	return
}

type CreateTempDBInstanceArgs struct {
	ResourceOwnerId      int64
	RestoreTime          string
	ResourceOwnerAccount string
	BackupId             int
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type CreateTempDBInstanceResponse struct {
	RequestId        string
	TempDBInstanceId string
}

func (c *RdsClient) DescribeOptimizeAdviceOnExcessIndex(req *DescribeOptimizeAdviceOnExcessIndexArgs) (resp *DescribeOptimizeAdviceOnExcessIndexResponse, err error) {
	resp = &DescribeOptimizeAdviceOnExcessIndexResponse{}
	err = c.Request("DescribeOptimizeAdviceOnExcessIndex", req, resp)
	return
}

type DescribeOptimizeAdviceOnExcessIndexAdviceOnExcessIndex struct {
	DBName     string
	TableName  string
	IndexCount int64
}
type DescribeOptimizeAdviceOnExcessIndexArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	DBInstanceId         string
	OwnerId              int64
	PageNumber           int
}
type DescribeOptimizeAdviceOnExcessIndexResponse struct {
	RequestId         string
	TotalRecordsCount int
	PageNumber        int
	PageRecordCount   int
	Items             DescribeOptimizeAdviceOnExcessIndexAdviceOnExcessIndexList
}

type DescribeOptimizeAdviceOnExcessIndexAdviceOnExcessIndexList []DescribeOptimizeAdviceOnExcessIndexAdviceOnExcessIndex

func (list *DescribeOptimizeAdviceOnExcessIndexAdviceOnExcessIndexList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOptimizeAdviceOnExcessIndexAdviceOnExcessIndex)
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

func (c *RdsClient) ModifyDBInstanceTDE(req *ModifyDBInstanceTDEArgs) (resp *ModifyDBInstanceTDEResponse, err error) {
	resp = &ModifyDBInstanceTDEResponse{}
	err = c.Request("ModifyDBInstanceTDE", req, resp)
	return
}

type ModifyDBInstanceTDEArgs struct {
	ResourceOwnerId      int64
	DBName               string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
	TDEStatus            string
}
type ModifyDBInstanceTDEResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeReplicaPerformance(req *DescribeReplicaPerformanceArgs) (resp *DescribeReplicaPerformanceResponse, err error) {
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

func (c *RdsClient) DescribeOptimizeAdviceOnMissIndex(req *DescribeOptimizeAdviceOnMissIndexArgs) (resp *DescribeOptimizeAdviceOnMissIndexResponse, err error) {
	resp = &DescribeOptimizeAdviceOnMissIndexResponse{}
	err = c.Request("DescribeOptimizeAdviceOnMissIndex", req, resp)
	return
}

type DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndex struct {
	DBName      string
	TableName   string
	QueryColumn string
	SQLText     string
}
type DescribeOptimizeAdviceOnMissIndexArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	DBInstanceId         string
	OwnerId              int64
	PageNumber           int
}
type DescribeOptimizeAdviceOnMissIndexResponse struct {
	RequestId         string
	DBInstanceId      string
	TotalRecordsCount int
	PageNumber        int
	PageRecordCount   int
	Items             DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndexList
}

type DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndexList []DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndex

func (list *DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndexList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeOptimizeAdviceOnMissIndexAdviceOnMissIndex)
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

func (c *RdsClient) DeleteDampPolicy(req *DeleteDampPolicyArgs) (resp *DeleteDampPolicyResponse, err error) {
	resp = &DeleteDampPolicyResponse{}
	err = c.Request("DeleteDampPolicy", req, resp)
	return
}

type DeleteDampPolicyArgs struct {
	ResourceOwnerId      int64
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	PolicyName           string
	OwnerId              int64
}
type DeleteDampPolicyResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeRegions(req *DescribeRegionsArgs) (resp *DescribeRegionsResponse, err error) {
	resp = &DescribeRegionsResponse{}
	err = c.Request("DescribeRegions", req, resp)
	return
}

type DescribeRegionsRDSRegion struct {
	RegionId string
	ZoneId   string
}
type DescribeRegionsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeRegionsResponse struct {
	RequestId string
	Regions   DescribeRegionsRDSRegionList
}

type DescribeRegionsRDSRegionList []DescribeRegionsRDSRegion

func (list *DescribeRegionsRDSRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsRDSRegion)
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

func (c *RdsClient) DescribeParameters(req *DescribeParametersArgs) (resp *DescribeParametersResponse, err error) {
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
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DescribeParametersResponse struct {
	RequestId         string
	Engine            string
	EngineVersion     string
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

func (c *RdsClient) DescribeDBInstancePerformance(req *DescribeDBInstancePerformanceArgs) (resp *DescribeDBInstancePerformanceResponse, err error) {
	resp = &DescribeDBInstancePerformanceResponse{}
	err = c.Request("DescribeDBInstancePerformance", req, resp)
	return
}

type DescribeDBInstancePerformancePerformanceKey struct {
	Key         string
	Unit        string
	ValueFormat string
	Values      DescribeDBInstancePerformancePerformanceValueList
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
	DBInstanceId         string
	StartTime            string
	OwnerId              int64
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

func (c *RdsClient) RevokeOperatorPermission(req *RevokeOperatorPermissionArgs) (resp *RevokeOperatorPermissionResponse, err error) {
	resp = &RevokeOperatorPermissionResponse{}
	err = c.Request("RevokeOperatorPermission", req, resp)
	return
}

type RevokeOperatorPermissionArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type RevokeOperatorPermissionResponse struct {
	RequestId string
}

func (c *RdsClient) DescibeImportsFromDatabase(req *DescibeImportsFromDatabaseArgs) (resp *DescibeImportsFromDatabaseResponse, err error) {
	resp = &DescibeImportsFromDatabaseResponse{}
	err = c.Request("DescibeImportsFromDatabase", req, resp)
	return
}

type DescibeImportsFromDatabaseImportResultFromDB struct {
	ImportId                    int
	ImportDataType              string
	ImportDataStatus            string
	ImportDataStatusDescription string
	IncrementalImportingTime    string
}
type DescibeImportsFromDatabaseArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	EndTime              string
	StartTime            string
	OwnerId              int64
	PageNumber           int
	ImportId             int
	Engine               string
	PageSize             int
	DBInstanceId         string
}
type DescibeImportsFromDatabaseResponse struct {
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescibeImportsFromDatabaseImportResultFromDBList
}

type DescibeImportsFromDatabaseImportResultFromDBList []DescibeImportsFromDatabaseImportResultFromDB

func (list *DescibeImportsFromDatabaseImportResultFromDBList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescibeImportsFromDatabaseImportResultFromDB)
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

func (c *RdsClient) DescribeInstanceAutoRenewalAttribute(req *DescribeInstanceAutoRenewalAttributeArgs) (resp *DescribeInstanceAutoRenewalAttributeResponse, err error) {
	resp = &DescribeInstanceAutoRenewalAttributeResponse{}
	err = c.Request("DescribeInstanceAutoRenewalAttribute", req, resp)
	return
}

type DescribeInstanceAutoRenewalAttributeItem struct {
	DBInstanceId string
	RegionId     string
	Duration     int
	Status       string
	AutoRenew    string
}
type DescribeInstanceAutoRenewalAttributeArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	PageSize             int
	DBInstanceId         string
	OwnerId              int64
	PageNumber           int
	ProxyId              string
}
type DescribeInstanceAutoRenewalAttributeResponse struct {
	RequestId        string
	PageNumber       int
	TotalRecordCount int
	PageRecordCount  int
	Items            DescribeInstanceAutoRenewalAttributeItemList
}

type DescribeInstanceAutoRenewalAttributeItemList []DescribeInstanceAutoRenewalAttributeItem

func (list *DescribeInstanceAutoRenewalAttributeItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceAutoRenewalAttributeItem)
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

func (c *RdsClient) DeleteDBInstance(req *DeleteDBInstanceArgs) (resp *DeleteDBInstanceResponse, err error) {
	resp = &DeleteDBInstanceResponse{}
	err = c.Request("DeleteDBInstance", req, resp)
	return
}

type DeleteDBInstanceArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type DeleteDBInstanceResponse struct {
	RequestId string
}

func (c *RdsClient) ModifyResourceGroup(req *ModifyResourceGroupArgs) (resp *ModifyResourceGroupResponse, err error) {
	resp = &ModifyResourceGroupResponse{}
	err = c.Request("ModifyResourceGroup", req, resp)
	return
}

type ModifyResourceGroupArgs struct {
	ResourceGroupId      string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type ModifyResourceGroupResponse struct {
	RequestId string
}

func (c *RdsClient) RequestServiceOfCloudDBA(req *RequestServiceOfCloudDBAArgs) (resp *RequestServiceOfCloudDBAResponse, err error) {
	resp = &RequestServiceOfCloudDBAResponse{}
	err = c.Request("RequestServiceOfCloudDBA", req, resp)
	return
}

type RequestServiceOfCloudDBAArgs struct {
	ServiceRequestParam string
	DBInstanceId        string
	ServiceRequestType  string
}
type RequestServiceOfCloudDBAResponse struct {
	RequestId string
	ListData  string
	AttrData  string
}

func (c *RdsClient) DescribeBinlogFiles(req *DescribeBinlogFilesArgs) (resp *DescribeBinlogFilesResponse, err error) {
	resp = &DescribeBinlogFilesResponse{}
	err = c.Request("DescribeBinlogFiles", req, resp)
	return
}

type DescribeBinlogFilesBinLogFile struct {
	FileSize             int64
	LogBeginTime         string
	LogEndTime           string
	DownloadLink         string
	IntranetDownloadLink string
	LinkExpiredTime      string
	Checksum             string
	HostInstanceID       string
}
type DescribeBinlogFilesArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	EndTime              string
	DBInstanceId         string
	StartTime            string
	OwnerId              int64
	PageNumber           int
}
type DescribeBinlogFilesResponse struct {
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	TotalFileSize    int64
	Items            DescribeBinlogFilesBinLogFileList
}

type DescribeBinlogFilesBinLogFileList []DescribeBinlogFilesBinLogFile

func (list *DescribeBinlogFilesBinLogFileList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBinlogFilesBinLogFile)
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

func (c *RdsClient) DescribePrice(req *DescribePriceArgs) (resp *DescribePriceResponse, err error) {
	resp = &DescribePriceResponse{}
	err = c.Request("DescribePrice", req, resp)
	return
}

type DescribePriceRule struct {
	RuleId      int64
	Name        string
	Description string
}

type DescribePricePriceInfo struct {
	Currency      string
	OriginalPrice float32
	TradePrice    float32
	DiscountPrice float32
	Coupons       DescribePriceCouponList
	RuleIds       DescribePriceRuleIdList
	ActivityInfo  DescribePriceActivityInfo
}

type DescribePriceCoupon struct {
	CouponNo    string
	Name        string
	Description string
	IsSelected  string
}

type DescribePriceActivityInfo struct {
	CheckErrMsg string
	ErrorCode   string
	Success     string
}
type DescribePriceArgs struct {
	ResourceOwnerId      int64
	DBInstanceStorage    int
	Quantity             int
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	CommodityCode        string
	EngineVersion        string
	OwnerId              int64
	UsedTime             string
	DBInstanceClass      string
	InstanceUsedType     int
	Engine               string
	ZoneId               string
	TimeType             string
	PayType              string
	OrderType            string
}
type DescribePriceResponse struct {
	RequestId string
	Rules     DescribePriceRuleList
	PriceInfo DescribePricePriceInfo
}

type DescribePriceCouponList []DescribePriceCoupon

func (list *DescribePriceCouponList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePriceCoupon)
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

type DescribePriceRuleIdList []string

func (list *DescribePriceRuleIdList) UnmarshalJSON(data []byte) error {
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

type DescribePriceRuleList []DescribePriceRule

func (list *DescribePriceRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePriceRule)
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

func (c *RdsClient) ModifyDBInstanceNetworkExpireTime(req *ModifyDBInstanceNetworkExpireTimeArgs) (resp *ModifyDBInstanceNetworkExpireTimeResponse, err error) {
	resp = &ModifyDBInstanceNetworkExpireTimeResponse{}
	err = c.Request("ModifyDBInstanceNetworkExpireTime", req, resp)
	return
}

type ModifyDBInstanceNetworkExpireTimeArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ConnectionString     string
	ClassicExpiredDays   int
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type ModifyDBInstanceNetworkExpireTimeResponse struct {
	RequestId string
}

func (c *RdsClient) DescribeDatabaseLockDiagnosis(req *DescribeDatabaseLockDiagnosisArgs) (resp *DescribeDatabaseLockDiagnosisResponse, err error) {
	resp = &DescribeDatabaseLockDiagnosisResponse{}
	err = c.Request("DescribeDatabaseLockDiagnosis", req, resp)
	return
}

type DescribeDatabaseLockDiagnosisArgs struct {
	EndTime      string
	DBInstanceId string
	StartTime    string
}
type DescribeDatabaseLockDiagnosisResponse struct {
	RequestId    string
	DeadLockList DescribeDatabaseLockDiagnosisDeadLockListList
}

type DescribeDatabaseLockDiagnosisDeadLockListList []string

func (list *DescribeDatabaseLockDiagnosisDeadLockListList) UnmarshalJSON(data []byte) error {
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

func (c *RdsClient) DescribeDBInstanceAttribute(req *DescribeDBInstanceAttributeArgs) (resp *DescribeDBInstanceAttributeResponse, err error) {
	resp = &DescribeDBInstanceAttributeResponse{}
	err = c.Request("DescribeDBInstanceAttribute", req, resp)
	return
}

type DescribeDBInstanceAttributeDBInstanceAttribute struct {
	InsId                       int
	DBInstanceId                string
	PayType                     string
	DBInstanceClassType         string
	DBInstanceType              string
	RegionId                    string
	ConnectionString            string
	Port                        string
	Engine                      string
	EngineVersion               string
	DBInstanceClass             string
	DBInstanceMemory            int64
	DBInstanceStorage           int
	DBInstanceNetType           string
	DBInstanceStatus            string
	DBInstanceDescription       string
	LockMode                    string
	LockReason                  string
	ReadDelayTime               string
	DBMaxQuantity               int
	AccountMaxQuantity          int
	CreationTime                string
	ExpireTime                  string
	MaintainTime                string
	AvailabilityValue           string
	MaxIOPS                     int
	MaxConnections              int
	MasterInstanceId            string
	DBInstanceCPU               string
	IncrementSourceDBInstanceId string
	GuardDBInstanceId           string
	TempDBInstanceId            string
	SecurityIPList              string
	ZoneId                      string
	InstanceNetworkType         string
	Category                    string
	AccountType                 string
	SupportUpgradeAccountType   string
	VpcId                       string
	VSwitchId                   string
	ConnectionMode              string
	ResourceGroupId             string
	ReadOnlyDBInstanceIds       DescribeDBInstanceAttributeReadOnlyDBInstanceIdList
}

type DescribeDBInstanceAttributeReadOnlyDBInstanceId struct {
	DBInstanceId string
}
type DescribeDBInstanceAttributeArgs struct {
	ResourceGroupId      string
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

type DescribeDBInstanceAttributeReadOnlyDBInstanceIdList []DescribeDBInstanceAttributeReadOnlyDBInstanceId

func (list *DescribeDBInstanceAttributeReadOnlyDBInstanceIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceAttributeReadOnlyDBInstanceId)
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

func (c *RdsClient) CreatePolicyWithSpecifiedPolicy(req *CreatePolicyWithSpecifiedPolicyArgs) (resp *CreatePolicyWithSpecifiedPolicyResponse, err error) {
	resp = &CreatePolicyWithSpecifiedPolicyResponse{}
	err = c.Request("CreatePolicyWithSpecifiedPolicy", req, resp)
	return
}

type CreatePolicyWithSpecifiedPolicyArgs struct {
	ResourceOwnerId      int64
	PolicyId             string
	SecurityToken        string
	ResourceOwnerAccount string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type CreatePolicyWithSpecifiedPolicyResponse struct {
	RequestId string
}

func (c *RdsClient) CreateReadOnlyDBInstance(req *CreateReadOnlyDBInstanceArgs) (resp *CreateReadOnlyDBInstanceResponse, err error) {
	resp = &CreateReadOnlyDBInstanceResponse{}
	err = c.Request("CreateReadOnlyDBInstance", req, resp)
	return
}

type CreateReadOnlyDBInstanceArgs struct {
	ResourceOwnerId       int64
	DBInstanceStorage     int
	ResourceOwnerAccount  string
	ClientToken           string
	OwnerAccount          string
	EngineVersion         string
	OwnerId               int64
	DBInstanceClass       string
	VSwitchId             string
	PrivateIpAddress      string
	ResourceGroupId       string
	VPCId                 string
	ZoneId                string
	DBInstanceId          string
	DBInstanceDescription string
	PayType               string
	InstanceNetworkType   string
}
type CreateReadOnlyDBInstanceResponse struct {
	RequestId        string
	DBInstanceId     string
	OrderId          string
	ConnectionString string
	Port             string
}

func (c *RdsClient) PurgeDBInstanceLog(req *PurgeDBInstanceLogArgs) (resp *PurgeDBInstanceLogResponse, err error) {
	resp = &PurgeDBInstanceLogResponse{}
	err = c.Request("PurgeDBInstanceLog", req, resp)
	return
}

type PurgeDBInstanceLogArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DBInstanceId         string
	OwnerId              int64
}
type PurgeDBInstanceLogResponse struct {
	RequestId string
}

func (c *RdsClient) GrantOperatorPermission(req *GrantOperatorPermissionArgs) (resp *GrantOperatorPermissionResponse, err error) {
	resp = &GrantOperatorPermissionResponse{}
	err = c.Request("GrantOperatorPermission", req, resp)
	return
}

type GrantOperatorPermissionArgs struct {
	Privileges           string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	ExpiredTime          string
	DBInstanceId         string
	OwnerId              int64
}
type GrantOperatorPermissionResponse struct {
	RequestId string
}
