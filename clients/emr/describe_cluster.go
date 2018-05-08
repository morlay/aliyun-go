package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId   common.String
	ClusterInfo DescribeClusterClusterInfo
}

type DescribeClusterClusterInfo struct {
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
	EcsOrderInfoList       DescribeClusterEcsOrderInfoList
	BootstrapActionList    DescribeClusterBootstrapActionList
	FailReason             DescribeClusterFailReason
	SoftwareInfo           DescribeClusterSoftwareInfo
}

type DescribeClusterEcsOrderInfo struct {
	NodeType       common.String
	InstanceType   common.String
	CpuCore        common.Integer
	MemoryCapacity common.Integer
	DiskType       common.String
	DiskCapacity   common.Integer
	DiskCount      common.Integer
	BandWidth      common.String
	Nodes          DescribeClusterNodeList
}

type DescribeClusterNode struct {
	ZoneId         common.String
	InstanceId     common.String
	Status         common.String
	PubIp          common.String
	InnerIp        common.String
	ExpiredTime    common.String
	EmrExpiredTime common.String
	DaemonInfos    DescribeClusterDaemonInfoList
	DiskInfos      DescribeClusterDiskInfoList
}

type DescribeClusterDaemonInfo struct {
	Name common.String
}

type DescribeClusterDiskInfo struct {
	Device   common.String
	DiskName common.String
	DiskId   common.String
	Type     common.String
	Size     common.Integer
}

type DescribeClusterBootstrapAction struct {
	Name common.String
	Path common.String
	Arg  common.String
}

type DescribeClusterFailReason struct {
	ErrorCode common.String
	ErrorMsg  common.String
	RequestId common.String
}

type DescribeClusterSoftwareInfo struct {
	EmrVer      common.String
	ClusterType common.String
	Softwares   DescribeClusterSoftwareList
}

type DescribeClusterSoftware struct {
	DisplayName common.String
	Name        common.String
	OnlyDisplay bool
	StartTpe    common.Integer
	Version     common.String
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
