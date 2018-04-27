package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstanceAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeInstanceAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeInstanceAttributeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceAttribute", "ecs", "")
	resp = &DescribeInstanceAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInstanceAttributeResponse struct {
	responses.BaseResponse
	RequestId               string
	InstanceId              string
	InstanceName            string
	ImageId                 string
	RegionId                string
	ZoneId                  string
	ClusterId               string
	InstanceType            string
	Cpu                     int
	Memory                  int
	HostName                string
	Status                  string
	InternetChargeType      string
	InternetMaxBandwidthIn  int
	InternetMaxBandwidthOut int
	VlanId                  string
	SerialNumber            string
	CreationTime            string
	Description             string
	InstanceNetworkType     string
	IoOptimized             string
	InstanceChargeType      string
	ExpiredTime             string
	StoppedMode             string
	OperationLocks          DescribeInstanceAttributeLockReasonList
	SecurityGroupIds        DescribeInstanceAttributeSecurityGroupIdList
	PublicIpAddress         DescribeInstanceAttributePublicIpAddresList
	InnerIpAddress          DescribeInstanceAttributeInnerIpAddresList
	VpcAttributes           DescribeInstanceAttributeVpcAttributes
	EipAddress              DescribeInstanceAttributeEipAddress
	DedicatedHostAttribute  DescribeInstanceAttributeDedicatedHostAttribute
}

type DescribeInstanceAttributeLockReason struct {
	LockReason string
}

type DescribeInstanceAttributeVpcAttributes struct {
	VpcId            string
	VSwitchId        string
	NatIpAddress     string
	PrivateIpAddress DescribeInstanceAttributePrivateIpAddresList
}

type DescribeInstanceAttributeEipAddress struct {
	AllocationId       string
	IpAddress          string
	Bandwidth          int
	InternetChargeType string
}

type DescribeInstanceAttributeDedicatedHostAttribute struct {
	DedicatedHostId   string
	DedicatedHostName string
}

type DescribeInstanceAttributeLockReasonList []DescribeInstanceAttributeLockReason

func (list *DescribeInstanceAttributeLockReasonList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceAttributeLockReason)
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

type DescribeInstanceAttributeSecurityGroupIdList []string

func (list *DescribeInstanceAttributeSecurityGroupIdList) UnmarshalJSON(data []byte) error {
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

type DescribeInstanceAttributePublicIpAddresList []string

func (list *DescribeInstanceAttributePublicIpAddresList) UnmarshalJSON(data []byte) error {
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

type DescribeInstanceAttributeInnerIpAddresList []string

func (list *DescribeInstanceAttributeInnerIpAddresList) UnmarshalJSON(data []byte) error {
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

type DescribeInstanceAttributePrivateIpAddresList []string

func (list *DescribeInstanceAttributePrivateIpAddresList) UnmarshalJSON(data []byte) error {
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
