package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstancesRequest struct {
	requests.RpcRequest
	Tag4Value            string `position:"Query" name:"Tag.4.Value"`
	InnerIpAddresses     string `position:"Query" name:"InnerIpAddresses"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Tag2Key              string `position:"Query" name:"Tag.2.Key"`
	PrivateIpAddresses   string `position:"Query" name:"PrivateIpAddresses"`
	HpcClusterId         string `position:"Query" name:"HpcClusterId"`
	Filter2Value         string `position:"Query" name:"Filter.2.Value"`
	Tag3Key              string `position:"Query" name:"Tag.3.Key"`
	KeyPairName          string `position:"Query" name:"KeyPairName"`
	Tag1Value            string `position:"Query" name:"Tag.1.Value"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	LockReason           string `position:"Query" name:"LockReason"`
	Filter1Key           string `position:"Query" name:"Filter.1.Key"`
	DeviceAvailable      string `position:"Query" name:"DeviceAvailable"`
	Filter3Value         string `position:"Query" name:"Filter.3.Value"`
	Tag5Key              string `position:"Query" name:"Tag.5.Key"`
	Filter1Value         string `position:"Query" name:"Filter.1.Value"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	InstanceName         string `position:"Query" name:"InstanceName"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	Tag4Key              string `position:"Query" name:"Tag.4.Key"`
	InstanceNetworkType  string `position:"Query" name:"InstanceNetworkType"`
	Status               string `position:"Query" name:"Status"`
	ImageId              string `position:"Query" name:"ImageId"`
	Filter4Value         string `position:"Query" name:"Filter.4.Value"`
	IoOptimized          string `position:"Query" name:"IoOptimized"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	Filter4Key           string `position:"Query" name:"Filter.4.Key"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	RdmaIpAddresses      string `position:"Query" name:"RdmaIpAddresses"`
	PageSize             int    `position:"Query" name:"PageSize"`
	PublicIpAddresses    string `position:"Query" name:"PublicIpAddresses"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	InstanceChargeType   string `position:"Query" name:"InstanceChargeType"`
	Tag3Value            string `position:"Query" name:"Tag.3.Value"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	InstanceTypeFamily   string `position:"Query" name:"InstanceTypeFamily"`
	Filter2Key           string `position:"Query" name:"Filter.2.Key"`
	Tag5Value            string `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string `position:"Query" name:"Tag.1.Key"`
	EipAddresses         string `position:"Query" name:"EipAddresses"`
	VpcId                string `position:"Query" name:"VpcId"`
	Tag2Value            string `position:"Query" name:"Tag.2.Value"`
	Filter3Key           string `position:"Query" name:"Filter.3.Key"`
}

func (req *DescribeInstancesRequest) Invoke(client *sdk.Client) (resp *DescribeInstancesResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstances", "ecs", "")
	resp = &DescribeInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInstancesResponse struct {
	responses.BaseResponse
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Instances  DescribeInstancesInstanceList
}

type DescribeInstancesInstance struct {
	InstanceId              string
	InstanceName            string
	Description             string
	ImageId                 string
	OSName                  string
	OSType                  string
	RegionId                string
	ZoneId                  string
	ClusterId               string
	InstanceType            string
	Cpu                     int
	Memory                  int
	HostName                string
	Status                  string
	SerialNumber            string
	InternetChargeType      string
	InternetMaxBandwidthIn  int
	InternetMaxBandwidthOut int
	VlanId                  string
	CreationTime            string
	InstanceNetworkType     string
	InstanceChargeType      string
	SaleCycle               string
	ExpiredTime             string
	AutoReleaseTime         string
	IoOptimized             bool
	DeviceAvailable         bool
	InstanceTypeFamily      string
	LocalStorageCapacity    int64
	LocalStorageAmount      int
	GPUAmount               int
	GPUSpec                 string
	SpotStrategy            string
	SpotPriceLimit          float32
	ResourceGroupId         string
	KeyPairName             string
	Recyclable              bool
	HpcClusterId            string
	StoppedMode             string
	NetworkInterfaces       DescribeInstancesNetworkInterfaceList
	OperationLocks          DescribeInstancesLockReasonList
	Tags                    DescribeInstancesTagList
	SecurityGroupIds        DescribeInstancesSecurityGroupIdList
	PublicIpAddress         DescribeInstancesPublicIpAddresList
	InnerIpAddress          DescribeInstancesInnerIpAddresList
	RdmaIpAddress           DescribeInstancesRdmaIpAddresList
	VpcAttributes           DescribeInstancesVpcAttributes
	EipAddress              DescribeInstancesEipAddress
}

type DescribeInstancesNetworkInterface struct {
	NetworkInterfaceId string
	MacAddress         string
	PrimaryIpAddress   string
}

type DescribeInstancesLockReason struct {
	LockReason string
	LockMsg    string
}

type DescribeInstancesTag struct {
	TagKey   string
	TagValue string
}

type DescribeInstancesVpcAttributes struct {
	VpcId            string
	VSwitchId        string
	NatIpAddress     string
	PrivateIpAddress DescribeInstancesPrivateIpAddresList
}

type DescribeInstancesEipAddress struct {
	AllocationId         string
	IpAddress            string
	Bandwidth            int
	InternetChargeType   string
	IsSupportUnassociate bool
}

type DescribeInstancesInstanceList []DescribeInstancesInstance

func (list *DescribeInstancesInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesInstance)
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

type DescribeInstancesNetworkInterfaceList []DescribeInstancesNetworkInterface

func (list *DescribeInstancesNetworkInterfaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesNetworkInterface)
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

type DescribeInstancesLockReasonList []DescribeInstancesLockReason

func (list *DescribeInstancesLockReasonList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesLockReason)
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

type DescribeInstancesTagList []DescribeInstancesTag

func (list *DescribeInstancesTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesTag)
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

type DescribeInstancesSecurityGroupIdList []string

func (list *DescribeInstancesSecurityGroupIdList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesPublicIpAddresList []string

func (list *DescribeInstancesPublicIpAddresList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesInnerIpAddresList []string

func (list *DescribeInstancesInnerIpAddresList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesRdmaIpAddresList []string

func (list *DescribeInstancesRdmaIpAddresList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesPrivateIpAddresList []string

func (list *DescribeInstancesPrivateIpAddresList) UnmarshalJSON(data []byte) error {
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
