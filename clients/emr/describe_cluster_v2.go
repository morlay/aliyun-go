package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeClusterV2Request struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *DescribeClusterV2Request) Invoke(client *sdk.Client) (resp *DescribeClusterV2Response, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterV2", "", "")
	resp = &DescribeClusterV2Response{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterV2Response struct {
	responses.BaseResponse
	RequestId   common.String
	ClusterInfo DescribeClusterV2ClusterInfo
}

type DescribeClusterV2ClusterInfo struct {
	Id                     common.String
	RegionId               common.String
	ZoneId                 common.String
	Name                   common.String
	CreateType             common.String
	StartTime              common.Long
	StopTime               common.Long
	LogEnable              bool
	LogPath                common.String
	UserId                 common.String
	Status                 common.String
	HighAvailabilityEnable bool
	ChargeType             common.String
	ExpiredTime            common.Long
	Period                 common.Integer
	RunningTime            common.Integer
	MasterNodeTotal        common.Integer
	MasterNodeInService    common.Integer
	CoreNodeTotal          common.Integer
	CoreNodeInService      common.Integer
	TaskNodeTotal          common.Integer
	TaskNodeInService      common.Integer
	ShowSoftwareInterface  bool
	CreateResource         common.String
	VpcId                  common.String
	VSwitchId              common.String
	NetType                common.String
	UserDefinedEmrEcsRole  common.String
	IoOptimized            bool
	InstanceGeneration     common.String
	ImageId                common.String
	SecurityGroupId        common.String
	SecurityGroupName      common.String
	BootstrapFailed        bool
	Configurations         common.String
	EasEnable              bool
	HostGroupList          DescribeClusterV2HostGroupList
	BootstrapActionList    DescribeClusterV2BootstrapActionList
	FailReason             DescribeClusterV2FailReason
	SoftwareInfo           DescribeClusterV2SoftwareInfo
}

type DescribeClusterV2HostGroup struct {
	HostGroupId    common.String
	HostGroupName  common.String
	HostGroupType  common.String
	ChargeType     common.String
	Period         common.String
	NodeCount      common.Integer
	InstanceType   common.String
	CpuCore        common.Integer
	MemoryCapacity common.Integer
	DiskType       common.String
	DiskCapacity   common.Integer
	DiskCount      common.Integer
	BandWidth      common.String
	Nodes          DescribeClusterV2NodeList
}

type DescribeClusterV2Node struct {
	ZoneId         common.String
	InstanceId     common.String
	Status         common.String
	PubIp          common.String
	InnerIp        common.String
	ExpiredTime    common.String
	CreateTime     common.String
	EmrExpiredTime common.String
	DaemonInfos    DescribeClusterV2DaemonInfoList
	DiskInfos      DescribeClusterV2DiskInfoList
}

type DescribeClusterV2DaemonInfo struct {
	Name common.String
}

type DescribeClusterV2DiskInfo struct {
	Device   common.String
	DiskName common.String
	DiskId   common.String
	Type     common.String
	Size     common.Integer
}

type DescribeClusterV2BootstrapAction struct {
	Name common.String
	Path common.String
	Arg  common.String
}

type DescribeClusterV2FailReason struct {
	ErrorCode common.String
	ErrorMsg  common.String
	RequestId common.String
}

type DescribeClusterV2SoftwareInfo struct {
	EmrVer      common.String
	ClusterType common.String
	Softwares   DescribeClusterV2SoftwareList
}

type DescribeClusterV2Software struct {
	DisplayName common.String
	Name        common.String
	OnlyDisplay bool
	StartTpe    common.Integer
	Version     common.String
}

type DescribeClusterV2HostGroupList []DescribeClusterV2HostGroup

func (list *DescribeClusterV2HostGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterV2HostGroup)
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

type DescribeClusterV2BootstrapActionList []DescribeClusterV2BootstrapAction

func (list *DescribeClusterV2BootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterV2BootstrapAction)
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

type DescribeClusterV2NodeList []DescribeClusterV2Node

func (list *DescribeClusterV2NodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterV2Node)
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

type DescribeClusterV2DaemonInfoList []DescribeClusterV2DaemonInfo

func (list *DescribeClusterV2DaemonInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterV2DaemonInfo)
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

type DescribeClusterV2DiskInfoList []DescribeClusterV2DiskInfo

func (list *DescribeClusterV2DiskInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterV2DiskInfo)
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

type DescribeClusterV2SoftwareList []DescribeClusterV2Software

func (list *DescribeClusterV2SoftwareList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterV2Software)
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
