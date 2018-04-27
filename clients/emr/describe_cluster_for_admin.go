package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId   string
	ClusterInfo DescribeClusterForAdminClusterInfo
}

type DescribeClusterForAdminClusterInfo struct {
	Id                     string
	RegionId               string
	ZoneId                 string
	Name                   string
	CreateType             string
	StartTime              int64
	StopTime               int64
	LogEnable              bool
	LogPath                string
	Status                 string
	HighAvailabilityEnable bool
	ChargeType             string
	ExpiredTime            int64
	Period                 int
	RunningTime            int
	MasterNodeTotal        int
	MasterNodeInService    int
	CoreNodeTotal          int
	CoreNodeInService      int
	TaskNodeTotal          int
	TaskNodeInService      int
	ShowSoftwareInterface  bool
	CreateResource         string
	VpcId                  string
	VSwitchId              string
	NetType                string
	UserDefinedEmrEcsRole  string
	IoOptimized            bool
	InstanceGeneration     string
	ImageId                string
	SecurityGroupId        string
	SecurityGroupName      string
	BootstrapFailed        bool
	Configurations         string
	HostGroupList          DescribeClusterForAdminHostGroupList
	EcsOrderInfoList       DescribeClusterForAdminEcsOrderInfoList
	BootstrapActionList    DescribeClusterForAdminBootstrapActionList
	FailReason             DescribeClusterForAdminFailReason
	SoftwareInfo           DescribeClusterForAdminSoftwareInfo
}

type DescribeClusterForAdminHostGroup struct {
	HostGroupId    string
	HostGroupName  string
	HostGroupType  string
	ChargeType     string
	Period         string
	NodeCount      int
	InstanceType   string
	CpuCore        int
	MemoryCapacity int
	DiskType       string
	DiskCapacity   int
	DiskCount      int
	BandWidth      string
	Nodes          DescribeClusterForAdminNodeList
}

type DescribeClusterForAdminNode struct {
	ZoneId         string
	InstanceId     string
	Status         string
	PubIp          string
	InnerIp        string
	ExpiredTime    string
	EmrExpiredTime string
	DaemonInfos    DescribeClusterForAdminDaemonInfoList
	DiskInfos      DescribeClusterForAdminDiskInfoList
}

type DescribeClusterForAdminDaemonInfo struct {
	Name string
}

type DescribeClusterForAdminDiskInfo struct {
	Device   string
	DiskName string
	DiskId   string
	Type     string
	Size     int
}

type DescribeClusterForAdminEcsOrderInfo struct {
	NodeType       string
	InstanceType   string
	CpuCore        int
	MemoryCapacity int
	DiskType       string
	DiskCapacity   int
	DiskCount      int
	BandWidth      string
	Nodes1         DescribeClusterForAdminNode2List
}

type DescribeClusterForAdminNode2 struct {
	ZoneId         string
	InstanceId     string
	Status         string
	PubIp          string
	InnerIp        string
	ExpiredTime    string
	EmrExpiredTime string
	DaemonInfos3   DescribeClusterForAdminDaemonInfo5List
	DiskInfos4     DescribeClusterForAdminDiskInfo6List
}

type DescribeClusterForAdminDaemonInfo5 struct {
	Name string
}

type DescribeClusterForAdminDiskInfo6 struct {
	Device   string
	DiskName string
	DiskId   string
	Type     string
	Size     int
}

type DescribeClusterForAdminBootstrapAction struct {
	Name string
	Path string
	Arg  string
}

type DescribeClusterForAdminFailReason struct {
	ErrorCode string
	ErrorMsg  string
	RequestId string
}

type DescribeClusterForAdminSoftwareInfo struct {
	EmrVer      string
	ClusterType string
	Softwares   DescribeClusterForAdminSoftwareList
}

type DescribeClusterForAdminSoftware struct {
	DisplayName string
	Name        string
	OnlyDisplay bool
	StartTpe    int
	Version     string
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
