package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyExecutionPlanClusterInfoRequest struct {
	requests.RpcRequest
	ResourceOwnerId        int64                                                 `position:"Query" name:"ResourceOwnerId"`
	LogPath                string                                                `position:"Query" name:"LogPath"`
	ClusterName            string                                                `position:"Query" name:"ClusterName"`
	Configurations         string                                                `position:"Query" name:"Configurations"`
	IoOptimized            string                                                `position:"Query" name:"IoOptimized"`
	SecurityGroupId        string                                                `position:"Query" name:"SecurityGroupId"`
	EasEnable              string                                                `position:"Query" name:"EasEnable"`
	CreateClusterOnDemand  string                                                `position:"Query" name:"CreateClusterOnDemand"`
	BootstrapActions       *ModifyExecutionPlanClusterInfoBootstrapActionList    `position:"Query" type:"Repeated" name:"BootstrapAction"`
	UseLocalMetaDb         string                                                `position:"Query" name:"UseLocalMetaDb"`
	EmrVer                 string                                                `position:"Query" name:"EmrVer"`
	Id                     string                                                `position:"Query" name:"Id"`
	IsOpenPublicIp         string                                                `position:"Query" name:"IsOpenPublicIp"`
	ClusterId              string                                                `position:"Query" name:"ClusterId"`
	InstanceGeneration     string                                                `position:"Query" name:"InstanceGeneration"`
	ClusterType            string                                                `position:"Query" name:"ClusterType"`
	VSwitchId              string                                                `position:"Query" name:"VSwitchId"`
	OptionSoftWareLists    *ModifyExecutionPlanClusterInfoOptionSoftWareListList `position:"Query" type:"Repeated" name:"OptionSoftWareList"`
	VpcId                  string                                                `position:"Query" name:"VpcId"`
	NetType                string                                                `position:"Query" name:"NetType"`
	EcsOrders              *ModifyExecutionPlanClusterInfoEcsOrderList           `position:"Query" type:"Repeated" name:"EcsOrder"`
	ZoneId                 string                                                `position:"Query" name:"ZoneId"`
	HighAvailabilityEnable string                                                `position:"Query" name:"HighAvailabilityEnable"`
	LogEnable              string                                                `position:"Query" name:"LogEnable"`
}

func (req *ModifyExecutionPlanClusterInfoRequest) Invoke(client *sdk.Client) (resp *ModifyExecutionPlanClusterInfoResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyExecutionPlanClusterInfo", "", "")
	resp = &ModifyExecutionPlanClusterInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyExecutionPlanClusterInfoBootstrapAction struct {
	Name string `name:"Name"`
	Path string `name:"Path"`
	Arg  string `name:"Arg"`
}

type ModifyExecutionPlanClusterInfoEcsOrder struct {
	Index        int    `name:"Index"`
	NodeCount    int    `name:"NodeCount"`
	InstanceType string `name:"InstanceType"`
	DiskType     string `name:"DiskType"`
	DiskCapacity int    `name:"DiskCapacity"`
	NodeType     string `name:"NodeType"`
	DiskCount    int    `name:"DiskCount"`
}

type ModifyExecutionPlanClusterInfoResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type ModifyExecutionPlanClusterInfoBootstrapActionList []ModifyExecutionPlanClusterInfoBootstrapAction

func (list *ModifyExecutionPlanClusterInfoBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyExecutionPlanClusterInfoBootstrapAction)
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

type ModifyExecutionPlanClusterInfoOptionSoftWareListList []string

func (list *ModifyExecutionPlanClusterInfoOptionSoftWareListList) UnmarshalJSON(data []byte) error {
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

type ModifyExecutionPlanClusterInfoEcsOrderList []ModifyExecutionPlanClusterInfoEcsOrder

func (list *ModifyExecutionPlanClusterInfoEcsOrderList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyExecutionPlanClusterInfoEcsOrder)
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
