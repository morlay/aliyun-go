package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	DryRun               string `position:"Query" name:"DryRun"`
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
	RequestId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	Instances  DescribeInstancesInstanceList
}

type DescribeInstancesInstance struct {
	InstanceId              common.String
	InstanceName            common.String
	Description             common.String
	ImageId                 common.String
	OSName                  common.String
	OSType                  common.String
	RegionId                common.String
	ZoneId                  common.String
	ClusterId               common.String
	InstanceType            common.String
	Cpu                     common.Integer
	Memory                  common.Integer
	HostName                common.String
	Status                  common.String
	SerialNumber            common.String
	InternetChargeType      common.String
	InternetMaxBandwidthIn  common.Integer
	InternetMaxBandwidthOut common.Integer
	VlanId                  common.String
	CreationTime            common.String
	StartTime               common.String
	InstanceNetworkType     common.String
	InstanceChargeType      common.String
	SaleCycle               common.String
	ExpiredTime             common.String
	AutoReleaseTime         common.String
	IoOptimized             bool
	DeviceAvailable         bool
	InstanceTypeFamily      common.String
	LocalStorageCapacity    common.Long
	LocalStorageAmount      common.Integer
	GPUAmount               common.Integer
	GPUSpec                 common.String
	SpotStrategy            common.String
	SpotPriceLimit          common.Float
	ResourceGroupId         common.String
	KeyPairName             common.String
	Recyclable              bool
	HpcClusterId            common.String
	StoppedMode             common.String
	NetworkInterfaces       DescribeInstancesNetworkInterfaceList
	OperationLocks          DescribeInstancesLockReasonList
	Tags                    DescribeInstancesTagList
	SecurityGroupIds        DescribeInstancesSecurityGroupIdList
	PublicIpAddress         DescribeInstancesPublicIpAddresList
	InnerIpAddress          DescribeInstancesInnerIpAddresList
	RdmaIpAddress           DescribeInstancesRdmaIpAddresList
	VpcAttributes           DescribeInstancesVpcAttributes
	EipAddress              DescribeInstancesEipAddress
	DedicatedHostAttribute  DescribeInstancesDedicatedHostAttribute
}

type DescribeInstancesNetworkInterface struct {
	NetworkInterfaceId common.String
	MacAddress         common.String
	PrimaryIpAddress   common.String
}

type DescribeInstancesLockReason struct {
	LockReason common.String
	LockMsg    common.String
}

type DescribeInstancesTag struct {
	TagKey   common.String
	TagValue common.String
}

type DescribeInstancesVpcAttributes struct {
	VpcId            common.String
	VSwitchId        common.String
	NatIpAddress     common.String
	PrivateIpAddress DescribeInstancesPrivateIpAddresList
}

type DescribeInstancesEipAddress struct {
	AllocationId         common.String
	IpAddress            common.String
	Bandwidth            common.Integer
	InternetChargeType   common.String
	IsSupportUnassociate bool
}

type DescribeInstancesDedicatedHostAttribute struct {
	DedicatedHostId   common.String
	DedicatedHostName common.String
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

type DescribeInstancesSecurityGroupIdList []common.String

func (list *DescribeInstancesSecurityGroupIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type DescribeInstancesPublicIpAddresList []common.String

func (list *DescribeInstancesPublicIpAddresList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type DescribeInstancesInnerIpAddresList []common.String

func (list *DescribeInstancesInnerIpAddresList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type DescribeInstancesRdmaIpAddresList []common.String

func (list *DescribeInstancesRdmaIpAddresList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type DescribeInstancesPrivateIpAddresList []common.String

func (list *DescribeInstancesPrivateIpAddresList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
