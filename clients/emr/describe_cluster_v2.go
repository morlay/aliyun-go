package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId   string
	ClusterInfo DescribeClusterV2ClusterInfo
}

type DescribeClusterV2ClusterInfo struct {
	Id                     string
	RegionId               string
	ZoneId                 string
	Name                   string
	CreateType             string
	StartTime              int64
	StopTime               int64
	LogEnable              bool
	LogPath                string
	UserId                 string
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
	EasEnable              bool
	HostGroupList          DescribeClusterV2HostGroupList
	BootstrapActionList    DescribeClusterV2BootstrapActionList
	FailReason             DescribeClusterV2FailReason
	SoftwareInfo           DescribeClusterV2SoftwareInfo
}

type DescribeClusterV2HostGroup struct {
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
	Nodes          DescribeClusterV2NodeList
}

type DescribeClusterV2Node struct {
	ZoneId         string
	InstanceId     string
	Status         string
	PubIp          string
	InnerIp        string
	ExpiredTime    string
	CreateTime     string
	EmrExpiredTime string
	DaemonInfos    DescribeClusterV2DaemonInfoList
	DiskInfos      DescribeClusterV2DiskInfoList
}

type DescribeClusterV2DaemonInfo struct {
	Name string
}

type DescribeClusterV2DiskInfo struct {
	Device   string
	DiskName string
	DiskId   string
	Type     string
	Size     int
}

type DescribeClusterV2BootstrapAction struct {
	Name string
	Path string
	Arg  string
}

type DescribeClusterV2FailReason struct {
	ErrorCode string
	ErrorMsg  string
	RequestId string
}

type DescribeClusterV2SoftwareInfo struct {
	EmrVer      string
	ClusterType string
	Softwares   DescribeClusterV2SoftwareList
}

type DescribeClusterV2Software struct {
	DisplayName string
	Name        string
	OnlyDisplay bool
	StartTpe    int
	Version     string
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
