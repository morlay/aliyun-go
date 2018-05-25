package market_inner

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InnerQueryAvailableImagesRequest struct {
	requests.RpcRequest
	SupportIoOptimized string                                `position:"Query" name:"SupportIoOptimized"`
	SnapshotId         string                                `position:"Query" name:"SnapshotId"`
	ImageIds           *InnerQueryAvailableImagesImageIdList `position:"Query" type:"Repeated" name:"ImageId"`
	ImageName          string                                `position:"Query" name:"ImageName"`
	PageSize           int                                   `position:"Query" name:"PageSize"`
	EcsUnits           *InnerQueryAvailableImagesEcsUnitList `position:"Query" type:"Repeated" name:"EcsUnit"`
	PageIndex          int                                   `position:"Query" name:"PageIndex"`
	AliUid             int64                                 `position:"Query" name:"AliUid"`
	OsTypes            *InnerQueryAvailableImagesOsTypeList  `position:"Query" type:"Repeated" name:"OsType"`
	RegionNo           string                                `position:"Query" name:"RegionNo"`
}

func (req *InnerQueryAvailableImagesRequest) Invoke(client *sdk.Client) (resp *InnerQueryAvailableImagesResponse, err error) {
	req.InitWithApiInfo("Market-Inner", "2016-08-01", "InnerQueryAvailableImages", "yunmarket", "")
	resp = &InnerQueryAvailableImagesResponse{}
	err = client.DoAction(req, resp)
	return
}

type InnerQueryAvailableImagesEcsUnit struct {
	Core   int `name:"Core"`
	Memory int `name:"Memory"`
}

type InnerQueryAvailableImagesOsType struct {
	OsKind string `name:"OsKind"`
	OsBit  int    `name:"OsBit"`
}

type InnerQueryAvailableImagesResponse struct {
	responses.BaseResponse
	PageNumber common.Integer
	PageSize   common.Integer
	TotalCount common.Integer
	RequestId  common.String
	ImageList  InnerQueryAvailableImagesImageList
}

type InnerQueryAvailableImagesImage struct {
	ImageId            common.String
	ImageSize          common.Integer
	SupplierId         common.Long
	SnapshotId         common.String
	Device             common.String
	ProductCode        common.String
	ProductSkuCode     common.String
	ImageVersion       common.String
	RegionId           common.String
	ImageName          common.String
	SupplierName       common.String
	OsName             common.String
	Architecture       common.String
	Description        common.String
	ImageOwnerAlias    common.String
	IsSubscribed       bool
	GmtCreated         common.String
	SupportIoOptimized bool
	VmType             common.String
	DiskDeviceMappings InnerQueryAvailableImagesDiskDeviceMappingList
}

type InnerQueryAvailableImagesDiskDeviceMapping struct {
	DiskType        common.String
	Format          common.String
	SnapshotId      common.String
	Size            common.Integer
	Device          common.String
	ImportOSSBucket common.String
	ImportOSSObject common.String
}

type InnerQueryAvailableImagesImageIdList []string

func (list *InnerQueryAvailableImagesImageIdList) UnmarshalJSON(data []byte) error {
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

type InnerQueryAvailableImagesEcsUnitList []InnerQueryAvailableImagesEcsUnit

func (list *InnerQueryAvailableImagesEcsUnitList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerQueryAvailableImagesEcsUnit)
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

type InnerQueryAvailableImagesOsTypeList []InnerQueryAvailableImagesOsType

func (list *InnerQueryAvailableImagesOsTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerQueryAvailableImagesOsType)
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

type InnerQueryAvailableImagesImageList []InnerQueryAvailableImagesImage

func (list *InnerQueryAvailableImagesImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerQueryAvailableImagesImage)
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

type InnerQueryAvailableImagesDiskDeviceMappingList []InnerQueryAvailableImagesDiskDeviceMapping

func (list *InnerQueryAvailableImagesDiskDeviceMappingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerQueryAvailableImagesDiskDeviceMapping)
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
