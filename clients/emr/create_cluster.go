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
	Name                   string                               `position:"Query" name:"Name"`
	ZoneId                 string                               `position:"Query" name:"ZoneId"`
	LogEnable              string                               `position:"Query" name:"LogEnable"`
	LogPath                string                               `position:"Query" name:"LogPath"`
	SecurityGroupId        string                               `position:"Query" name:"SecurityGroupId"`
	IsOpenPublicIp         string                               `position:"Query" name:"IsOpenPublicIp"`
	SecurityGroupName      string                               `position:"Query" name:"SecurityGroupName"`
	ChargeType             string                               `position:"Query" name:"ChargeType"`
	Period                 int                                  `position:"Query" name:"Period"`
	AutoRenew              string                               `position:"Query" name:"AutoRenew"`
	AutoRenewPeriod        int                                  `position:"Query" name:"AutoRenewPeriod"`
	VpcId                  string                               `position:"Query" name:"VpcId"`
	VSwitchId              string                               `position:"Query" name:"VSwitchId"`
	NetType                string                               `position:"Query" name:"NetType"`
	EmrVer                 string                               `position:"Query" name:"EmrVer"`
	ClusterType            string                               `position:"Query" name:"ClusterType"`
	HighAvailabilityEnable string                               `position:"Query" name:"HighAvailabilityEnable"`
	IoOptimized            string                               `position:"Query" name:"IoOptimized"`
	InstanceGeneration     string                               `position:"Query" name:"InstanceGeneration"`
	MasterPwdEnable        string                               `position:"Query" name:"MasterPwdEnable"`
	MasterPwd              string                               `position:"Query" name:"MasterPwd"`
	Configurations         string                               `position:"Query" name:"Configurations"`
	OptionSoftWareLists    *CreateClusterOptionSoftWareListList `position:"Query" type:"Repeated" name:"OptionSoftWareList"`
	EcsOrders              *CreateClusterEcsOrderList           `position:"Query" type:"Repeated" name:"EcsOrder"`
	BootstrapActions       *CreateClusterBootstrapActionList    `position:"Query" type:"Repeated" name:"BootstrapAction"`
}

func (req *CreateClusterRequest) Invoke(client *sdk.Client) (resp *CreateClusterResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateCluster", "", "")
	resp = &CreateClusterResponse{}
	err = client.DoAction(req, resp)
	return
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

type CreateClusterBootstrapAction struct {
	Name string `name:"Name"`
	Path string `name:"Path"`
	Arg  string `name:"Arg"`
}

type CreateClusterResponse struct {
	responses.BaseResponse
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
