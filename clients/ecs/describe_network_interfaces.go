package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeNetworkInterfacesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64                                            `position:"Query" name:"ResourceOwnerId"`
	SecurityGroupId      string                                           `position:"Query" name:"SecurityGroupId"`
	Type                 string                                           `position:"Query" name:"Type"`
	PageNumber           int                                              `position:"Query" name:"PageNumber"`
	PageSize             int                                              `position:"Query" name:"PageSize"`
	NetworkInterfaceName string                                           `position:"Query" name:"NetworkInterfaceName"`
	ResourceOwnerAccount string                                           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                           `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                                            `position:"Query" name:"OwnerId"`
	VSwitchId            string                                           `position:"Query" name:"VSwitchId"`
	InstanceId           string                                           `position:"Query" name:"InstanceId"`
	PrimaryIpAddress     string                                           `position:"Query" name:"PrimaryIpAddress"`
	NetworkInterfaceIds  *DescribeNetworkInterfacesNetworkInterfaceIdList `position:"Query" type:"Repeated" name:"NetworkInterfaceId"`
}

func (req *DescribeNetworkInterfacesRequest) Invoke(client *sdk.Client) (resp *DescribeNetworkInterfacesResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeNetworkInterfaces", "ecs", "")
	resp = &DescribeNetworkInterfacesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeNetworkInterfacesResponse struct {
	responses.BaseResponse
	RequestId            common.String
	TotalCount           common.Integer
	PageNumber           common.Integer
	PageSize             common.Integer
	NetworkInterfaceSets DescribeNetworkInterfacesNetworkInterfaceSetList
}

type DescribeNetworkInterfacesNetworkInterfaceSet struct {
	NetworkInterfaceId   common.String
	Status               common.String
	Type                 common.String
	VpcId                common.String
	VSwitchId            common.String
	ZoneId               common.String
	PrivateIpAddress     common.String
	MacAddress           common.String
	NetworkInterfaceName common.String
	Description          common.String
	InstanceId           common.String
	CreationTime         common.String
	PrivateIpSets        DescribeNetworkInterfacesPrivateIpSetList
	SecurityGroupIds     DescribeNetworkInterfacesSecurityGroupIdList
	AssociatedPublicIp   DescribeNetworkInterfacesAssociatedPublicIp
}

type DescribeNetworkInterfacesPrivateIpSet struct {
	PrivateIpAddress    common.String
	Primary             bool
	AssociatedPublicIp1 DescribeNetworkInterfacesAssociatedPublicIp1
}

type DescribeNetworkInterfacesAssociatedPublicIp1 struct {
	PublicIpAddress common.String
	AllocationId    common.String
}

type DescribeNetworkInterfacesAssociatedPublicIp struct {
	PublicIpAddress common.String
	AllocationId    common.String
}

type DescribeNetworkInterfacesNetworkInterfaceIdList []string

func (list *DescribeNetworkInterfacesNetworkInterfaceIdList) UnmarshalJSON(data []byte) error {
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

type DescribeNetworkInterfacesNetworkInterfaceSetList []DescribeNetworkInterfacesNetworkInterfaceSet

func (list *DescribeNetworkInterfacesNetworkInterfaceSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNetworkInterfacesNetworkInterfaceSet)
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

type DescribeNetworkInterfacesPrivateIpSetList []DescribeNetworkInterfacesPrivateIpSet

func (list *DescribeNetworkInterfacesPrivateIpSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNetworkInterfacesPrivateIpSet)
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

type DescribeNetworkInterfacesSecurityGroupIdList []common.String

func (list *DescribeNetworkInterfacesSecurityGroupIdList) UnmarshalJSON(data []byte) error {
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
