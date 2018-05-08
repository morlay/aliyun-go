package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeClusterForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *DescribeClusterForAdminRequest) Invoke(client *sdk.Client) (resp *DescribeClusterForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterForAdmin", "", "")
	resp = &DescribeClusterForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterForAdminResponse struct {
	responses.BaseResponse
	RequestId   common.String
	ClusterInfo DescribeClusterForAdminClusterInfo
}

type DescribeClusterForAdminClusterInfo struct {
	Id                     common.String
	RegionId               common.String
	ZoneId                 common.String
	Name                   common.String
	CreateType             common.String
	StartTime              common.Long
	StopTime               common.Long
	LogEnable              bool
	LogPath                common.String
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
	HostGroupList          DescribeClusterForAdminHostGroupList
	EcsOrderInfoList       DescribeClusterForAdminEcsOrderInfoList
	BootstrapActionList    DescribeClusterForAdminBootstrapActionList
	FailReason             DescribeClusterForAdminFailReason
	SoftwareInfo           DescribeClusterForAdminSoftwareInfo
}

type DescribeClusterForAdminHostGroup struct {
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
	Nodes          DescribeClusterForAdminNodeList
}

type DescribeClusterForAdminNode struct {
	ZoneId         common.String
	InstanceId     common.String
	Status         common.String
	PubIp          common.String
	InnerIp        common.String
	ExpiredTime    common.String
	EmrExpiredTime common.String
	DaemonInfos    DescribeClusterForAdminDaemonInfoList
	DiskInfos      DescribeClusterForAdminDiskInfoList
}

type DescribeClusterForAdminDaemonInfo struct {
	Name common.String
}

type DescribeClusterForAdminDiskInfo struct {
	Device   common.String
	DiskName common.String
	DiskId   common.String
	Type     common.String
	Size     common.Integer
}

type DescribeClusterForAdminEcsOrderInfo struct {
	NodeType       common.String
	InstanceType   common.String
	CpuCore        common.Integer
	MemoryCapacity common.Integer
	DiskType       common.String
	DiskCapacity   common.Integer
	DiskCount      common.Integer
	BandWidth      common.String
	Nodes1         DescribeClusterForAdminNode2List
}

type DescribeClusterForAdminNode2 struct {
	ZoneId         common.String
	InstanceId     common.String
	Status         common.String
	PubIp          common.String
	InnerIp        common.String
	ExpiredTime    common.String
	EmrExpiredTime common.String
	DaemonInfos3   DescribeClusterForAdminDaemonInfo5List
	DiskInfos4     DescribeClusterForAdminDiskInfo6List
}

type DescribeClusterForAdminDaemonInfo5 struct {
	Name common.String
}

type DescribeClusterForAdminDiskInfo6 struct {
	Device   common.String
	DiskName common.String
	DiskId   common.String
	Type     common.String
	Size     common.Integer
}

type DescribeClusterForAdminBootstrapAction struct {
	Name common.String
	Path common.String
	Arg  common.String
}

type DescribeClusterForAdminFailReason struct {
	ErrorCode common.String
	ErrorMsg  common.String
	RequestId common.String
}

type DescribeClusterForAdminSoftwareInfo struct {
	EmrVer      common.String
	ClusterType common.String
	Softwares   DescribeClusterForAdminSoftwareList
}

type DescribeClusterForAdminSoftware struct {
	DisplayName common.String
	Name        common.String
	OnlyDisplay bool
	StartTpe    common.Integer
	Version     common.String
}

type DescribeClusterForAdminHostGroupList []DescribeClusterForAdminHostGroup

func (list *DescribeClusterForAdminHostGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminHostGroup)
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

type DescribeClusterForAdminEcsOrderInfoList []DescribeClusterForAdminEcsOrderInfo

func (list *DescribeClusterForAdminEcsOrderInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminEcsOrderInfo)
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

type DescribeClusterForAdminBootstrapActionList []DescribeClusterForAdminBootstrapAction

func (list *DescribeClusterForAdminBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminBootstrapAction)
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

type DescribeClusterForAdminNodeList []DescribeClusterForAdminNode

func (list *DescribeClusterForAdminNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminNode)
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

type DescribeClusterForAdminDaemonInfoList []DescribeClusterForAdminDaemonInfo

func (list *DescribeClusterForAdminDaemonInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminDaemonInfo)
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

type DescribeClusterForAdminDiskInfoList []DescribeClusterForAdminDiskInfo

func (list *DescribeClusterForAdminDiskInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminDiskInfo)
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

type DescribeClusterForAdminNode2List []DescribeClusterForAdminNode2

func (list *DescribeClusterForAdminNode2List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminNode2)
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

type DescribeClusterForAdminDaemonInfo5List []DescribeClusterForAdminDaemonInfo5

func (list *DescribeClusterForAdminDaemonInfo5List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminDaemonInfo5)
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

type DescribeClusterForAdminDiskInfo6List []DescribeClusterForAdminDiskInfo6

func (list *DescribeClusterForAdminDiskInfo6List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminDiskInfo6)
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

type DescribeClusterForAdminSoftwareList []DescribeClusterForAdminSoftware

func (list *DescribeClusterForAdminSoftwareList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminSoftware)
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
