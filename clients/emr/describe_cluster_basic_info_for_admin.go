package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeClusterBasicInfoForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *DescribeClusterBasicInfoForAdminRequest) Invoke(client *sdk.Client) (resp *DescribeClusterBasicInfoForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterBasicInfoForAdmin", "", "")
	resp = &DescribeClusterBasicInfoForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterBasicInfoForAdminResponse struct {
	responses.BaseResponse
	RequestId   common.String
	ClusterInfo DescribeClusterBasicInfoForAdminClusterInfo
}

type DescribeClusterBasicInfoForAdminClusterInfo struct {
	Id                     common.String
	BizId                  common.String
	RegionId               common.String
	ZoneId                 common.String
	Name                   common.String
	CreateType             common.String
	ClusterType            common.String
	StartTime              common.Long
	StopTime               common.Long
	LogEnable              bool
	LogPath                common.String
	UserId                 common.String
	Status                 common.String
	HighAvailabilityEnable bool
	PayType                common.String
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
	EcmClusterId           common.String
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
	EcsOrderInfoList       DescribeClusterBasicInfoForAdminEcsOrderInfoList
	BootstrapActionList    DescribeClusterBasicInfoForAdminBootstrapActionList
	FailReason             DescribeClusterBasicInfoForAdminFailReason
	SoftwareInfo           DescribeClusterBasicInfoForAdminSoftwareInfo
}

type DescribeClusterBasicInfoForAdminEcsOrderInfo struct {
	NodeType       common.String
	InstanceType   common.String
	CpuCore        common.Integer
	MemoryCapacity common.Integer
	DiskType       common.String
	DiskCapacity   common.Integer
	DiskCount      common.Integer
	BandWidth      common.String
	Nodes          DescribeClusterBasicInfoForAdminNodeList
}

type DescribeClusterBasicInfoForAdminNode struct {
	ZoneId         common.String
	InstanceId     common.String
	Status         common.String
	PubIp          common.String
	InnerIp        common.String
	ExpiredTime    common.String
	EmrExpiredTime common.String
	DiskInfos      DescribeClusterBasicInfoForAdminDiskInfoList
}

type DescribeClusterBasicInfoForAdminDiskInfo struct {
	Device   common.String
	DiskName common.String
	DiskId   common.String
	Type     common.String
	Size     common.Integer
}

type DescribeClusterBasicInfoForAdminBootstrapAction struct {
	Name common.String
	Path common.String
	Arg  common.String
}

type DescribeClusterBasicInfoForAdminFailReason struct {
	ErrorCode common.String
	ErrorMsg  common.String
	RequestId common.String
}

type DescribeClusterBasicInfoForAdminSoftwareInfo struct {
	EmrVer      common.String
	ClusterType common.String
	Softwares   DescribeClusterBasicInfoForAdminSoftwareList
}

type DescribeClusterBasicInfoForAdminSoftware struct {
	DisplayName common.String
	Name        common.String
	OnlyDisplay bool
	StartTpe    common.Integer
	Version     common.String
}

type DescribeClusterBasicInfoForAdminEcsOrderInfoList []DescribeClusterBasicInfoForAdminEcsOrderInfo

func (list *DescribeClusterBasicInfoForAdminEcsOrderInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterBasicInfoForAdminEcsOrderInfo)
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

type DescribeClusterBasicInfoForAdminBootstrapActionList []DescribeClusterBasicInfoForAdminBootstrapAction

func (list *DescribeClusterBasicInfoForAdminBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterBasicInfoForAdminBootstrapAction)
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

type DescribeClusterBasicInfoForAdminNodeList []DescribeClusterBasicInfoForAdminNode

func (list *DescribeClusterBasicInfoForAdminNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterBasicInfoForAdminNode)
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

type DescribeClusterBasicInfoForAdminDiskInfoList []DescribeClusterBasicInfoForAdminDiskInfo

func (list *DescribeClusterBasicInfoForAdminDiskInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterBasicInfoForAdminDiskInfo)
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

type DescribeClusterBasicInfoForAdminSoftwareList []DescribeClusterBasicInfoForAdminSoftware

func (list *DescribeClusterBasicInfoForAdminSoftwareList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterBasicInfoForAdminSoftware)
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
