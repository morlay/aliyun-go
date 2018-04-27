package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId   string
	ClusterInfo DescribeClusterBasicInfoForAdminClusterInfo
}

type DescribeClusterBasicInfoForAdminClusterInfo struct {
	Id                     string
	BizId                  string
	RegionId               string
	ZoneId                 string
	Name                   string
	CreateType             string
	ClusterType            string
	StartTime              int64
	StopTime               int64
	LogEnable              bool
	LogPath                string
	UserId                 string
	Status                 string
	HighAvailabilityEnable bool
	PayType                string
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
	EcmClusterId           string
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
	EcsOrderInfoList       DescribeClusterBasicInfoForAdminEcsOrderInfoList
	BootstrapActionList    DescribeClusterBasicInfoForAdminBootstrapActionList
	FailReason             DescribeClusterBasicInfoForAdminFailReason
	SoftwareInfo           DescribeClusterBasicInfoForAdminSoftwareInfo
}

type DescribeClusterBasicInfoForAdminEcsOrderInfo struct {
	NodeType       string
	InstanceType   string
	CpuCore        int
	MemoryCapacity int
	DiskType       string
	DiskCapacity   int
	DiskCount      int
	BandWidth      string
	Nodes          DescribeClusterBasicInfoForAdminNodeList
}

type DescribeClusterBasicInfoForAdminNode struct {
	ZoneId         string
	InstanceId     string
	Status         string
	PubIp          string
	InnerIp        string
	ExpiredTime    string
	EmrExpiredTime string
	DiskInfos      DescribeClusterBasicInfoForAdminDiskInfoList
}

type DescribeClusterBasicInfoForAdminDiskInfo struct {
	Device   string
	DiskName string
	DiskId   string
	Type     string
	Size     int
}

type DescribeClusterBasicInfoForAdminBootstrapAction struct {
	Name string
	Path string
	Arg  string
}

type DescribeClusterBasicInfoForAdminFailReason struct {
	ErrorCode string
	ErrorMsg  string
	RequestId string
}

type DescribeClusterBasicInfoForAdminSoftwareInfo struct {
	EmrVer      string
	ClusterType string
	Softwares   DescribeClusterBasicInfoForAdminSoftwareList
}

type DescribeClusterBasicInfoForAdminSoftware struct {
	DisplayName string
	Name        string
	OnlyDisplay bool
	StartTpe    int
	Version     string
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
