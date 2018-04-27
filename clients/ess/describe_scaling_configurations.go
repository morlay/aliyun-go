package ess

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeScalingConfigurationsRequest struct {
	requests.RpcRequest
	ScalingConfigurationId6    string `position:"Query" name:"ScalingConfigurationId.6"`
	ScalingConfigurationId7    string `position:"Query" name:"ScalingConfigurationId.7"`
	ResourceOwnerId            int64  `position:"Query" name:"ResourceOwnerId"`
	ScalingConfigurationId4    string `position:"Query" name:"ScalingConfigurationId.4"`
	ScalingConfigurationId5    string `position:"Query" name:"ScalingConfigurationId.5"`
	ScalingGroupId             string `position:"Query" name:"ScalingGroupId"`
	ScalingConfigurationId8    string `position:"Query" name:"ScalingConfigurationId.8"`
	ScalingConfigurationId9    string `position:"Query" name:"ScalingConfigurationId.9"`
	ScalingConfigurationId10   string `position:"Query" name:"ScalingConfigurationId.10"`
	PageNumber                 int    `position:"Query" name:"PageNumber"`
	ScalingConfigurationName2  string `position:"Query" name:"ScalingConfigurationName.2"`
	ScalingConfigurationName3  string `position:"Query" name:"ScalingConfigurationName.3"`
	ScalingConfigurationName1  string `position:"Query" name:"ScalingConfigurationName.1"`
	PageSize                   int    `position:"Query" name:"PageSize"`
	ScalingConfigurationId2    string `position:"Query" name:"ScalingConfigurationId.2"`
	ScalingConfigurationId3    string `position:"Query" name:"ScalingConfigurationId.3"`
	ScalingConfigurationId1    string `position:"Query" name:"ScalingConfigurationId.1"`
	ResourceOwnerAccount       string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount               string `position:"Query" name:"OwnerAccount"`
	ScalingConfigurationName6  string `position:"Query" name:"ScalingConfigurationName.6"`
	ScalingConfigurationName7  string `position:"Query" name:"ScalingConfigurationName.7"`
	ScalingConfigurationName4  string `position:"Query" name:"ScalingConfigurationName.4"`
	ScalingConfigurationName5  string `position:"Query" name:"ScalingConfigurationName.5"`
	OwnerId                    int64  `position:"Query" name:"OwnerId"`
	ScalingConfigurationName8  string `position:"Query" name:"ScalingConfigurationName.8"`
	ScalingConfigurationName9  string `position:"Query" name:"ScalingConfigurationName.9"`
	ScalingConfigurationName10 string `position:"Query" name:"ScalingConfigurationName.10"`
}

func (req *DescribeScalingConfigurationsRequest) Invoke(client *sdk.Client) (resp *DescribeScalingConfigurationsResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DescribeScalingConfigurations", "ess", "")
	resp = &DescribeScalingConfigurationsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeScalingConfigurationsResponse struct {
	responses.BaseResponse
	TotalCount            int
	PageNumber            int
	PageSize              int
	RequestId             string
	ScalingConfigurations DescribeScalingConfigurationsScalingConfigurationList
}

type DescribeScalingConfigurationsScalingConfiguration struct {
	ScalingConfigurationId      string
	ScalingConfigurationName    string
	ScalingGroupId              string
	InstanceName                string
	ImageId                     string
	InstanceType                string
	InstanceGeneration          string
	SecurityGroupId             string
	IoOptimized                 string
	InternetChargeType          string
	InternetMaxBandwidthIn      int
	InternetMaxBandwidthOut     int
	SystemDiskCategory          string
	SystemDiskSize              int
	LifecycleState              string
	CreationTime                string
	LoadBalancerWeight          int
	UserData                    string
	KeyPairName                 string
	RamRoleName                 string
	DeploymentSetId             string
	SecurityEnhancementStrategy string
	SpotStrategy                string
	DataDisks                   DescribeScalingConfigurationsDataDiskList
	Tags                        DescribeScalingConfigurationsTagList
	SpotPriceLimit              DescribeScalingConfigurationsSpotPriceModelList
	InstanceTypes               DescribeScalingConfigurationsInstanceTypeList
}

type DescribeScalingConfigurationsDataDisk struct {
	Size       int
	Category   string
	SnapshotId string
	Device     string
}

type DescribeScalingConfigurationsTag struct {
	Key   string
	Value string
}

type DescribeScalingConfigurationsSpotPriceModel struct {
	InstanceType string
	PriceLimit   float32
}

type DescribeScalingConfigurationsScalingConfigurationList []DescribeScalingConfigurationsScalingConfiguration

func (list *DescribeScalingConfigurationsScalingConfigurationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingConfigurationsScalingConfiguration)
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

type DescribeScalingConfigurationsDataDiskList []DescribeScalingConfigurationsDataDisk

func (list *DescribeScalingConfigurationsDataDiskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingConfigurationsDataDisk)
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

type DescribeScalingConfigurationsTagList []DescribeScalingConfigurationsTag

func (list *DescribeScalingConfigurationsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingConfigurationsTag)
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

type DescribeScalingConfigurationsSpotPriceModelList []DescribeScalingConfigurationsSpotPriceModel

func (list *DescribeScalingConfigurationsSpotPriceModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingConfigurationsSpotPriceModel)
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

type DescribeScalingConfigurationsInstanceTypeList []string

func (list *DescribeScalingConfigurationsInstanceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
