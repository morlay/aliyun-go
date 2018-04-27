package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyClusterTemplateRequest struct {
	requests.RpcRequest
	ResourceOwnerId        int64                                        `position:"Query" name:"ResourceOwnerId"`
	LogPath                string                                       `position:"Query" name:"LogPath"`
	MasterPwd              string                                       `position:"Query" name:"MasterPwd"`
	Configurations         string                                       `position:"Query" name:"Configurations"`
	SecurityGroupId        string                                       `position:"Query" name:"SecurityGroupId"`
	SshEnable              string                                       `position:"Query" name:"SshEnable"`
	EasEnable              string                                       `position:"Query" name:"EasEnable"`
	SecurityGroupName      string                                       `position:"Query" name:"SecurityGroupName"`
	BootstrapActions       *ModifyClusterTemplateBootstrapActionList    `position:"Query" type:"Repeated" name:"BootstrapAction"`
	UseLocalMetaDb         string                                       `position:"Query" name:"UseLocalMetaDb"`
	EmrVer                 string                                       `position:"Query" name:"EmrVer"`
	TemplateName           string                                       `position:"Query" name:"TemplateName"`
	Id                     string                                       `position:"Query" name:"Id"`
	UserDefinedEmrEcsRole  string                                       `position:"Query" name:"UserDefinedEmrEcsRole"`
	IsOpenPublicIp         string                                       `position:"Query" name:"IsOpenPublicIp"`
	Period                 int                                          `position:"Query" name:"Period"`
	LIoOptimized           string                                       `position:"Query" name:"LIoOptimized"`
	InstanceGeneration     string                                       `position:"Query" name:"InstanceGeneration"`
	VSwitchId              string                                       `position:"Query" name:"VSwitchId"`
	ClusterType            string                                       `position:"Query" name:"ClusterType"`
	AutoRenew              string                                       `position:"Query" name:"AutoRenew"`
	OptionSoftWareLists    *ModifyClusterTemplateOptionSoftWareListList `position:"Query" type:"Repeated" name:"OptionSoftWareList"`
	VpcId                  string                                       `position:"Query" name:"VpcId"`
	NetType                string                                       `position:"Query" name:"NetType"`
	HostGroups             *ModifyClusterTemplateHostGroupList          `position:"Query" type:"Repeated" name:"HostGroup"`
	ZoneId                 string                                       `position:"Query" name:"ZoneId"`
	ChargeType             string                                       `position:"Query" name:"ChargeType"`
	HighAvailabilityEnable string                                       `position:"Query" name:"HighAvailabilityEnable"`
}

func (req *ModifyClusterTemplateRequest) Invoke(client *sdk.Client) (resp *ModifyClusterTemplateResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyClusterTemplate", "", "")
	resp = &ModifyClusterTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyClusterTemplateBootstrapAction struct {
	Name string `name:"Name"`
	Path string `name:"Path"`
	Arg  string `name:"Arg"`
}

type ModifyClusterTemplateHostGroup struct {
	HostGroupId     string `name:"HostGroupId"`
	HostGroupName   string `name:"HostGroupName"`
	HostGroupType   string `name:"HostGroupType"`
	ClusterId       string `name:"ClusterId"`
	Comment         string `name:"Comment"`
	CreateType      string `name:"CreateType"`
	ChargeType      string `name:"ChargeType"`
	Period          int    `name:"Period"`
	NodeCount       int    `name:"NodeCount"`
	InstanceType    string `name:"InstanceType"`
	DiskType        string `name:"DiskType"`
	DiskCapacity    int    `name:"DiskCapacity"`
	DiskCount       int    `name:"DiskCount"`
	SysDiskCapacity int    `name:"SysDiskCapacity"`
	AutoRenew       string `name:"AutoRenew"`
	VSwitchId       string `name:"VSwitchId"`
}

type ModifyClusterTemplateResponse struct {
	responses.BaseResponse
	RequestId         string
	Success           string
	ErrCode           string
	ErrMsg            string
	ClusterTemplateId string
}

type ModifyClusterTemplateBootstrapActionList []ModifyClusterTemplateBootstrapAction

func (list *ModifyClusterTemplateBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyClusterTemplateBootstrapAction)
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

type ModifyClusterTemplateOptionSoftWareListList []string

func (list *ModifyClusterTemplateOptionSoftWareListList) UnmarshalJSON(data []byte) error {
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

type ModifyClusterTemplateHostGroupList []ModifyClusterTemplateHostGroup

func (list *ModifyClusterTemplateHostGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyClusterTemplateHostGroup)
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
