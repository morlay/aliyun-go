package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateClusterRequest struct {
	requests.RpcRequest
	ResourceOwnerId        int64                                `position:"Query" name:"ResourceOwnerId"`
	LogPath                string                               `position:"Query" name:"LogPath"`
	MasterPwd              string                               `position:"Query" name:"MasterPwd"`
	Configurations         string                               `position:"Query" name:"Configurations"`
	IoOptimized            string                               `position:"Query" name:"IoOptimized"`
	SecurityGroupId        string                               `position:"Query" name:"SecurityGroupId"`
	SshEnable              string                               `position:"Query" name:"SshEnable"`
	EasEnable              string                               `position:"Query" name:"EasEnable"`
	SecurityGroupName      string                               `position:"Query" name:"SecurityGroupName"`
	DepositType            string                               `position:"Query" name:"DepositType"`
	MachineType            string                               `position:"Query" name:"MachineType"`
	BootstrapActions       *CreateClusterBootstrapActionList    `position:"Query" type:"Repeated" name:"BootstrapAction"`
	UseLocalMetaDb         string                               `position:"Query" name:"UseLocalMetaDb"`
	EmrVer                 string                               `position:"Query" name:"EmrVer"`
	UserDefinedEmrEcsRole  string                               `position:"Query" name:"UserDefinedEmrEcsRole"`
	IsOpenPublicIp         string                               `position:"Query" name:"IsOpenPublicIp"`
	Period                 int                                  `position:"Query" name:"Period"`
	RelatedClusterId       string                               `position:"Query" name:"RelatedClusterId"`
	InstanceGeneration     string                               `position:"Query" name:"InstanceGeneration"`
	VSwitchId              string                               `position:"Query" name:"VSwitchId"`
	ClusterType            string                               `position:"Query" name:"ClusterType"`
	AutoRenew              string                               `position:"Query" name:"AutoRenew"`
	OptionSoftWareLists    *CreateClusterOptionSoftWareListList `position:"Query" type:"Repeated" name:"OptionSoftWareList"`
	VpcId                  string                               `position:"Query" name:"VpcId"`
	NetType                string                               `position:"Query" name:"NetType"`
	EcsOrders              *CreateClusterEcsOrderList           `position:"Query" type:"Repeated" name:"EcsOrder"`
	Name                   string                               `position:"Query" name:"Name"`
	ZoneId                 string                               `position:"Query" name:"ZoneId"`
	ChargeType             string                               `position:"Query" name:"ChargeType"`
	HighAvailabilityEnable string                               `position:"Query" name:"HighAvailabilityEnable"`
	MasterPwdEnable        string                               `position:"Query" name:"MasterPwdEnable"`
	LogEnable              string                               `position:"Query" name:"LogEnable"`
}

func (req *CreateClusterRequest) Invoke(client *sdk.Client) (resp *CreateClusterResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateCluster", "", "")
	resp = &CreateClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateClusterBootstrapAction struct {
	Name string `name:"Name"`
	Path string `name:"Path"`
	Arg  string `name:"Arg"`
}

type CreateClusterEcsOrder struct {
	Index        int    `name:"Index"`
	NodeCount    int    `name:"NodeCount"`
	NodeType     string `name:"NodeType"`
	InstanceType string `name:"InstanceType"`
	DiskType     string `name:"DiskType"`
	DiskCapacity int    `name:"DiskCapacity"`
	DiskCount    int    `name:"DiskCount"`
}

type CreateClusterResponse struct {
	responses.BaseResponse
	RequestId     string
	ClusterId     string
	EmrOrderId    string
	MasterOrderId string
	CoreOrderId   string
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
