package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateExecutionPlanRequest struct {
	ResourceOwnerId        int64                                      `position:"Query" name:"ResourceOwnerId"`
	Name                   string                                     `position:"Query" name:"Name"`
	Strategy               string                                     `position:"Query" name:"Strategy"`
	TimeInterval           int                                        `position:"Query" name:"TimeInterval"`
	StartTime              int64                                      `position:"Query" name:"StartTime"`
	TimeUnit               string                                     `position:"Query" name:"TimeUnit"`
	ClusterId              string                                     `position:"Query" name:"ClusterId"`
	CreateClusterOnDemand  string                                     `position:"Query" name:"CreateClusterOnDemand"`
	ClusterName            string                                     `position:"Query" name:"ClusterName"`
	ZoneId                 string                                     `position:"Query" name:"ZoneId"`
	LogEnable              string                                     `position:"Query" name:"LogEnable"`
	LogPath                string                                     `position:"Query" name:"LogPath"`
	SecurityGroupId        string                                     `position:"Query" name:"SecurityGroupId"`
	IsOpenPublicIp         string                                     `position:"Query" name:"IsOpenPublicIp"`
	EmrVer                 string                                     `position:"Query" name:"EmrVer"`
	ClusterType            string                                     `position:"Query" name:"ClusterType"`
	HighAvailabilityEnable string                                     `position:"Query" name:"HighAvailabilityEnable"`
	VpcId                  string                                     `position:"Query" name:"VpcId"`
	VSwitchId              string                                     `position:"Query" name:"VSwitchId"`
	NetType                string                                     `position:"Query" name:"NetType"`
	IoOptimized            string                                     `position:"Query" name:"IoOptimized"`
	InstanceGeneration     string                                     `position:"Query" name:"InstanceGeneration"`
	Configurations         string                                     `position:"Query" name:"Configurations"`
	JobIdLists             *CreateExecutionPlanJobIdListList          `position:"Query" type:"Repeated" name:"JobIdList"`
	OptionSoftWareLists    *CreateExecutionPlanOptionSoftWareListList `position:"Query" type:"Repeated" name:"OptionSoftWareList"`
	EcsOrders              *CreateExecutionPlanEcsOrderList           `position:"Query" type:"Repeated" name:"EcsOrder"`
	BootstrapActions       *CreateExecutionPlanBootstrapActionList    `position:"Query" type:"Repeated" name:"BootstrapAction"`
}

func (r CreateExecutionPlanRequest) Invoke(client *sdk.Client) (response *CreateExecutionPlanResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateExecutionPlanRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateExecutionPlan", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateExecutionPlanResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateExecutionPlanResponse

	err = client.DoAction(&req, &resp)
	return
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

type CreateExecutionPlanBootstrapAction struct {
	Name string `name:"Name"`
	Path string `name:"Path"`
	Arg  string `name:"Arg"`
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
