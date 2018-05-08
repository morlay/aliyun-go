package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeExecutionPlanRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *DescribeExecutionPlanRequest) Invoke(client *sdk.Client) (resp *DescribeExecutionPlanResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeExecutionPlan", "", "")
	resp = &DescribeExecutionPlanResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeExecutionPlanResponse struct {
	responses.BaseResponse
	RequestId             common.String
	Id                    common.String
	Name                  common.String
	Status                common.String
	Strategy              common.String
	TimeInterval          common.Integer
	StartTime             common.Long
	TimeUnit              common.String
	DayOfWeek             common.String
	DayOfMonth            common.String
	CreateClusterOnDemand bool
	ClusterId             common.String
	ClusterName           common.String
	WorkflowApp           common.String
	ExecutionPlanVersion  common.Long
	JobInfoList           DescribeExecutionPlanJobInfoList
	ClusterInfo           DescribeExecutionPlanClusterInfo
}

type DescribeExecutionPlanJobInfo struct {
	Id           common.String
	Name         common.String
	Type         common.String
	RunParameter common.String
	FailAct      common.String
}

type DescribeExecutionPlanClusterInfo struct {
	Name                   common.String
	ZoneId                 common.String
	LogEnable              bool
	LogPath                common.String
	SecurityGroupId        common.String
	EmrVer                 common.String
	ClusterType            common.String
	HighAvailabilityEnable bool
	VpcId                  common.String
	VSwitchId              common.String
	NetType                common.String
	IoOptimized            bool
	InstanceGeneration     common.String
	Configurations         common.String
	EasEnable              bool
	UserDefinedEmrEcsRole  common.String
	EcsOrders              DescribeExecutionPlanEcsOrderInfoList
	BootstrapActionList    DescribeExecutionPlanBootstrapActionList
	SoftwareInfo           DescribeExecutionPlanSoftwareInfo
}

type DescribeExecutionPlanEcsOrderInfo struct {
	Index        common.Integer
	NodeCount    common.Integer
	InstanceType common.String
	DiskType     common.String
	DiskCapacity common.Integer
	NodeType     common.String
	DiskCount    common.Integer
}

type DescribeExecutionPlanBootstrapAction struct {
	Name common.String
	Path common.String
	Arg  common.String
}

type DescribeExecutionPlanSoftwareInfo struct {
	EmrVer      common.String
	ClusterType common.String
	Softwares   DescribeExecutionPlanSoftwareList
}

type DescribeExecutionPlanSoftware struct {
	DisplayName common.String
	Name        common.String
	OnlyDisplay bool
	StartTpe    common.Integer
	Version     common.String
	Optional    bool
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
