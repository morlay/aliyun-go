package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeZonesRequest struct {
	requests.RpcRequest
	SpotStrategy         string `position:"Query" name:"SpotStrategy"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceChargeType   string `position:"Query" name:"InstanceChargeType"`
	Verbose              string `position:"Query" name:"Verbose"`
}

func (req *DescribeZonesRequest) Invoke(client *sdk.Client) (resp *DescribeZonesResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeZones", "ecs", "")
	resp = &DescribeZonesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeZonesResponse struct {
	responses.BaseResponse
	RequestId common.String
	Zones     DescribeZonesZoneList
}

type DescribeZonesZone struct {
	ZoneId                      common.String
	LocalName                   common.String
	AvailableResources          DescribeZonesResourcesInfoList
	AvailableResourceCreation   DescribeZonesAvailableResourceCreationList
	AvailableDiskCategories     DescribeZonesAvailableDiskCategoryList
	AvailableInstanceTypes      DescribeZonesAvailableInstanceTypeList
	AvailableVolumeCategories   DescribeZonesAvailableVolumeCategoryList
	AvailableDedicatedHostTypes DescribeZonesAvailableDedicatedHostTypeList
	DedicatedHostGenerations    DescribeZonesDedicatedHostGenerationList
}

type DescribeZonesResourcesInfo struct {
	IoOptimized          bool
	SystemDiskCategories DescribeZonesSystemDiskCategoryList
	DataDiskCategories   DescribeZonesDataDiskCategoryList
	NetworkTypes         DescribeZonesNetworkTypeList
	InstanceTypes        DescribeZonesInstanceTypeList
	InstanceTypeFamilies DescribeZonesInstanceTypeFamilyList
	InstanceGenerations  DescribeZonesInstanceGenerationList
}

type DescribeZonesZoneList []DescribeZonesZone

func (list *DescribeZonesZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZonesZone)
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

type DescribeZonesResourcesInfoList []DescribeZonesResourcesInfo

func (list *DescribeZonesResourcesInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZonesResourcesInfo)
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

type DescribeZonesAvailableResourceCreationList []common.String

func (list *DescribeZonesAvailableResourceCreationList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesAvailableDiskCategoryList []common.String

func (list *DescribeZonesAvailableDiskCategoryList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesAvailableInstanceTypeList []common.String

func (list *DescribeZonesAvailableInstanceTypeList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesAvailableVolumeCategoryList []common.String

func (list *DescribeZonesAvailableVolumeCategoryList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesAvailableDedicatedHostTypeList []common.String

func (list *DescribeZonesAvailableDedicatedHostTypeList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesDedicatedHostGenerationList []common.String

func (list *DescribeZonesDedicatedHostGenerationList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesSystemDiskCategoryList []common.String

func (list *DescribeZonesSystemDiskCategoryList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesDataDiskCategoryList []common.String

func (list *DescribeZonesDataDiskCategoryList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesNetworkTypeList []common.String

func (list *DescribeZonesNetworkTypeList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesInstanceTypeList []common.String

func (list *DescribeZonesInstanceTypeList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesInstanceTypeFamilyList []common.String

func (list *DescribeZonesInstanceTypeFamilyList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesInstanceGenerationList []common.String

func (list *DescribeZonesInstanceGenerationList) UnmarshalJSON(data []byte) error {
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
