package ess

import "encoding/json"

func (c *EssClient) DescribeLimitation(req *DescribeLimitationArgs) (resp *DescribeLimitationResponse, err error) {
	resp = &DescribeLimitationResponse{}
	err = c.Request("DescribeLimitation", req, resp)
	return
}

type DescribeLimitationArgs struct {
	ResourceOwnerAccount string
	OwnerId              int64
}
type DescribeLimitationResponse struct {
	MaxNumberOfScalingGroups         int
	MaxNumberOfScalingConfigurations int
	MaxNumberOfScalingRules          int
	MaxNumberOfScheduledTasks        int
	MaxNumberOfScalingInstances      int
	MaxNumberOfDBInstances           int
	MaxNumberOfLoadBalancers         int
	MaxNumberOfMinSize               int
	MaxNumberOfMaxSize               int
}

func (c *EssClient) DescribeCapacityHistory(req *DescribeCapacityHistoryArgs) (resp *DescribeCapacityHistoryResponse, err error) {
	resp = &DescribeCapacityHistoryResponse{}
	err = c.Request("DescribeCapacityHistory", req, resp)
	return
}

type DescribeCapacityHistoryCapacityHistoryModel struct {
	ScalingGroupId      string
	TotalCapacity       int
	AttachedCapacity    int
	AutoCreatedCapacity int
	Timestamp           string
}
type DescribeCapacityHistoryArgs struct {
	ResourceOwnerAccount string
	ScalingGroupId       string
	PageSize             int
	EndTime              string
	StartTime            string
	OwnerId              int64
	PageNumber           int
}
type DescribeCapacityHistoryResponse struct {
	TotalCount           int
	PageNumber           int
	PageSize             int
	CapacityHistoryItems DescribeCapacityHistoryCapacityHistoryModelList
}

type DescribeCapacityHistoryCapacityHistoryModelList []DescribeCapacityHistoryCapacityHistoryModel

func (list *DescribeCapacityHistoryCapacityHistoryModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCapacityHistoryCapacityHistoryModel)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

func (c *EssClient) DescribeScalingInstances(req *DescribeScalingInstancesArgs) (resp *DescribeScalingInstancesResponse, err error) {
	resp = &DescribeScalingInstancesResponse{}
	err = c.Request("DescribeScalingInstances", req, resp)
	return
}

type DescribeScalingInstancesScalingInstance struct {
	InstanceId             string
	ScalingConfigurationId string
	ScalingGroupId         string
	HealthStatus           string
	LoadBalancerWeight     int
	LifecycleState         string
	CreationTime           string
	CreationType           string
}
type DescribeScalingInstancesArgs struct {
	InstanceId10           string
	ResourceOwnerId        int64
	InstanceId12           string
	InstanceId11           string
	ScalingGroupId         string
	LifecycleState         string
	CreationType           string
	PageNumber             int
	PageSize               int
	InstanceId20           string
	InstanceId1            string
	InstanceId3            string
	ResourceOwnerAccount   string
	InstanceId2            string
	InstanceId5            string
	InstanceId4            string
	OwnerAccount           string
	InstanceId7            string
	InstanceId6            string
	InstanceId9            string
	InstanceId8            string
	OwnerId                int64
	ScalingConfigurationId string
	HealthStatus           string
	InstanceId18           string
	InstanceId17           string
	InstanceId19           string
	InstanceId14           string
	InstanceId13           string
	InstanceId16           string
	InstanceId15           string
}
type DescribeScalingInstancesResponse struct {
	TotalCount       int
	PageNumber       int
	PageSize         int
	RequestId        string
	ScalingInstances DescribeScalingInstancesScalingInstanceList
}

type DescribeScalingInstancesScalingInstanceList []DescribeScalingInstancesScalingInstance

func (list *DescribeScalingInstancesScalingInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingInstancesScalingInstance)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

func (c *EssClient) DeleteScalingGroup(req *DeleteScalingGroupArgs) (resp *DeleteScalingGroupResponse, err error) {
	resp = &DeleteScalingGroupResponse{}
	err = c.Request("DeleteScalingGroup", req, resp)
	return
}

type DeleteScalingGroupArgs struct {
	ResourceOwnerAccount string
	ScalingGroupId       string
	ForceDelete          bool
	OwnerAccount         string
	OwnerId              int64
}
type DeleteScalingGroupResponse struct {
	RequestId string
}

func (c *EssClient) DescribeScalingActivityDetail(req *DescribeScalingActivityDetailArgs) (resp *DescribeScalingActivityDetailResponse, err error) {
	resp = &DescribeScalingActivityDetailResponse{}
	err = c.Request("DescribeScalingActivityDetail", req, resp)
	return
}

type DescribeScalingActivityDetailArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerId              int64
	ScalingActivityId    string
}
type DescribeScalingActivityDetailResponse struct {
	ScalingActivityId string
	Detail            string
}

func (c *EssClient) DescribeAlertConfig(req *DescribeAlertConfigArgs) (resp *DescribeAlertConfigResponse, err error) {
	resp = &DescribeAlertConfigResponse{}
	err = c.Request("DescribeAlertConfig", req, resp)
	return
}

type DescribeAlertConfigArgs struct {
	ResourceOwnerAccount string
	ScalingGroupId       string
	OwnerId              int64
}
type DescribeAlertConfigResponse struct {
	SuccessConfig int
	FailConfig    int
	RejectConfig  int
	RequestId     string
}

func (c *EssClient) VerifyUser(req *VerifyUserArgs) (resp *VerifyUserResponse, err error) {
	resp = &VerifyUserResponse{}
	err = c.Request("VerifyUser", req, resp)
	return
}

type VerifyUserArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerId              int64
}
type VerifyUserResponse struct {
}

func (c *EssClient) DescribeScheduledTasks(req *DescribeScheduledTasksArgs) (resp *DescribeScheduledTasksResponse, err error) {
	resp = &DescribeScheduledTasksResponse{}
	err = c.Request("DescribeScheduledTasks", req, resp)
	return
}

type DescribeScheduledTasksScheduledTask struct {
	ScheduledTaskId      string
	ScheduledTaskName    string
	Description          string
	ScheduledAction      string
	RecurrenceEndTime    string
	LaunchTime           string
	RecurrenceType       string
	RecurrenceValue      string
	LaunchExpirationTime int
	TaskEnabled          bool
}
type DescribeScheduledTasksArgs struct {
	ResourceOwnerId      int64
	ScheduledAction2     string
	ScheduledAction1     string
	ScheduledAction6     string
	ScheduledAction5     string
	ScheduledAction4     string
	ScheduledAction3     string
	ScheduledAction9     string
	ScheduledAction8     string
	ScheduledAction7     string
	OwnerId              int64
	ScheduledTaskName20  string
	ScheduledTaskName19  string
	ScheduledTaskName18  string
	ScheduledTaskId20    string
	ScheduledTaskName13  string
	ScheduledTaskName12  string
	ScheduledTaskName11  string
	ScheduledTaskName10  string
	ScheduledTaskName17  string
	ScheduledTaskName16  string
	PageNumber           int
	ScheduledTaskName15  string
	ScheduledTaskName14  string
	ScheduledTaskId2     string
	ScheduledTaskId1     string
	ScheduledTaskId4     string
	ScheduledTaskId18    string
	ScheduledTaskId3     string
	ScheduledTaskId19    string
	ScheduledTaskId6     string
	ScheduledTaskId5     string
	ScheduledTaskId8     string
	ScheduledTaskName9   string
	ScheduledAction20    string
	ScheduledTaskId7     string
	PageSize             int
	ScheduledTaskId12    string
	ScheduledTaskName7   string
	ScheduledTaskId9     string
	ScheduledTaskId13    string
	ScheduledTaskName8   string
	ScheduledTaskId10    string
	ScheduledTaskName5   string
	ScheduledTaskId11    string
	ScheduledTaskName6   string
	ScheduledTaskId16    string
	ScheduledTaskName3   string
	ScheduledTaskId17    string
	ScheduledTaskName4   string
	ScheduledTaskId14    string
	ScheduledTaskName1   string
	ScheduledTaskId15    string
	ScheduledTaskName2   string
	ResourceOwnerAccount string
	OwnerAccount         string
	ScheduledAction18    string
	ScheduledAction19    string
	ScheduledAction16    string
	ScheduledAction17    string
	ScheduledAction14    string
	ScheduledAction15    string
	ScheduledAction12    string
	ScheduledAction13    string
	ScheduledAction10    string
	ScheduledAction11    string
}
type DescribeScheduledTasksResponse struct {
	TotalCount     int
	PageNumber     int
	PageSize       int
	RequestId      string
	ScheduledTasks DescribeScheduledTasksScheduledTaskList
}

type DescribeScheduledTasksScheduledTaskList []DescribeScheduledTasksScheduledTask

func (list *DescribeScheduledTasksScheduledTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScheduledTasksScheduledTask)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

func (c *EssClient) AttachInstances(req *AttachInstancesArgs) (resp *AttachInstancesResponse, err error) {
	resp = &AttachInstancesResponse{}
	err = c.Request("AttachInstances", req, resp)
	return
}

type AttachInstancesArgs struct {
	InstanceId10         string
	LoadBalancerWeight6  int
	LoadBalancerWeight11 int
	LoadBalancerWeight7  int
	LoadBalancerWeight12 int
	ResourceOwnerId      int64
	InstanceId12         string
	LoadBalancerWeight8  int
	InstanceId11         string
	LoadBalancerWeight9  int
	LoadBalancerWeight10 int
	LoadBalancerWeight2  int
	LoadBalancerWeight15 int
	LoadBalancerWeight3  int
	LoadBalancerWeight16 int
	ScalingGroupId       string
	LoadBalancerWeight4  int
	LoadBalancerWeight13 int
	LoadBalancerWeight5  int
	LoadBalancerWeight14 int
	LoadBalancerWeight1  int
	InstanceId20         string
	InstanceId1          string
	LoadBalancerWeight20 int
	InstanceId3          string
	ResourceOwnerAccount string
	InstanceId2          string
	InstanceId5          string
	InstanceId4          string
	OwnerAccount         string
	InstanceId7          string
	InstanceId6          string
	InstanceId9          string
	InstanceId8          string
	OwnerId              int64
	InstanceId18         string
	LoadBalancerWeight19 int
	InstanceId17         string
	LoadBalancerWeight17 int
	InstanceId19         string
	LoadBalancerWeight18 int
	InstanceId14         string
	InstanceId13         string
	InstanceId16         string
	InstanceId15         string
}
type AttachInstancesResponse struct {
	ScalingActivityId string
	RequestId         string
}

func (c *EssClient) ModifyScheduledTask(req *ModifyScheduledTaskArgs) (resp *ModifyScheduledTaskResponse, err error) {
	resp = &ModifyScheduledTaskResponse{}
	err = c.Request("ModifyScheduledTask", req, resp)
	return
}

type ModifyScheduledTaskArgs struct {
	LaunchTime           string
	ResourceOwnerId      int64
	ScheduledAction      string
	ResourceOwnerAccount string
	OwnerAccount         string
	Description          string
	OwnerId              int64
	RecurrenceValue      string
	LaunchExpirationTime int
	RecurrenceEndTime    string
	ScheduledTaskName    string
	TaskEnabled          bool
	ScheduledTaskId      string
	RecurrenceType       string
}
type ModifyScheduledTaskResponse struct {
	RequestId string
}

func (c *EssClient) DescribeScalingActivities(req *DescribeScalingActivitiesArgs) (resp *DescribeScalingActivitiesResponse, err error) {
	resp = &DescribeScalingActivitiesResponse{}
	err = c.Request("DescribeScalingActivities", req, resp)
	return
}

type DescribeScalingActivitiesScalingActivity struct {
	ScalingActivityId   string
	ScalingGroupId      string
	Description         string
	Cause               string
	StartTime           string
	EndTime             string
	Progress            int
	StatusCode          string
	StatusMessage       string
	TotalCapacity       string
	AttachedCapacity    string
	AutoCreatedCapacity string
}
type DescribeScalingActivitiesArgs struct {
	ScalingActivityId9   string
	ResourceOwnerId      int64
	ScalingActivityId5   string
	ScalingActivityId6   string
	ScalingGroupId       string
	ScalingActivityId7   string
	ScalingActivityId8   string
	ScalingActivityId1   string
	ScalingActivityId2   string
	ScalingActivityId3   string
	ScalingActivityId4   string
	PageNumber           int
	StatusCode           string
	PageSize             int
	ScalingActivityId11  string
	ScalingActivityId10  string
	ScalingActivityId13  string
	ScalingActivityId12  string
	ScalingActivityId15  string
	ScalingActivityId14  string
	ScalingActivityId17  string
	ScalingActivityId16  string
	ScalingActivityId19  string
	ResourceOwnerAccount string
	ScalingActivityId18  string
	OwnerAccount         string
	OwnerId              int64
	ScalingActivityId20  string
}
type DescribeScalingActivitiesResponse struct {
	TotalCount        int
	PageNumber        int
	PageSize          int
	RequestId         string
	ScalingActivities DescribeScalingActivitiesScalingActivityList
}

type DescribeScalingActivitiesScalingActivityList []DescribeScalingActivitiesScalingActivity

func (list *DescribeScalingActivitiesScalingActivityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingActivitiesScalingActivity)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

func (c *EssClient) ModifyAlertConfig(req *ModifyAlertConfigArgs) (resp *ModifyAlertConfigResponse, err error) {
	resp = &ModifyAlertConfigResponse{}
	err = c.Request("ModifyAlertConfig", req, resp)
	return
}

type ModifyAlertConfigArgs struct {
	SuccessConfig        int
	RejectConfig         int
	ResourceOwnerAccount string
	ScalingGroupId       string
	OwnerId              int64
	FailConfig           int
}
type ModifyAlertConfigResponse struct {
	RequestId string
}

func (c *EssClient) DescribeAccountAttributes(req *DescribeAccountAttributesArgs) (resp *DescribeAccountAttributesResponse, err error) {
	resp = &DescribeAccountAttributesResponse{}
	err = c.Request("DescribeAccountAttributes", req, resp)
	return
}

type DescribeAccountAttributesArgs struct {
	ResourceOwnerAccount string
	OwnerId              int64
}
type DescribeAccountAttributesResponse struct {
	MaxNumberOfScalingGroups         int
	MaxNumberOfScalingConfigurations int
	MaxNumberOfScalingRules          int
	MaxNumberOfScheduledTasks        int
	MaxNumberOfScalingInstances      int
	MaxNumberOfDBInstances           int
	MaxNumberOfLoadBalancers         int
	MaxNumberOfMinSize               int
	MaxNumberOfMaxSize               int
}

func (c *EssClient) ExecuteScalingRule(req *ExecuteScalingRuleArgs) (resp *ExecuteScalingRuleResponse, err error) {
	resp = &ExecuteScalingRuleResponse{}
	err = c.Request("ExecuteScalingRule", req, resp)
	return
}

type ExecuteScalingRuleArgs struct {
	ResourceOwnerId      int64
	ScalingRuleAri       string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	OwnerId              int64
}
type ExecuteScalingRuleResponse struct {
	ScalingActivityId string
	RequestId         string
}

func (c *EssClient) DetachInstances(req *DetachInstancesArgs) (resp *DetachInstancesResponse, err error) {
	resp = &DetachInstancesResponse{}
	err = c.Request("DetachInstances", req, resp)
	return
}

type DetachInstancesArgs struct {
	InstanceId10         string
	ResourceOwnerId      int64
	InstanceId12         string
	InstanceId11         string
	ScalingGroupId       string
	InstanceId20         string
	InstanceId1          string
	InstanceId3          string
	ResourceOwnerAccount string
	InstanceId2          string
	InstanceId5          string
	InstanceId4          string
	OwnerAccount         string
	InstanceId7          string
	InstanceId6          string
	InstanceId9          string
	InstanceId8          string
	OwnerId              int64
	InstanceId18         string
	InstanceId17         string
	InstanceId19         string
	InstanceId14         string
	InstanceId13         string
	InstanceId16         string
	InstanceId15         string
}
type DetachInstancesResponse struct {
	ScalingActivityId string
	RequestId         string
}

func (c *EssClient) ModifyScalingGroup(req *ModifyScalingGroupArgs) (resp *ModifyScalingGroupResponse, err error) {
	resp = &ModifyScalingGroupResponse{}
	err = c.Request("ModifyScalingGroup", req, resp)
	return
}

type ModifyScalingGroupArgs struct {
	ResourceOwnerId              int64
	ResourceOwnerAccount         string
	ScalingGroupName             string
	ScalingGroupId               string
	OwnerAccount                 string
	ActiveScalingConfigurationId string
	MinSize                      int
	OwnerId                      int64
	MaxSize                      int
	DefaultCooldown              int
	RemovalPolicy1               string
	RemovalPolicy2               string
}
type ModifyScalingGroupResponse struct {
	RequestId string
}

func (c *EssClient) VerifyAuthentication(req *VerifyAuthenticationArgs) (resp *VerifyAuthenticationResponse, err error) {
	resp = &VerifyAuthenticationResponse{}
	err = c.Request("VerifyAuthentication", req, resp)
	return
}

type VerifyAuthenticationArgs struct {
	Uid                  int64
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerId              int64
}
type VerifyAuthenticationResponse struct {
	RequestId string
}

func (c *EssClient) DeleteScalingRule(req *DeleteScalingRuleArgs) (resp *DeleteScalingRuleResponse, err error) {
	resp = &DeleteScalingRuleResponse{}
	err = c.Request("DeleteScalingRule", req, resp)
	return
}

type DeleteScalingRuleArgs struct {
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	ScalingRuleId        string
}
type DeleteScalingRuleResponse struct {
	RequestId string
}

func (c *EssClient) DisableScalingGroup(req *DisableScalingGroupArgs) (resp *DisableScalingGroupResponse, err error) {
	resp = &DisableScalingGroupResponse{}
	err = c.Request("DisableScalingGroup", req, resp)
	return
}

type DisableScalingGroupArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ScalingGroupId       string
	OwnerAccount         string
	OwnerId              int64
}
type DisableScalingGroupResponse struct {
	RequestId string
}

func (c *EssClient) ModifyScalingRule(req *ModifyScalingRuleArgs) (resp *ModifyScalingRuleResponse, err error) {
	resp = &ModifyScalingRuleResponse{}
	err = c.Request("ModifyScalingRule", req, resp)
	return
}

type ModifyScalingRuleArgs struct {
	ScalingRuleName      string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	AdjustmentValue      int
	OwnerAccount         string
	Cooldown             int
	AdjustmentType       string
	OwnerId              int64
	ScalingRuleId        string
}
type ModifyScalingRuleResponse struct {
	RequestId string
}

func (c *EssClient) RemoveInstances(req *RemoveInstancesArgs) (resp *RemoveInstancesResponse, err error) {
	resp = &RemoveInstancesResponse{}
	err = c.Request("RemoveInstances", req, resp)
	return
}

type RemoveInstancesArgs struct {
	InstanceId10         string
	ResourceOwnerId      int64
	InstanceId12         string
	InstanceId11         string
	ScalingGroupId       string
	InstanceId20         string
	InstanceId1          string
	InstanceId3          string
	ResourceOwnerAccount string
	InstanceId2          string
	InstanceId5          string
	InstanceId4          string
	OwnerAccount         string
	InstanceId7          string
	InstanceId6          string
	InstanceId9          string
	InstanceId8          string
	OwnerId              int64
	InstanceId18         string
	InstanceId17         string
	InstanceId19         string
	InstanceId14         string
	InstanceId13         string
	InstanceId16         string
	InstanceId15         string
}
type RemoveInstancesResponse struct {
	ScalingActivityId string
	RequestId         string
}

func (c *EssClient) CreateScalingConfiguration(req *CreateScalingConfigurationArgs) (resp *CreateScalingConfigurationResponse, err error) {
	resp = &CreateScalingConfigurationResponse{}
	err = c.Request("CreateScalingConfiguration", req, resp)
	return
}

type CreateScalingConfigurationArgs struct {
	DataDisk3Size               int
	ImageId                     string
	DataDisk1SnapshotId         string
	DataDisk3Category           string
	DataDisk1Device             string
	ScalingGroupId              string
	DataDisk2Device             string
	InstanceTypess              CreateScalingConfigurationInstanceTypesList
	IoOptimized                 string
	SecurityGroupId             string
	InternetMaxBandwidthOut     int
	SecurityEnhancementStrategy string
	KeyPairName                 string
	SystemDiskCategory          string
	UserData                    string
	DataDisk4Category           string
	DataDisk2SnapshotId         string
	DataDisk4Size               int
	InstanceType                string
	DataDisk2Category           string
	DataDisk1Size               int
	DataDisk3SnapshotId         string
	ResourceOwnerAccount        string
	OwnerAccount                string
	DataDisk2Size               int
	RamRoleName                 string
	OwnerId                     int64
	ScalingConfigurationName    string
	Tags                        string
	DataDisk2DeleteWithInstance string
	DataDisk1Category           string
	DataDisk3DeleteWithInstance string
	LoadBalancerWeight          int
	InstanceName                string
	SystemDiskSize              int
	DataDisk4SnapshotId         string
	DataDisk4Device             string
	InternetChargeType          string
	DataDisk3Device             string
	DataDisk4DeleteWithInstance string
	InternetMaxBandwidthIn      int
	DataDisk1DeleteWithInstance string
}
type CreateScalingConfigurationResponse struct {
	ScalingConfigurationId string
	RequestId              string
}

type CreateScalingConfigurationInstanceTypesList []string

func (list *CreateScalingConfigurationInstanceTypesList) UnmarshalJSON(data []byte) error {
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

func (c *EssClient) DeleteScalingConfiguration(req *DeleteScalingConfigurationArgs) (resp *DeleteScalingConfigurationResponse, err error) {
	resp = &DeleteScalingConfigurationResponse{}
	err = c.Request("DeleteScalingConfiguration", req, resp)
	return
}

type DeleteScalingConfigurationArgs struct {
	ScalingConfigurationId string
	ResourceOwnerAccount   string
	OwnerAccount           string
	OwnerId                int64
}
type DeleteScalingConfigurationResponse struct {
	RequestId string
}

func (c *EssClient) CreateScalingRule(req *CreateScalingRuleArgs) (resp *CreateScalingRuleResponse, err error) {
	resp = &CreateScalingRuleResponse{}
	err = c.Request("CreateScalingRule", req, resp)
	return
}

type CreateScalingRuleArgs struct {
	ScalingRuleName      string
	ResourceOwnerAccount string
	AdjustmentValue      int
	ScalingGroupId       string
	OwnerAccount         string
	Cooldown             int
	AdjustmentType       string
	OwnerId              int64
}
type CreateScalingRuleResponse struct {
	ScalingRuleId  string
	ScalingRuleAri string
	RequestId      string
}

func (c *EssClient) DeactivateScalingConfiguration(req *DeactivateScalingConfigurationArgs) (resp *DeactivateScalingConfigurationResponse, err error) {
	resp = &DeactivateScalingConfigurationResponse{}
	err = c.Request("DeactivateScalingConfiguration", req, resp)
	return
}

type DeactivateScalingConfigurationArgs struct {
	ScalingConfigurationId string
	ResourceOwnerAccount   string
	OwnerAccount           string
	OwnerId                int64
}
type DeactivateScalingConfigurationResponse struct {
	RequestId string
}

func (c *EssClient) DescribeScalingConfigurations(req *DescribeScalingConfigurationsArgs) (resp *DescribeScalingConfigurationsResponse, err error) {
	resp = &DescribeScalingConfigurationsResponse{}
	err = c.Request("DescribeScalingConfigurations", req, resp)
	return
}

type DescribeScalingConfigurationsScalingConfiguration struct {
	ScalingConfigurationId      string
	ScalingConfigurationName    string
	ScalingGroupId              string
	InstanceName                string
	ImageId                     string
	InstanceType                string
	InstanceGeneration          string
	SecurityGroupId             string
	IoOptimized                 string
	InternetChargeType          string
	InternetMaxBandwidthIn      int
	InternetMaxBandwidthOut     int
	SystemDiskCategory          string
	SystemDiskSize              int
	LifecycleState              string
	CreationTime                string
	LoadBalancerWeight          int
	UserData                    string
	KeyPairName                 string
	RamRoleName                 string
	DeploymentSetId             string
	SecurityEnhancementStrategy string
	DataDisks                   DescribeScalingConfigurationsDataDiskList
	Tags                        DescribeScalingConfigurationsTagList
	InstanceTypes               DescribeScalingConfigurationsInstanceTypeList
}

type DescribeScalingConfigurationsDataDisk struct {
	Size       int
	Category   string
	SnapshotId string
	Device     string
}

type DescribeScalingConfigurationsTag struct {
	Key   string
	Value string
}
type DescribeScalingConfigurationsArgs struct {
	ScalingConfigurationId6    string
	ScalingConfigurationId7    string
	ResourceOwnerId            int64
	ScalingConfigurationId4    string
	ScalingConfigurationId5    string
	ScalingGroupId             string
	ScalingConfigurationId8    string
	ScalingConfigurationId9    string
	ScalingConfigurationId10   string
	PageNumber                 int
	ScalingConfigurationName2  string
	ScalingConfigurationName3  string
	ScalingConfigurationName1  string
	PageSize                   int
	ScalingConfigurationId2    string
	ScalingConfigurationId3    string
	ScalingConfigurationId1    string
	ResourceOwnerAccount       string
	OwnerAccount               string
	ScalingConfigurationName6  string
	ScalingConfigurationName7  string
	ScalingConfigurationName4  string
	ScalingConfigurationName5  string
	OwnerId                    int64
	ScalingConfigurationName8  string
	ScalingConfigurationName9  string
	ScalingConfigurationName10 string
}
type DescribeScalingConfigurationsResponse struct {
	TotalCount            int
	PageNumber            int
	PageSize              int
	RequestId             string
	ScalingConfigurations DescribeScalingConfigurationsScalingConfigurationList
}

type DescribeScalingConfigurationsDataDiskList []DescribeScalingConfigurationsDataDisk

func (list *DescribeScalingConfigurationsDataDiskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingConfigurationsDataDisk)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type DescribeScalingConfigurationsTagList []DescribeScalingConfigurationsTag

func (list *DescribeScalingConfigurationsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingConfigurationsTag)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type DescribeScalingConfigurationsInstanceTypeList []string

func (list *DescribeScalingConfigurationsInstanceTypeList) UnmarshalJSON(data []byte) error {
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

type DescribeScalingConfigurationsScalingConfigurationList []DescribeScalingConfigurationsScalingConfiguration

func (list *DescribeScalingConfigurationsScalingConfigurationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingConfigurationsScalingConfiguration)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

func (c *EssClient) DeleteScheduledTask(req *DeleteScheduledTaskArgs) (resp *DeleteScheduledTaskResponse, err error) {
	resp = &DeleteScheduledTaskResponse{}
	err = c.Request("DeleteScheduledTask", req, resp)
	return
}

type DeleteScheduledTaskArgs struct {
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	ScheduledTaskId      string
}
type DeleteScheduledTaskResponse struct {
	RequestId string
}

func (c *EssClient) DescribeScalingRules(req *DescribeScalingRulesArgs) (resp *DescribeScalingRulesResponse, err error) {
	resp = &DescribeScalingRulesResponse{}
	err = c.Request("DescribeScalingRules", req, resp)
	return
}

type DescribeScalingRulesScalingRule struct {
	ScalingRuleId   string
	ScalingGroupId  string
	ScalingRuleName string
	Cooldown        int
	AdjustmentType  string
	AdjustmentValue int
	MinSize         int
	MaxSize         int
	ScalingRuleAri  string
}
type DescribeScalingRulesArgs struct {
	ScalingRuleName1     string
	ResourceOwnerId      int64
	ScalingRuleName2     string
	ScalingRuleName3     string
	ScalingRuleName4     string
	ScalingRuleName5     string
	ScalingGroupId       string
	ScalingRuleName6     string
	ScalingRuleName7     string
	ScalingRuleName8     string
	ScalingRuleAri9      string
	ScalingRuleName9     string
	PageNumber           int
	PageSize             int
	ScalingRuleId10      string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	ScalingRuleAri1      string
	ScalingRuleAri2      string
	ScalingRuleName10    string
	ScalingRuleAri3      string
	ScalingRuleAri4      string
	ScalingRuleId8       string
	ScalingRuleAri5      string
	ScalingRuleId9       string
	ScalingRuleAri6      string
	ScalingRuleAri7      string
	ScalingRuleAri10     string
	ScalingRuleAri8      string
	ScalingRuleId4       string
	ScalingRuleId5       string
	ScalingRuleId6       string
	ScalingRuleId7       string
	ScalingRuleId1       string
	ScalingRuleId2       string
	ScalingRuleId3       string
}
type DescribeScalingRulesResponse struct {
	TotalCount   int
	PageNumber   int
	PageSize     int
	RequestId    string
	ScalingRules DescribeScalingRulesScalingRuleList
}

type DescribeScalingRulesScalingRuleList []DescribeScalingRulesScalingRule

func (list *DescribeScalingRulesScalingRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingRulesScalingRule)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

func (c *EssClient) DescribeRegions(req *DescribeRegionsArgs) (resp *DescribeRegionsResponse, err error) {
	resp = &DescribeRegionsResponse{}
	err = c.Request("DescribeRegions", req, resp)
	return
}

type DescribeRegionsRegion struct {
	RegionId           string
	ClassicUnavailable bool
	VpcUnavailable     bool
}
type DescribeRegionsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerId              int64
}
type DescribeRegionsResponse struct {
	Regions DescribeRegionsRegionList
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

func (c *EssClient) DescribeScalingGroups(req *DescribeScalingGroupsArgs) (resp *DescribeScalingGroupsResponse, err error) {
	resp = &DescribeScalingGroupsResponse{}
	err = c.Request("DescribeScalingGroups", req, resp)
	return
}

type DescribeScalingGroupsScalingGroup struct {
	DefaultCooldown              int
	MaxSize                      int
	PendingCapacity              int
	RemovingCapacity             int
	ScalingGroupName             string
	ActiveCapacity               int
	ActiveScalingConfigurationId string
	ScalingGroupId               string
	RegionId                     string
	TotalCapacity                int
	MinSize                      int
	LifecycleState               string
	CreationTime                 string
	VpcId                        string
	VSwitchId                    string
	VSwitchIds                   DescribeScalingGroupsVSwitchIdList
	RemovalPolicies              DescribeScalingGroupsRemovalPolicyList
	DBInstanceIds                DescribeScalingGroupsDBInstanceIdList
	LoadBalancerIds              DescribeScalingGroupsLoadBalancerIdList
}
type DescribeScalingGroupsArgs struct {
	ResourceOwnerId      int64
	ScalingGroupId10     string
	ScalingGroupId12     string
	ScalingGroupId13     string
	ScalingGroupId14     string
	ScalingGroupId15     string
	PageNumber           int
	PageSize             int
	ScalingGroupName20   string
	ScalingGroupName19   string
	ScalingGroupId20     string
	ScalingGroupName18   string
	ScalingGroupName17   string
	ScalingGroupName16   string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	ScalingGroupName1    string
	ScalingGroupName2    string
	ScalingGroupId2      string
	ScalingGroupId1      string
	ScalingGroupId6      string
	ScalingGroupId16     string
	ScalingGroupName7    string
	ScalingGroupName11   string
	ScalingGroupId5      string
	ScalingGroupId17     string
	ScalingGroupName8    string
	ScalingGroupName10   string
	ScalingGroupId4      string
	ScalingGroupId18     string
	ScalingGroupName9    string
	ScalingGroupId3      string
	ScalingGroupId19     string
	ScalingGroupName3    string
	ScalingGroupName15   string
	ScalingGroupId9      string
	ScalingGroupName4    string
	ScalingGroupName14   string
	ScalingGroupId8      string
	ScalingGroupName5    string
	ScalingGroupName13   string
	ScalingGroupId7      string
	ScalingGroupName6    string
	ScalingGroupName12   string
}
type DescribeScalingGroupsResponse struct {
	TotalCount    int
	PageNumber    int
	PageSize      int
	RequestId     string
	ScalingGroups DescribeScalingGroupsScalingGroupList
}

type DescribeScalingGroupsVSwitchIdList []string

func (list *DescribeScalingGroupsVSwitchIdList) UnmarshalJSON(data []byte) error {
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

type DescribeScalingGroupsRemovalPolicyList []string

func (list *DescribeScalingGroupsRemovalPolicyList) UnmarshalJSON(data []byte) error {
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

type DescribeScalingGroupsDBInstanceIdList []string

func (list *DescribeScalingGroupsDBInstanceIdList) UnmarshalJSON(data []byte) error {
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

type DescribeScalingGroupsLoadBalancerIdList []string

func (list *DescribeScalingGroupsLoadBalancerIdList) UnmarshalJSON(data []byte) error {
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

type DescribeScalingGroupsScalingGroupList []DescribeScalingGroupsScalingGroup

func (list *DescribeScalingGroupsScalingGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingGroupsScalingGroup)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

func (c *EssClient) CreateScheduledTask(req *CreateScheduledTaskArgs) (resp *CreateScheduledTaskResponse, err error) {
	resp = &CreateScheduledTaskResponse{}
	err = c.Request("CreateScheduledTask", req, resp)
	return
}

type CreateScheduledTaskArgs struct {
	LaunchTime           string
	ScheduledAction      string
	ResourceOwnerAccount string
	OwnerAccount         string
	Description          string
	OwnerId              int64
	RecurrenceValue      string
	LaunchExpirationTime int
	RecurrenceEndTime    string
	ScheduledTaskName    string
	TaskEnabled          bool
	RecurrenceType       string
}
type CreateScheduledTaskResponse struct {
	ScheduledTaskId string
	RequestId       string
}

func (c *EssClient) EnableScalingGroup(req *EnableScalingGroupArgs) (resp *EnableScalingGroupResponse, err error) {
	resp = &EnableScalingGroupResponse{}
	err = c.Request("EnableScalingGroup", req, resp)
	return
}

type EnableScalingGroupArgs struct {
	InstanceId10                 string
	LoadBalancerWeight6          int
	LoadBalancerWeight11         int
	LoadBalancerWeight7          int
	LoadBalancerWeight12         int
	ResourceOwnerId              int64
	InstanceId12                 string
	LoadBalancerWeight8          int
	InstanceId11                 string
	LoadBalancerWeight9          int
	LoadBalancerWeight10         int
	LoadBalancerWeight2          int
	LoadBalancerWeight15         int
	LoadBalancerWeight3          int
	LoadBalancerWeight16         int
	ScalingGroupId               string
	LoadBalancerWeight4          int
	LoadBalancerWeight13         int
	LoadBalancerWeight5          int
	LoadBalancerWeight14         int
	ActiveScalingConfigurationId string
	LoadBalancerWeight1          int
	InstanceId20                 string
	InstanceId1                  string
	LoadBalancerWeight20         int
	InstanceId3                  string
	ResourceOwnerAccount         string
	InstanceId2                  string
	InstanceId5                  string
	InstanceId4                  string
	OwnerAccount                 string
	InstanceId7                  string
	InstanceId6                  string
	InstanceId9                  string
	InstanceId8                  string
	OwnerId                      int64
	InstanceId18                 string
	LoadBalancerWeight19         int
	InstanceId17                 string
	LoadBalancerWeight17         int
	InstanceId19                 string
	LoadBalancerWeight18         int
	InstanceId14                 string
	InstanceId13                 string
	InstanceId16                 string
	InstanceId15                 string
}
type EnableScalingGroupResponse struct {
	RequestId string
}

func (c *EssClient) CreateScalingGroup(req *CreateScalingGroupArgs) (resp *CreateScalingGroupResponse, err error) {
	resp = &CreateScalingGroupResponse{}
	err = c.Request("CreateScalingGroup", req, resp)
	return
}

type CreateScalingGroupArgs struct {
	DBInstanceIds        string
	LoadBalancerIds      string
	ResourceOwnerAccount string
	ScalingGroupName     string
	VSwitchIdss          CreateScalingGroupVSwitchIdsList
	OwnerAccount         string
	MinSize              int
	OwnerId              int64
	VSwitchId            string
	MaxSize              int
	DefaultCooldown      int
	RemovalPolicy1       string
	RemovalPolicy2       string
}
type CreateScalingGroupResponse struct {
	ScalingGroupId string
	RequestId      string
}

type CreateScalingGroupVSwitchIdsList []string

func (list *CreateScalingGroupVSwitchIdsList) UnmarshalJSON(data []byte) error {
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
