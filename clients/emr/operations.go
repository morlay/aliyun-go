package emr

import (
	"encoding/json"

	"github.com/morlay/aliyun-go/core"
)

func (c *EmrClient) SuspendExecutionPlanScheduler(req *SuspendExecutionPlanSchedulerArgs) (resp *SuspendExecutionPlanSchedulerResponse, err error) {
	resp = &SuspendExecutionPlanSchedulerResponse{}
	err = c.Request("SuspendExecutionPlanScheduler", req, resp)
	return
}

type SuspendExecutionPlanSchedulerArgs struct {
	ResourceOwnerId int64
	Id              string
}
type SuspendExecutionPlanSchedulerResponse struct {
	RequestId string
}

func (c *EmrClient) RenewCluster(req *RenewClusterArgs) (resp *RenewClusterResponse, err error) {
	resp = &RenewClusterResponse{}
	err = c.Request("RenewCluster", req, resp)
	return
}

type RenewClusterArgs struct {
	ResourceOwnerId int64
	Id              string
	ECSIds          core.Bool
	Period          string
}
type RenewClusterResponse struct {
	RequestId string
	ClusterId string
}

func (c *EmrClient) CreateJob(req *CreateJobArgs) (resp *CreateJobResponse, err error) {
	resp = &CreateJobResponse{}
	err = c.Request("CreateJob", req, resp)
	return
}

type CreateJobArgs struct {
	ResourceOwnerId int64
	Name            string
	Type            string
	RunParameter    string
	FailAct         string
}
type CreateJobResponse struct {
	RequestId string
	Id        string
}

func (c *EmrClient) ListAvailableConfig(req *ListAvailableConfigArgs) (resp *ListAvailableConfigResponse, err error) {
	resp = &ListAvailableConfigResponse{}
	err = c.Request("ListAvailableConfig", req, resp)
	return
}

type ListAvailableConfigSecurityGroupType struct {
	Name  string
	State string
	Id    string
}

type ListAvailableConfigInstanceType struct {
	Classify               string
	Type                   string
	CpuCore                int
	MemSize                int
	HasCloudDisk           core.Bool
	HasEfficiencyCloudDisk core.Bool
	HasSSDCloudDisk        core.Bool
}

type ListAvailableConfigEmrVerType struct {
	Name       string
	SubModules ListAvailableConfigSubModuleList
}

type ListAvailableConfigSubModule struct {
	Name         string
	RequiredList ListAvailableConfigRequiredList
	OptionalList ListAvailableConfigOptionalList
}

type ListAvailableConfigRequired struct {
	DisplayName string
	Name        string
	OnlyDisplay core.Bool
	StartTpe    int
	Version     string
}

type ListAvailableConfigOptional struct {
	DisplayName string
	Name        string
	OnlyDisplay core.Bool
	StartTpe    int
	Version     string
}

type ListAvailableConfigZoneType struct {
	Name                          string
	Id                            string
	AvailableResources            ListAvailableConfigAvailableResourceList
	AvailableResourceCreationList ListAvailableConfigAvailableResourceCreationListList
	AvailableDiskCategoryList     ListAvailableConfigAvailableDiskCategoryListList
}

type ListAvailableConfigAvailableResource struct {
	IoOptimized            core.Bool
	SystemDiskCategories   ListAvailableConfigSystemDiskCategoryList
	DataDiskCategories     ListAvailableConfigDataDiskCategoryList
	NetworkTypes           ListAvailableConfigNetworkTypeList
	SupportedInstanceTypes ListAvailableConfigSupportedInstanceTypeList
	InstanceTypeFamilies   ListAvailableConfigInstanceTypeFamilyList
	InstanceGenerations    ListAvailableConfigInstanceGenerationList
}

type ListAvailableConfigVpc struct {
	Id             string
	CidrBlock      string
	VSwitchs       ListAvailableConfigVSwitchList
	SecurityGroups ListAvailableConfigSecurityGroupList
}

type ListAvailableConfigVSwitch struct {
	Id        string
	CidrBlock string
	ZoneId    string
}

type ListAvailableConfigSecurityGroup struct {
	Name string
	Id   string
}
type ListAvailableConfigArgs struct {
	ResourceOwnerId int64
}
type ListAvailableConfigResponse struct {
	RequestId          string
	SecurityGroupTypes ListAvailableConfigSecurityGroupTypeList
	InstanceTypes      ListAvailableConfigInstanceTypeList
	EmrVerTypes        ListAvailableConfigEmrVerTypeList
	ZoneTypes          ListAvailableConfigZoneTypeList
	Vpcs               ListAvailableConfigVpcList
}

type ListAvailableConfigSubModuleList []ListAvailableConfigSubModule

func (list *ListAvailableConfigSubModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigSubModule)
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

type ListAvailableConfigRequiredList []ListAvailableConfigRequired

func (list *ListAvailableConfigRequiredList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigRequired)
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

type ListAvailableConfigOptionalList []ListAvailableConfigOptional

func (list *ListAvailableConfigOptionalList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigOptional)
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

type ListAvailableConfigAvailableResourceList []ListAvailableConfigAvailableResource

func (list *ListAvailableConfigAvailableResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigAvailableResource)
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

type ListAvailableConfigAvailableResourceCreationListList []string

func (list *ListAvailableConfigAvailableResourceCreationListList) UnmarshalJSON(data []byte) error {
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

type ListAvailableConfigAvailableDiskCategoryListList []string

func (list *ListAvailableConfigAvailableDiskCategoryListList) UnmarshalJSON(data []byte) error {
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

type ListAvailableConfigSystemDiskCategoryList []string

func (list *ListAvailableConfigSystemDiskCategoryList) UnmarshalJSON(data []byte) error {
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

type ListAvailableConfigDataDiskCategoryList []string

func (list *ListAvailableConfigDataDiskCategoryList) UnmarshalJSON(data []byte) error {
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

type ListAvailableConfigNetworkTypeList []string

func (list *ListAvailableConfigNetworkTypeList) UnmarshalJSON(data []byte) error {
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

type ListAvailableConfigSupportedInstanceTypeList []string

func (list *ListAvailableConfigSupportedInstanceTypeList) UnmarshalJSON(data []byte) error {
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

type ListAvailableConfigInstanceTypeFamilyList []string

func (list *ListAvailableConfigInstanceTypeFamilyList) UnmarshalJSON(data []byte) error {
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

type ListAvailableConfigInstanceGenerationList []string

func (list *ListAvailableConfigInstanceGenerationList) UnmarshalJSON(data []byte) error {
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

type ListAvailableConfigVSwitchList []ListAvailableConfigVSwitch

func (list *ListAvailableConfigVSwitchList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigVSwitch)
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

type ListAvailableConfigSecurityGroupList []ListAvailableConfigSecurityGroup

func (list *ListAvailableConfigSecurityGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigSecurityGroup)
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

type ListAvailableConfigSecurityGroupTypeList []ListAvailableConfigSecurityGroupType

func (list *ListAvailableConfigSecurityGroupTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigSecurityGroupType)
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

type ListAvailableConfigInstanceTypeList []ListAvailableConfigInstanceType

func (list *ListAvailableConfigInstanceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigInstanceType)
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

type ListAvailableConfigEmrVerTypeList []ListAvailableConfigEmrVerType

func (list *ListAvailableConfigEmrVerTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigEmrVerType)
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

type ListAvailableConfigZoneTypeList []ListAvailableConfigZoneType

func (list *ListAvailableConfigZoneTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigZoneType)
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

type ListAvailableConfigVpcList []ListAvailableConfigVpc

func (list *ListAvailableConfigVpcList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigVpc)
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

func (c *EmrClient) ModifyClusterName(req *ModifyClusterNameArgs) (resp *ModifyClusterNameResponse, err error) {
	resp = &ModifyClusterNameResponse{}
	err = c.Request("ModifyClusterName", req, resp)
	return
}

type ModifyClusterNameArgs struct {
	ResourceOwnerId int64
	Id              string
	Name            string
}
type ModifyClusterNameResponse struct {
	RequestId string
}

func (c *EmrClient) DescribeExecutionPlan(req *DescribeExecutionPlanArgs) (resp *DescribeExecutionPlanResponse, err error) {
	resp = &DescribeExecutionPlanResponse{}
	err = c.Request("DescribeExecutionPlan", req, resp)
	return
}

type DescribeExecutionPlanJobInfo struct {
	Id string
}

type DescribeExecutionPlanClusterInfo struct {
	Name                   string
	ZoneId                 string
	LogEnable              core.Bool
	LogPath                string
	SecurityGroupId        string
	EmrVer                 string
	ClusterType            string
	HighAvailabilityEnable core.Bool
	VpcId                  string
	VSwitchId              string
	NetType                string
	IoOptimized            core.Bool
	InstanceGeneration     string
	Configurations         string
	EcsOrders              DescribeExecutionPlanEcsOrderInfoList
	BootstrapActionList    DescribeExecutionPlanBootstrapActionList
	SoftwareInfo           DescribeExecutionPlanSoftwareInfo
}

type DescribeExecutionPlanEcsOrderInfo struct {
	Index        int
	NodeCount    int
	InstanceType string
	DiskType     string
	DiskCapacity int
	NodeType     string
	DiskCount    int
}

type DescribeExecutionPlanBootstrapAction struct {
	Name string
	Path string
	Arg  string
}

type DescribeExecutionPlanSoftwareInfo struct {
	EmrVer      string
	ClusterType string
	Softwares   DescribeExecutionPlanSoftwareList
}

type DescribeExecutionPlanSoftware struct {
	DisplayName string
	Name        string
	OnlyDisplay core.Bool
	StartTpe    int
	Version     string
	Optional    core.Bool
}
type DescribeExecutionPlanArgs struct {
	ResourceOwnerId int64
	Id              string
}
type DescribeExecutionPlanResponse struct {
	RequestId             string
	Id                    string
	Name                  string
	Strategy              string
	TimeInterval          int
	StartTime             int64
	TimeUnit              string
	CreateClusterOnDemand core.Bool
	ClusterId             string
	ExecutionPlanVersion  int64
	JobInfoList           DescribeExecutionPlanJobInfoList
	ClusterInfo           DescribeExecutionPlanClusterInfo
}

type DescribeExecutionPlanEcsOrderInfoList []DescribeExecutionPlanEcsOrderInfo

func (list *DescribeExecutionPlanEcsOrderInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeExecutionPlanEcsOrderInfo)
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

type DescribeExecutionPlanBootstrapActionList []DescribeExecutionPlanBootstrapAction

func (list *DescribeExecutionPlanBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeExecutionPlanBootstrapAction)
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

type DescribeExecutionPlanSoftwareList []DescribeExecutionPlanSoftware

func (list *DescribeExecutionPlanSoftwareList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeExecutionPlanSoftware)
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

type DescribeExecutionPlanJobInfoList []DescribeExecutionPlanJobInfo

func (list *DescribeExecutionPlanJobInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeExecutionPlanJobInfo)
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

func (c *EmrClient) ListClusters(req *ListClustersArgs) (resp *ListClustersResponse, err error) {
	resp = &ListClustersResponse{}
	err = c.Request("ListClusters", req, resp)
	return
}

type ListClustersClusterInfo struct {
	Id                  string
	Name                string
	Type                string
	CreateTime          int64
	RunningTime         int
	Status              string
	ChargeType          string
	ExpiredTime         int64
	Period              int
	HasUncompletedOrder core.Bool
	OrderList           string
	FailReason          ListClustersFailReason
}

type ListClustersFailReason struct {
	ErrorCode string
	ErrorMsg  string
	RequestId string
}
type ListClustersArgs struct {
	ResourceOwnerId  int64
	CreateType       string
	IsDesc           core.Bool
	PageNumber       int
	PageSize         int
	DefaultStatus    core.Bool
	ClusterTypeLists ListClustersClusterTypeListList
	StatusLists      ListClustersStatusListList
}
type ListClustersResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Clusters   ListClustersClusterInfoList
}

type ListClustersClusterTypeListList []string

func (list *ListClustersClusterTypeListList) UnmarshalJSON(data []byte) error {
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

type ListClustersStatusListList []string

func (list *ListClustersStatusListList) UnmarshalJSON(data []byte) error {
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

type ListClustersClusterInfoList []ListClustersClusterInfo

func (list *ListClustersClusterInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClustersClusterInfo)
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

func (c *EmrClient) ResizeCluster(req *ResizeClusterArgs) (resp *ResizeClusterResponse, err error) {
	resp = &ResizeClusterResponse{}
	err = c.Request("ResizeCluster", req, resp)
	return
}

type ResizeClusterArgs struct {
	ClusterId          string
	NewMasterInstances int
	NewCoreInstances   int
	NewTaskInstances   int
	ChargeType         string
	Period             int
}
type ResizeClusterResponse struct {
	RequestId   string
	ClusterId   string
	EmrOrderId  string
	CoreOrderId string
}

func (c *EmrClient) DeleteExecutionPlan(req *DeleteExecutionPlanArgs) (resp *DeleteExecutionPlanResponse, err error) {
	resp = &DeleteExecutionPlanResponse{}
	err = c.Request("DeleteExecutionPlan", req, resp)
	return
}

type DeleteExecutionPlanArgs struct {
	ResourceOwnerId int64
	Id              string
}
type DeleteExecutionPlanResponse struct {
	RequestId string
}

func (c *EmrClient) ListExecutionPlanInstances(req *ListExecutionPlanInstancesArgs) (resp *ListExecutionPlanInstancesResponse, err error) {
	resp = &ListExecutionPlanInstancesResponse{}
	err = c.Request("ListExecutionPlanInstances", req, resp)
	return
}

type ListExecutionPlanInstancesExecutionPlanInstance struct {
	Id                string
	ExecutionPlanId   string
	ExecutionPlanName string
	StartTime         int64
	RunTime           int
	ClusterId         string
	ClusterName       string
	ClusterType       string
	Status            string
	LogEnable         core.Bool
	LogPath           string
}
type ListExecutionPlanInstancesArgs struct {
	ResourceOwnerId      int64
	OnlyLastInstance     core.Bool
	IsDesc               core.Bool
	PageNumber           int
	PageSize             int
	ExecutionPlanIdLists ListExecutionPlanInstancesExecutionPlanIdListList
	StatusLists          ListExecutionPlanInstancesStatusListList
}
type ListExecutionPlanInstancesResponse struct {
	RequestId              string
	TotalCount             int
	PageNumber             int
	PageSize               int
	ExecutionPlanInstances ListExecutionPlanInstancesExecutionPlanInstanceList
}

type ListExecutionPlanInstancesExecutionPlanIdListList []string

func (list *ListExecutionPlanInstancesExecutionPlanIdListList) UnmarshalJSON(data []byte) error {
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

type ListExecutionPlanInstancesStatusListList []string

func (list *ListExecutionPlanInstancesStatusListList) UnmarshalJSON(data []byte) error {
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

type ListExecutionPlanInstancesExecutionPlanInstanceList []ListExecutionPlanInstancesExecutionPlanInstance

func (list *ListExecutionPlanInstancesExecutionPlanInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListExecutionPlanInstancesExecutionPlanInstance)
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

func (c *EmrClient) CreateExecutionPlan(req *CreateExecutionPlanArgs) (resp *CreateExecutionPlanResponse, err error) {
	resp = &CreateExecutionPlanResponse{}
	err = c.Request("CreateExecutionPlan", req, resp)
	return
}

type CreateExecutionPlanEcsOrder struct {
	Index        int
	NodeCount    int
	NodeType     string
	InstanceType string
	DiskType     string
	DiskCapacity int
	DiskCount    int
}

type CreateExecutionPlanBootstrapAction struct {
	Name string
	Path string
	Arg  string
}
type CreateExecutionPlanArgs struct {
	ResourceOwnerId        int64
	Name                   string
	Strategy               string
	TimeInterval           int
	StartTime              int64
	TimeUnit               string
	ClusterId              string
	CreateClusterOnDemand  core.Bool
	ClusterName            string
	ZoneId                 string
	LogEnable              core.Bool
	LogPath                string
	SecurityGroupId        string
	IsOpenPublicIp         core.Bool
	EmrVer                 string
	ClusterType            string
	HighAvailabilityEnable core.Bool
	VpcId                  string
	VSwitchId              string
	NetType                string
	IoOptimized            core.Bool
	InstanceGeneration     string
	Configurations         string
	JobIdLists             CreateExecutionPlanJobIdListList
	OptionSoftWareLists    CreateExecutionPlanOptionSoftWareListList
	EcsOrders              CreateExecutionPlanEcsOrderList
	BootstrapActions       CreateExecutionPlanBootstrapActionList
}
type CreateExecutionPlanResponse struct {
	RequestId string
	Id        string
}

type CreateExecutionPlanJobIdListList []string

func (list *CreateExecutionPlanJobIdListList) UnmarshalJSON(data []byte) error {
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

type CreateExecutionPlanOptionSoftWareListList []string

func (list *CreateExecutionPlanOptionSoftWareListList) UnmarshalJSON(data []byte) error {
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

type CreateExecutionPlanEcsOrderList []CreateExecutionPlanEcsOrder

func (list *CreateExecutionPlanEcsOrderList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateExecutionPlanEcsOrder)
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

type CreateExecutionPlanBootstrapActionList []CreateExecutionPlanBootstrapAction

func (list *CreateExecutionPlanBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateExecutionPlanBootstrapAction)
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

func (c *EmrClient) ListJobs(req *ListJobsArgs) (resp *ListJobsResponse, err error) {
	resp = &ListJobsResponse{}
	err = c.Request("ListJobs", req, resp)
	return
}

type ListJobsJobInfo struct {
	Id           string
	Name         string
	Type         string
	RunParameter string
	FailAct      string
}
type ListJobsArgs struct {
	ResourceOwnerId int64
	IsDesc          core.Bool
	PageNumber      int
	PageSize        int
}
type ListJobsResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Jobs       ListJobsJobInfoList
}

type ListJobsJobInfoList []ListJobsJobInfo

func (list *ListJobsJobInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobsJobInfo)
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

func (c *EmrClient) RunExecutionPlan(req *RunExecutionPlanArgs) (resp *RunExecutionPlanResponse, err error) {
	resp = &RunExecutionPlanResponse{}
	err = c.Request("RunExecutionPlan", req, resp)
	return
}

type RunExecutionPlanArgs struct {
	ResourceOwnerId int64
	Id              string
}
type RunExecutionPlanResponse struct {
	RequestId               string
	ExecutionPlanInstanceId string
}

func (c *EmrClient) ReleaseCluster(req *ReleaseClusterArgs) (resp *ReleaseClusterResponse, err error) {
	resp = &ReleaseClusterResponse{}
	err = c.Request("ReleaseCluster", req, resp)
	return
}

type ReleaseClusterArgs struct {
	ResourceOwnerId int64
	Id              string
	ForceRelease    core.Bool
}
type ReleaseClusterResponse struct {
	RequestId string
}

func (c *EmrClient) DescribeCluster(req *DescribeClusterArgs) (resp *DescribeClusterResponse, err error) {
	resp = &DescribeClusterResponse{}
	err = c.Request("DescribeCluster", req, resp)
	return
}

type DescribeClusterClusterInfo struct {
	Id                     string
	RegionId               string
	ZoneId                 string
	Name                   string
	CreateType             string
	StartTime              int64
	StopTime               int64
	LogEnable              core.Bool
	LogPath                string
	Status                 string
	HighAvailabilityEnable core.Bool
	ChargeType             string
	ExpiredTime            int64
	Period                 int
	RunningTime            int
	MasterNodeTotal        int
	MasterNodeInService    int
	CoreNodeTotal          int
	CoreNodeInService      int
	TaskNodeTotal          int
	TaskNodeInService      int
	VpcId                  string
	VSwitchId              string
	NetType                string
	IoOptimized            core.Bool
	InstanceGeneration     string
	ImageId                string
	SecurityGroupId        string
	SecurityGroupName      string
	BootstrapFailed        core.Bool
	Configurations         string
	EcsOrderInfoList       DescribeClusterEcsOrderInfoList
	BootstrapActionList    DescribeClusterBootstrapActionList
	FailReason             DescribeClusterFailReason
	SoftwareInfo           DescribeClusterSoftwareInfo
}

type DescribeClusterEcsOrderInfo struct {
	NodeType       string
	InstanceType   string
	CpuCore        int
	MemoryCapacity int
	DiskType       string
	DiskCapacity   int
	DiskCount      int
	BandWidth      string
	Nodes          DescribeClusterNodeList
}

type DescribeClusterNode struct {
	ZoneId         string
	InstanceId     string
	Status         string
	PubIp          string
	InnerIp        string
	ExpiredTime    string
	EmrExpiredTime string
	DaemonInfos    DescribeClusterDaemonInfoList
	DiskInfos      DescribeClusterDiskInfoList
}

type DescribeClusterDaemonInfo struct {
	Name string
}

type DescribeClusterDiskInfo struct {
	Device   string
	DiskName string
	DiskId   string
	Type     string
	Size     int
}

type DescribeClusterBootstrapAction struct {
	Name string
	Path string
	Arg  string
}

type DescribeClusterFailReason struct {
	ErrorCode string
	ErrorMsg  string
	RequestId string
}

type DescribeClusterSoftwareInfo struct {
	EmrVer      string
	ClusterType string
	Softwares   DescribeClusterSoftwareList
}

type DescribeClusterSoftware struct {
	DisplayName string
	Name        string
	OnlyDisplay core.Bool
	StartTpe    int
	Version     string
}
type DescribeClusterArgs struct {
	ResourceOwnerId int64
	Id              string
}
type DescribeClusterResponse struct {
	RequestId   string
	ClusterInfo DescribeClusterClusterInfo
}

type DescribeClusterEcsOrderInfoList []DescribeClusterEcsOrderInfo

func (list *DescribeClusterEcsOrderInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterEcsOrderInfo)
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

type DescribeClusterBootstrapActionList []DescribeClusterBootstrapAction

func (list *DescribeClusterBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterBootstrapAction)
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

type DescribeClusterNodeList []DescribeClusterNode

func (list *DescribeClusterNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterNode)
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

type DescribeClusterDaemonInfoList []DescribeClusterDaemonInfo

func (list *DescribeClusterDaemonInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterDaemonInfo)
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

type DescribeClusterDiskInfoList []DescribeClusterDiskInfo

func (list *DescribeClusterDiskInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterDiskInfo)
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

type DescribeClusterSoftwareList []DescribeClusterSoftware

func (list *DescribeClusterSoftwareList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterSoftware)
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

func (c *EmrClient) KillExecutionJobInstance(req *KillExecutionJobInstanceArgs) (resp *KillExecutionJobInstanceResponse, err error) {
	resp = &KillExecutionJobInstanceResponse{}
	err = c.Request("KillExecutionJobInstance", req, resp)
	return
}

type KillExecutionJobInstanceArgs struct {
	ResourceOwnerId int64
	JobInstanceId   string
}
type KillExecutionJobInstanceResponse struct {
	RequestId string
}

func (c *EmrClient) DescribeJob(req *DescribeJobArgs) (resp *DescribeJobResponse, err error) {
	resp = &DescribeJobResponse{}
	err = c.Request("DescribeJob", req, resp)
	return
}

type DescribeJobArgs struct {
	ResourceOwnerId int64
	Id              string
}
type DescribeJobResponse struct {
	RequestId    string
	Id           string
	Name         string
	FailAct      string
	Type         string
	RunParameter string
}

func (c *EmrClient) GetClusterStatus(req *GetClusterStatusArgs) (resp *GetClusterStatusResponse, err error) {
	resp = &GetClusterStatusResponse{}
	err = c.Request("GetClusterStatus", req, resp)
	return
}

type GetClusterStatusStatus struct {
	Name      string
	Legend    string
	Unit      string
	Series    GetClusterStatusSeriesInfoList
	LineNames GetClusterStatusLineNameList
	Times     GetClusterStatusTimeList
}

type GetClusterStatusSeriesInfo struct {
	SeriesItems GetClusterStatusSeriesItemList
}

type GetClusterStatusSeriesItem struct {
	SeriesValue float32
}
type GetClusterStatusArgs struct {
	ResourceOwnerId int64
	Id              string
	ItemType        string
	Interval        string
}
type GetClusterStatusResponse struct {
	RequestId  string
	StatusList GetClusterStatusStatusList
}

type GetClusterStatusSeriesInfoList []GetClusterStatusSeriesInfo

func (list *GetClusterStatusSeriesInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetClusterStatusSeriesInfo)
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

type GetClusterStatusLineNameList []string

func (list *GetClusterStatusLineNameList) UnmarshalJSON(data []byte) error {
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

type GetClusterStatusTimeList []string

func (list *GetClusterStatusTimeList) UnmarshalJSON(data []byte) error {
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

type GetClusterStatusSeriesItemList []GetClusterStatusSeriesItem

func (list *GetClusterStatusSeriesItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetClusterStatusSeriesItem)
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

type GetClusterStatusStatusList []GetClusterStatusStatus

func (list *GetClusterStatusStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetClusterStatusStatus)
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

func (c *EmrClient) KillExecutionPlanInstance(req *KillExecutionPlanInstanceArgs) (resp *KillExecutionPlanInstanceResponse, err error) {
	resp = &KillExecutionPlanInstanceResponse{}
	err = c.Request("KillExecutionPlanInstance", req, resp)
	return
}

type KillExecutionPlanInstanceArgs struct {
	ResourceOwnerId int64
	Id              string
}
type KillExecutionPlanInstanceResponse struct {
	RequestId string
}

func (c *EmrClient) ListJobExecutionInstances(req *ListJobExecutionInstancesArgs) (resp *ListJobExecutionInstancesResponse, err error) {
	resp = &ListJobExecutionInstancesResponse{}
	err = c.Request("ListJobExecutionInstances", req, resp)
	return
}

type ListJobExecutionInstancesJobInstance struct {
	Id        string
	JobName   string
	StartTime int64
	RunTime   int
	JobType   string
	JobId     string
	ClusterId string
	Status    string
}
type ListJobExecutionInstancesArgs struct {
	ResourceOwnerId         int64
	ExecutionPlanInstanceId string
	IsDesc                  core.Bool
	PageNumber              int
	PageSize                int
}
type ListJobExecutionInstancesResponse struct {
	RequestId    string
	TotalCount   int
	PageNumber   int
	PageSize     int
	JobInstances ListJobExecutionInstancesJobInstanceList
}

type ListJobExecutionInstancesJobInstanceList []ListJobExecutionInstancesJobInstance

func (list *ListJobExecutionInstancesJobInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobExecutionInstancesJobInstance)
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

func (c *EmrClient) ModifyJob(req *ModifyJobArgs) (resp *ModifyJobResponse, err error) {
	resp = &ModifyJobResponse{}
	err = c.Request("ModifyJob", req, resp)
	return
}

type ModifyJobArgs struct {
	ResourceOwnerId int64
	Id              string
	Name            string
	Type            string
	RunParameter    string
	FailAct         string
}
type ModifyJobResponse struct {
	RequestId string
}

func (c *EmrClient) ListJobInstanceWorkers(req *ListJobInstanceWorkersArgs) (resp *ListJobInstanceWorkersResponse, err error) {
	resp = &ListJobInstanceWorkersResponse{}
	err = c.Request("ListJobInstanceWorkers", req, resp)
	return
}

type ListJobInstanceWorkersJobInstanceWorkerInfo struct {
	ApplicationId string
	InstanceInfo  string
	ContainerInfo string
}
type ListJobInstanceWorkersArgs struct {
	ResourceOwnerId int64
	JobInstanceId   string
}
type ListJobInstanceWorkersResponse struct {
	RequestId          string
	JobInstanceWorkers ListJobInstanceWorkersJobInstanceWorkerInfoList
}

type ListJobInstanceWorkersJobInstanceWorkerInfoList []ListJobInstanceWorkersJobInstanceWorkerInfo

func (list *ListJobInstanceWorkersJobInstanceWorkerInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobInstanceWorkersJobInstanceWorkerInfo)
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

func (c *EmrClient) QueryPrice(req *QueryPriceArgs) (resp *QueryPriceResponse, err error) {
	resp = &QueryPriceResponse{}
	err = c.Request("QueryPrice", req, resp)
	return
}

type QueryPriceEmrPriceDO struct {
	OriginalPrice string
	DiscountPrice string
	TradePrice    string
}

type QueryPriceEcsPriceDO struct {
	OriginalPrice string
	DiscountPrice string
	TradePrice    string
}
type QueryPriceArgs struct {
	ResourceOwnerId        int64
	ZoneId                 string
	MasterInstanceType     string
	CoreInstanceType       string
	TaskInstanceType       string
	MasterInstanceQuantity int
	CoreInstanceQuantity   int
	TaskInstanceQuantity   int
	MasterDiskType         string
	CoreDiskType           string
	TaskDiskType           string
	MasterDiskQuantity     int
	CoreDiskQuantity       int
	TaskDiskQuantity       int
	Duration               int
	IoOptimized            core.Bool
	ChargeType             string
	NetType                string
	Period                 int
}
type QueryPriceResponse struct {
	RequestId  string
	EmrPrice   string
	EcsPrice   string
	EmrPriceDO QueryPriceEmrPriceDO
	EcsPriceDO QueryPriceEcsPriceDO
}

func (c *EmrClient) ResumeExecutionPlanScheduler(req *ResumeExecutionPlanSchedulerArgs) (resp *ResumeExecutionPlanSchedulerResponse, err error) {
	resp = &ResumeExecutionPlanSchedulerResponse{}
	err = c.Request("ResumeExecutionPlanScheduler", req, resp)
	return
}

type ResumeExecutionPlanSchedulerArgs struct {
	ResourceOwnerId int64
	Id              string
}
type ResumeExecutionPlanSchedulerResponse struct {
	RequestId string
}

func (c *EmrClient) ModifyExecutionPlan(req *ModifyExecutionPlanArgs) (resp *ModifyExecutionPlanResponse, err error) {
	resp = &ModifyExecutionPlanResponse{}
	err = c.Request("ModifyExecutionPlan", req, resp)
	return
}

type ModifyExecutionPlanEcsOrder struct {
	Index        int
	NodeCount    int
	InstanceType string
	DiskType     string
	DiskCapacity int
	NodeType     string
	DiskCount    int
}

type ModifyExecutionPlanBootstrapAction struct {
	Name string
	Path string
	Arg  string
}
type ModifyExecutionPlanArgs struct {
	ResourceOwnerId        int64
	ClusterName            string
	ClusterId              string
	ZoneId                 string
	LogEnable              core.Bool
	LogPath                string
	SecurityGroupId        string
	IsOpenPublicIp         core.Bool
	CreateClusterOnDemand  core.Bool
	EmrVer                 string
	ClusterType            string
	HighAvailabilityEnable core.Bool
	VpcId                  string
	VSwitchId              string
	NetType                string
	IoOptimized            core.Bool
	InstanceGeneration     string
	Configurations         string
	Id                     string
	ExecutionPlanVersion   int64
	Name                   string
	Strategy               string
	TimeInterval           int
	StartTime              int64
	TimeUnit               string
	OptionSoftWareLists    ModifyExecutionPlanOptionSoftWareListList
	EcsOrders              ModifyExecutionPlanEcsOrderList
	BootstrapActions       ModifyExecutionPlanBootstrapActionList
	JobIdLists             ModifyExecutionPlanJobIdListList
}
type ModifyExecutionPlanResponse struct {
	RequestId string
}

type ModifyExecutionPlanOptionSoftWareListList []string

func (list *ModifyExecutionPlanOptionSoftWareListList) UnmarshalJSON(data []byte) error {
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

type ModifyExecutionPlanEcsOrderList []ModifyExecutionPlanEcsOrder

func (list *ModifyExecutionPlanEcsOrderList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyExecutionPlanEcsOrder)
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

type ModifyExecutionPlanBootstrapActionList []ModifyExecutionPlanBootstrapAction

func (list *ModifyExecutionPlanBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyExecutionPlanBootstrapAction)
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

type ModifyExecutionPlanJobIdListList []string

func (list *ModifyExecutionPlanJobIdListList) UnmarshalJSON(data []byte) error {
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

func (c *EmrClient) DeleteJob(req *DeleteJobArgs) (resp *DeleteJobResponse, err error) {
	resp = &DeleteJobResponse{}
	err = c.Request("DeleteJob", req, resp)
	return
}

type DeleteJobArgs struct {
	ResourceOwnerId int64
	Id              string
}
type DeleteJobResponse struct {
	RequestId string
}

func (c *EmrClient) CreateCluster(req *CreateClusterArgs) (resp *CreateClusterResponse, err error) {
	resp = &CreateClusterResponse{}
	err = c.Request("CreateCluster", req, resp)
	return
}

type CreateClusterEcsOrder struct {
	Index        int
	NodeCount    int
	NodeType     string
	InstanceType string
	DiskType     string
	DiskCapacity int
	DiskCount    int
}

type CreateClusterBootstrapAction struct {
	Name string
	Path string
	Arg  string
}
type CreateClusterArgs struct {
	ResourceOwnerId        int64
	Name                   string
	ZoneId                 string
	LogEnable              core.Bool
	LogPath                string
	SecurityGroupId        string
	IsOpenPublicIp         core.Bool
	SecurityGroupName      string
	ChargeType             string
	Period                 int
	AutoRenew              core.Bool
	AutoRenewPeriod        int
	VpcId                  string
	VSwitchId              string
	NetType                string
	EmrVer                 string
	ClusterType            string
	HighAvailabilityEnable core.Bool
	IoOptimized            core.Bool
	InstanceGeneration     string
	MasterPwdEnable        core.Bool
	MasterPwd              string
	Configurations         string
	OptionSoftWareLists    CreateClusterOptionSoftWareListList
	EcsOrders              CreateClusterEcsOrderList
	BootstrapActions       CreateClusterBootstrapActionList
}
type CreateClusterResponse struct {
	RequestId     string
	ClusterId     string
	EmrOrderId    string
	MasterOrderId string
	CoreOrderId   string
}

type CreateClusterOptionSoftWareListList []string

func (list *CreateClusterOptionSoftWareListList) UnmarshalJSON(data []byte) error {
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

type CreateClusterEcsOrderList []CreateClusterEcsOrder

func (list *CreateClusterEcsOrderList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateClusterEcsOrder)
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

type CreateClusterBootstrapActionList []CreateClusterBootstrapAction

func (list *CreateClusterBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateClusterBootstrapAction)
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

func (c *EmrClient) ListExecutionPlans(req *ListExecutionPlansArgs) (resp *ListExecutionPlansResponse, err error) {
	resp = &ListExecutionPlansResponse{}
	err = c.Request("ListExecutionPlans", req, resp)
	return
}

type ListExecutionPlansExecutionPlanInfo struct {
	Id                    string
	Name                  string
	CreateClusterOnDemand core.Bool
	Stragety              string
	Status                string
	TimeInterval          int
	StartTime             int64
	TimeUnit              string
}
type ListExecutionPlansArgs struct {
	ResourceOwnerId int64
	ClusterId       string
	JobId           string
	Strategy        string
	IsDesc          core.Bool
	PageNumber      int
	PageSize        int
	StatusLists     ListExecutionPlansStatusListList
}
type ListExecutionPlansResponse struct {
	RequestId      string
	TotalCount     int
	PageNumber     int
	PageSize       int
	ExecutionPlans ListExecutionPlansExecutionPlanInfoList
}

type ListExecutionPlansStatusListList []string

func (list *ListExecutionPlansStatusListList) UnmarshalJSON(data []byte) error {
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

type ListExecutionPlansExecutionPlanInfoList []ListExecutionPlansExecutionPlanInfo

func (list *ListExecutionPlansExecutionPlanInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListExecutionPlansExecutionPlanInfo)
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
