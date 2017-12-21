package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyExecutionPlanRequest struct {
	requests.RpcRequest
	ResourceOwnerId        int64                                      `position:"Query" name:"ResourceOwnerId"`
	ClusterName            string                                     `position:"Query" name:"ClusterName"`
	ClusterId              string                                     `position:"Query" name:"ClusterId"`
	ZoneId                 string                                     `position:"Query" name:"ZoneId"`
	LogEnable              string                                     `position:"Query" name:"LogEnable"`
	LogPath                string                                     `position:"Query" name:"LogPath"`
	SecurityGroupId        string                                     `position:"Query" name:"SecurityGroupId"`
	IsOpenPublicIp         string                                     `position:"Query" name:"IsOpenPublicIp"`
	CreateClusterOnDemand  string                                     `position:"Query" name:"CreateClusterOnDemand"`
	EmrVer                 string                                     `position:"Query" name:"EmrVer"`
	ClusterType            string                                     `position:"Query" name:"ClusterType"`
	HighAvailabilityEnable string                                     `position:"Query" name:"HighAvailabilityEnable"`
	VpcId                  string                                     `position:"Query" name:"VpcId"`
	VSwitchId              string                                     `position:"Query" name:"VSwitchId"`
	NetType                string                                     `position:"Query" name:"NetType"`
	IoOptimized            string                                     `position:"Query" name:"IoOptimized"`
	InstanceGeneration     string                                     `position:"Query" name:"InstanceGeneration"`
	Configurations         string                                     `position:"Query" name:"Configurations"`
	Id                     string                                     `position:"Query" name:"Id"`
	ExecutionPlanVersion   int64                                      `position:"Query" name:"ExecutionPlanVersion"`
	Name                   string                                     `position:"Query" name:"Name"`
	Strategy               string                                     `position:"Query" name:"Strategy"`
	TimeInterval           int                                        `position:"Query" name:"TimeInterval"`
	StartTime              int64                                      `position:"Query" name:"StartTime"`
	TimeUnit               string                                     `position:"Query" name:"TimeUnit"`
	OptionSoftWareLists    *ModifyExecutionPlanOptionSoftWareListList `position:"Query" type:"Repeated" name:"OptionSoftWareList"`
	EcsOrders              *ModifyExecutionPlanEcsOrderList           `position:"Query" type:"Repeated" name:"EcsOrder"`
	BootstrapActions       *ModifyExecutionPlanBootstrapActionList    `position:"Query" type:"Repeated" name:"BootstrapAction"`
	JobIdLists             *ModifyExecutionPlanJobIdListList          `position:"Query" type:"Repeated" name:"JobIdList"`
}

func (req *ModifyExecutionPlanRequest) Invoke(client *sdk.Client) (resp *ModifyExecutionPlanResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyExecutionPlan", "", "")
	resp = &ModifyExecutionPlanResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyExecutionPlanEcsOrder struct {
	Index        int    `name:"Index"`
	NodeCount    int    `name:"NodeCount"`
	InstanceType string `name:"InstanceType"`
	DiskType     string `name:"DiskType"`
	DiskCapacity int    `name:"DiskCapacity"`
	NodeType     string `name:"NodeType"`
	DiskCount    int    `name:"DiskCount"`
}

type ModifyExecutionPlanBootstrapAction struct {
	Name string `name:"Name"`
	Path string `name:"Path"`
	Arg  string `name:"Arg"`
}

type ModifyExecutionPlanResponse struct {
	responses.BaseResponse
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
