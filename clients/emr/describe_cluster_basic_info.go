package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterBasicInfoRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *DescribeClusterBasicInfoRequest) Invoke(client *sdk.Client) (resp *DescribeClusterBasicInfoResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterBasicInfo", "", "")
	resp = &DescribeClusterBasicInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterBasicInfoResponse struct {
	responses.BaseResponse
	RequestId   string
	ClusterInfo DescribeClusterBasicInfoClusterInfo
}

type DescribeClusterBasicInfoClusterInfo struct {
	Id                     string
	BizId                  string
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
	VpcId                  string
	VSwitchId              string
	NetType                string
	UserDefinedEmrEcsRole  string
	IoOptimized            bool
	InstanceGeneration     string
	BootstrapFailed        bool
	Configurations         string
	ImageId                string
	SecurityGroupId        string
	SecurityGroupName      string
	EasEnable              bool
	BootstrapActionList    DescribeClusterBasicInfoBootstrapActionList
	SoftwareInfo           DescribeClusterBasicInfoSoftwareInfo
	FailReason             DescribeClusterBasicInfoFailReason
}

type DescribeClusterBasicInfoBootstrapAction struct {
	Name string
	Path string
	Arg  string
}

type DescribeClusterBasicInfoSoftwareInfo struct {
	EmrVer      string
	ClusterType string
	Softwares   DescribeClusterBasicInfoSoftwareList
}

type DescribeClusterBasicInfoSoftware struct {
	DisplayName string
	Name        string
	OnlyDisplay bool
	StartTpe    int
	Version     string
}

type DescribeClusterBasicInfoFailReason struct {
	ErrorCode string
	ErrorMsg  string
	RequestId string
}

type DescribeClusterBasicInfoBootstrapActionList []DescribeClusterBasicInfoBootstrapAction

func (list *DescribeClusterBasicInfoBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterBasicInfoBootstrapAction)
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

type DescribeClusterBasicInfoSoftwareList []DescribeClusterBasicInfoSoftware

func (list *DescribeClusterBasicInfoSoftwareList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterBasicInfoSoftware)
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
