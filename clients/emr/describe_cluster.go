package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *DescribeClusterRequest) Invoke(client *sdk.Client) (resp *DescribeClusterResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeCluster", "", "")
	resp = &DescribeClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterResponse struct {
	responses.BaseResponse
	RequestId   string
	ClusterInfo DescribeClusterClusterInfo
}

type DescribeClusterClusterInfo struct {
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
	VpcId                  string
	VSwitchId              string
	NetType                string
	IoOptimized            bool
	InstanceGeneration     string
	ImageId                string
	SecurityGroupId        string
	SecurityGroupName      string
	BootstrapFailed        bool
	Configurations         string
	EcsOrderInfoList       DescribeClusterEcsOrderInfoList
	BootstrapActionList    DescribeClusterBootstrapActionList
	FailReason             DescribeClusterFailReason
	SoftwareInfo           DescribeClusterSoftwareInfo
}

type DescribeClusterEcsOrderInfo struct {
	NodeType       string
	InstanceType   string
	CpuCore        int
	MemoryCapacity int
	DiskType       string
	DiskCapacity   int
	DiskCount      int
	BandWidth      string
	Nodes          DescribeClusterNodeList
}

type DescribeClusterNode struct {
	ZoneId         string
	InstanceId     string
	Status         string
	PubIp          string
	InnerIp        string
	ExpiredTime    string
	EmrExpiredTime string
	DaemonInfos    DescribeClusterDaemonInfoList
	DiskInfos      DescribeClusterDiskInfoList
}

type DescribeClusterDaemonInfo struct {
	Name string
}

type DescribeClusterDiskInfo struct {
	Device   string
	DiskName string
	DiskId   string
	Type     string
	Size     int
}

type DescribeClusterBootstrapAction struct {
	Name string
	Path string
	Arg  string
}

type DescribeClusterFailReason struct {
	ErrorCode string
	ErrorMsg  string
	RequestId string
}

type DescribeClusterSoftwareInfo struct {
	EmrVer      string
	ClusterType string
	Softwares   DescribeClusterSoftwareList
}

type DescribeClusterSoftware struct {
	DisplayName string
	Name        string
	OnlyDisplay bool
	StartTpe    int
	Version     string
}

type DescribeClusterEcsOrderInfoList []DescribeClusterEcsOrderInfo

func (list *DescribeClusterEcsOrderInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterEcsOrderInfo)
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

type DescribeClusterBootstrapActionList []DescribeClusterBootstrapAction

func (list *DescribeClusterBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterBootstrapAction)
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

type DescribeClusterNodeList []DescribeClusterNode

func (list *DescribeClusterNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterNode)
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

type DescribeClusterDaemonInfoList []DescribeClusterDaemonInfo

func (list *DescribeClusterDaemonInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterDaemonInfo)
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

type DescribeClusterDiskInfoList []DescribeClusterDiskInfo

func (list *DescribeClusterDiskInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterDiskInfo)
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

type DescribeClusterSoftwareList []DescribeClusterSoftware

func (list *DescribeClusterSoftwareList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterSoftware)
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
