package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeResourcesModificationRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	MigrateAcrossZone    string `position:"Query" name:"MigrateAcrossZone"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OperationType        string `position:"Query" name:"OperationType"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DestinationResource  string `position:"Query" name:"DestinationResource"`
}

func (req *DescribeResourcesModificationRequest) Invoke(client *sdk.Client) (resp *DescribeResourcesModificationResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeResourcesModification", "ecs", "")
	resp = &DescribeResourcesModificationResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeResourcesModificationResponse struct {
	responses.BaseResponse
	RequestId      string
	AvailableZones DescribeResourcesModificationAvailableZoneList
}

type DescribeResourcesModificationAvailableZone struct {
	RegionId           string
	ZoneId             string
	Status             string
	AvailableResources DescribeResourcesModificationAvailableResourceList
}

type DescribeResourcesModificationAvailableResource struct {
	Type               string
	SupportedResources DescribeResourcesModificationSupportedResourceList
}

type DescribeResourcesModificationSupportedResource struct {
	Value  string
	Status string
	Min    int
	Max    int
	Unit   string
}

type DescribeResourcesModificationAvailableZoneList []DescribeResourcesModificationAvailableZone

func (list *DescribeResourcesModificationAvailableZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeResourcesModificationAvailableZone)
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

type DescribeResourcesModificationAvailableResourceList []DescribeResourcesModificationAvailableResource

func (list *DescribeResourcesModificationAvailableResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeResourcesModificationAvailableResource)
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

type DescribeResourcesModificationSupportedResourceList []DescribeResourcesModificationSupportedResource

func (list *DescribeResourcesModificationSupportedResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeResourcesModificationSupportedResource)
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
