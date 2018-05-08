package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId               common.String
	InstanceId              common.String
	InstanceName            common.String
	ImageId                 common.String
	RegionId                common.String
	ZoneId                  common.String
	ClusterId               common.String
	InstanceType            common.String
	Cpu                     common.Integer
	Memory                  common.Integer
	HostName                common.String
	Status                  common.String
	InternetChargeType      common.String
	InternetMaxBandwidthIn  common.Integer
	InternetMaxBandwidthOut common.Integer
	VlanId                  common.String
	SerialNumber            common.String
	CreationTime            common.String
	Description             common.String
	InstanceNetworkType     common.String
	IoOptimized             common.String
	InstanceChargeType      common.String
	ExpiredTime             common.String
	StoppedMode             common.String
	OperationLocks          DescribeInstanceAttributeLockReasonList
	SecurityGroupIds        DescribeInstanceAttributeSecurityGroupIdList
	PublicIpAddress         DescribeInstanceAttributePublicIpAddresList
	InnerIpAddress          DescribeInstanceAttributeInnerIpAddresList
	VpcAttributes           DescribeInstanceAttributeVpcAttributes
	EipAddress              DescribeInstanceAttributeEipAddress
	DedicatedHostAttribute  DescribeInstanceAttributeDedicatedHostAttribute
}

type DescribeInstanceAttributeLockReason struct {
	LockReason common.String
}

type DescribeInstanceAttributeVpcAttributes struct {
	VpcId            common.String
	VSwitchId        common.String
	NatIpAddress     common.String
	PrivateIpAddress DescribeInstanceAttributePrivateIpAddresList
}

type DescribeInstanceAttributeEipAddress struct {
	AllocationId       common.String
	IpAddress          common.String
	Bandwidth          common.Integer
	InternetChargeType common.String
}

type DescribeInstanceAttributeDedicatedHostAttribute struct {
	DedicatedHostId   common.String
	DedicatedHostName common.String
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

type DescribeInstanceAttributeSecurityGroupIdList []common.String

func (list *DescribeInstanceAttributeSecurityGroupIdList) UnmarshalJSON(data []byte) error {
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

type DescribeInstanceAttributePublicIpAddresList []common.String

func (list *DescribeInstanceAttributePublicIpAddresList) UnmarshalJSON(data []byte) error {
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

type DescribeInstanceAttributeInnerIpAddresList []common.String

func (list *DescribeInstanceAttributeInnerIpAddresList) UnmarshalJSON(data []byte) error {
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

type DescribeInstanceAttributePrivateIpAddresList []common.String

func (list *DescribeInstanceAttributePrivateIpAddresList) UnmarshalJSON(data []byte) error {
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
