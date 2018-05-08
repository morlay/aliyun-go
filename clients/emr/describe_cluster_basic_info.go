package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId   common.String
	ClusterInfo DescribeClusterBasicInfoClusterInfo
}

type DescribeClusterBasicInfoClusterInfo struct {
	Id                     common.String
	BizId                  common.String
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
	VpcId                  common.String
	VSwitchId              common.String
	NetType                common.String
	UserDefinedEmrEcsRole  common.String
	IoOptimized            bool
	InstanceGeneration     common.String
	BootstrapFailed        bool
	Configurations         common.String
	ImageId                common.String
	SecurityGroupId        common.String
	SecurityGroupName      common.String
	EasEnable              bool
	BootstrapActionList    DescribeClusterBasicInfoBootstrapActionList
	SoftwareInfo           DescribeClusterBasicInfoSoftwareInfo
	FailReason             DescribeClusterBasicInfoFailReason
}

type DescribeClusterBasicInfoBootstrapAction struct {
	Name common.String
	Path common.String
	Arg  common.String
}

type DescribeClusterBasicInfoSoftwareInfo struct {
	EmrVer      common.String
	ClusterType common.String
	Softwares   DescribeClusterBasicInfoSoftwareList
}

type DescribeClusterBasicInfoSoftware struct {
	DisplayName common.String
	Name        common.String
	OnlyDisplay bool
	StartTpe    common.Integer
	Version     common.String
}

type DescribeClusterBasicInfoFailReason struct {
	ErrorCode common.String
	ErrorMsg  common.String
	RequestId common.String
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
