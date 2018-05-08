package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeAvailableResourceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	IoOptimized          string `position:"Query" name:"IoOptimized"`
	DataDiskCategory     string `position:"Query" name:"DataDiskCategory"`
	SystemDiskCategory   string `position:"Query" name:"SystemDiskCategory"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	NetworkCategory      string `position:"Query" name:"NetworkCategory"`
	InstanceChargeType   string `position:"Query" name:"InstanceChargeType"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DedicatedHostId      string `position:"Query" name:"DedicatedHostId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SpotStrategy         string `position:"Query" name:"SpotStrategy"`
	DestinationResource  string `position:"Query" name:"DestinationResource"`
	ZoneId               string `position:"Query" name:"ZoneId"`
}

func (req *DescribeAvailableResourceRequest) Invoke(client *sdk.Client) (resp *DescribeAvailableResourceResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeAvailableResource", "ecs", "")
	resp = &DescribeAvailableResourceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAvailableResourceResponse struct {
	responses.BaseResponse
	RequestId      common.String
	AvailableZones DescribeAvailableResourceAvailableZoneList
}

type DescribeAvailableResourceAvailableZone struct {
	RegionId           common.String
	ZoneId             common.String
	Status             common.String
	AvailableResources DescribeAvailableResourceAvailableResourceList
}

type DescribeAvailableResourceAvailableResource struct {
	Type               common.String
	SupportedResources DescribeAvailableResourceSupportedResourceList
}

type DescribeAvailableResourceSupportedResource struct {
	Value  common.String
	Status common.String
	Min    common.Integer
	Max    common.Integer
	Unit   common.String
}

type DescribeAvailableResourceAvailableZoneList []DescribeAvailableResourceAvailableZone

func (list *DescribeAvailableResourceAvailableZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAvailableResourceAvailableZone)
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

type DescribeAvailableResourceAvailableResourceList []DescribeAvailableResourceAvailableResource

func (list *DescribeAvailableResourceAvailableResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAvailableResourceAvailableResource)
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

type DescribeAvailableResourceSupportedResourceList []DescribeAvailableResourceSupportedResource

func (list *DescribeAvailableResourceSupportedResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAvailableResourceSupportedResource)
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
