package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateExecutionPlanRequest struct {
	requests.RpcRequest
	ResourceOwnerId        int64                                      `position:"Query" name:"ResourceOwnerId"`
	TimeInterval           int                                        `position:"Query" name:"TimeInterval"`
	LogPath                string                                     `position:"Query" name:"LogPath"`
	ClusterName            string                                     `position:"Query" name:"ClusterName"`
	Configurations         string                                     `position:"Query" name:"Configurations"`
	IoOptimized            string                                     `position:"Query" name:"IoOptimized"`
	SecurityGroupId        string                                     `position:"Query" name:"SecurityGroupId"`
	EasEnable              string                                     `position:"Query" name:"EasEnable"`
	CreateClusterOnDemand  string                                     `position:"Query" name:"CreateClusterOnDemand"`
	StartTime              int64                                      `position:"Query" name:"StartTime"`
	JobIdLists             *CreateExecutionPlanJobIdListList          `position:"Query" type:"Repeated" name:"JobIdList"`
	DayOfMonth             string                                     `position:"Query" name:"DayOfMonth"`
	BootstrapActions       *CreateExecutionPlanBootstrapActionList    `position:"Query" type:"Repeated" name:"BootstrapAction"`
	UseLocalMetaDb         string                                     `position:"Query" name:"UseLocalMetaDb"`
	EmrVer                 string                                     `position:"Query" name:"EmrVer"`
	UserDefinedEmrEcsRole  string                                     `position:"Query" name:"UserDefinedEmrEcsRole"`
	IsOpenPublicIp         string                                     `position:"Query" name:"IsOpenPublicIp"`
	ClusterId              string                                     `position:"Query" name:"ClusterId"`
	TimeUnit               string                                     `position:"Query" name:"TimeUnit"`
	InstanceGeneration     string                                     `position:"Query" name:"InstanceGeneration"`
	ClusterType            string                                     `position:"Query" name:"ClusterType"`
	VSwitchId              string                                     `position:"Query" name:"VSwitchId"`
	OptionSoftWareLists    *CreateExecutionPlanOptionSoftWareListList `position:"Query" type:"Repeated" name:"OptionSoftWareList"`
	VpcId                  string                                     `position:"Query" name:"VpcId"`
	NetType                string                                     `position:"Query" name:"NetType"`
	EcsOrders              *CreateExecutionPlanEcsOrderList           `position:"Query" type:"Repeated" name:"EcsOrder"`
	WorkflowDefinition     string                                     `position:"Query" name:"WorkflowDefinition"`
	Name                   string                                     `position:"Query" name:"Name"`
	DayOfWeek              string                                     `position:"Query" name:"DayOfWeek"`
	ZoneId                 string                                     `position:"Query" name:"ZoneId"`
	Strategy               string                                     `position:"Query" name:"Strategy"`
	HighAvailabilityEnable string                                     `position:"Query" name:"HighAvailabilityEnable"`
	LogEnable              string                                     `position:"Query" name:"LogEnable"`
}

func (req *CreateExecutionPlanRequest) Invoke(client *sdk.Client) (resp *CreateExecutionPlanResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateExecutionPlan", "", "")
	resp = &CreateExecutionPlanResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateExecutionPlanBootstrapAction struct {
	Name string `name:"Name"`
	Path string `name:"Path"`
	Arg  string `name:"Arg"`
}

type CreateExecutionPlanEcsOrder struct {
	Index        int    `name:"Index"`
	NodeCount    int    `name:"NodeCount"`
	NodeType     string `name:"NodeType"`
	InstanceType string `name:"InstanceType"`
	DiskType     string `name:"DiskType"`
	DiskCapacity int    `name:"DiskCapacity"`
	DiskCount    int    `name:"DiskCount"`
}

type CreateExecutionPlanResponse struct {
	responses.BaseResponse
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
