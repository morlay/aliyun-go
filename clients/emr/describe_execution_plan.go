package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId             string
	Id                    string
	Name                  string
	Status                string
	Strategy              string
	TimeInterval          int
	StartTime             int64
	TimeUnit              string
	DayOfWeek             string
	DayOfMonth            string
	CreateClusterOnDemand bool
	ClusterId             string
	ClusterName           string
	WorkflowApp           string
	ExecutionPlanVersion  int64
	JobInfoList           DescribeExecutionPlanJobInfoList
	ClusterInfo           DescribeExecutionPlanClusterInfo
}

type DescribeExecutionPlanJobInfo struct {
	Id           string
	Name         string
	Type         string
	RunParameter string
	FailAct      string
}

type DescribeExecutionPlanClusterInfo struct {
	Name                   string
	ZoneId                 string
	LogEnable              bool
	LogPath                string
	SecurityGroupId        string
	EmrVer                 string
	ClusterType            string
	HighAvailabilityEnable bool
	VpcId                  string
	VSwitchId              string
	NetType                string
	IoOptimized            bool
	InstanceGeneration     string
	Configurations         string
	EasEnable              bool
	UserDefinedEmrEcsRole  string
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
	OnlyDisplay bool
	StartTpe    int
	Version     string
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
