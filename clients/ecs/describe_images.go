package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeImagesRequest struct {
	requests.RpcRequest
	Tag4Value            string                    `position:"Query" name:"Tag.4.Value"`
	ActionType           string                    `position:"Query" name:"ActionType"`
	ResourceOwnerId      int64                     `position:"Query" name:"ResourceOwnerId"`
	ImageId              string                    `position:"Query" name:"ImageId"`
	SnapshotId           string                    `position:"Query" name:"SnapshotId"`
	Tag2Key              string                    `position:"Query" name:"Tag.2.Key"`
	Usage                string                    `position:"Query" name:"Usage"`
	Tag3Key              string                    `position:"Query" name:"Tag.3.Key"`
	PageNumber           int                       `position:"Query" name:"PageNumber"`
	ImageOwnerAlias      string                    `position:"Query" name:"ImageOwnerAlias"`
	Tag1Value            string                    `position:"Query" name:"Tag.1.Value"`
	IsSupportIoOptimized string                    `position:"Query" name:"IsSupportIoOptimized"`
	ImageName            string                    `position:"Query" name:"ImageName"`
	IsSupportCloudinit   string                    `position:"Query" name:"IsSupportCloudinit"`
	PageSize             int                       `position:"Query" name:"PageSize"`
	InstanceType         string                    `position:"Query" name:"InstanceType"`
	Tag3Value            string                    `position:"Query" name:"Tag.3.Value"`
	Architecture         string                    `position:"Query" name:"Architecture"`
	DryRun               string                    `position:"Query" name:"DryRun"`
	Tag5Key              string                    `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string                    `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                    `position:"Query" name:"OwnerAccount"`
	ShowExpired          string                    `position:"Query" name:"ShowExpired"`
	OSType               string                    `position:"Query" name:"OSType"`
	OwnerId              int64                     `position:"Query" name:"OwnerId"`
	Tag5Value            string                    `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string                    `position:"Query" name:"Tag.1.Key"`
	Filters              *DescribeImagesFilterList `position:"Query" type:"Repeated" name:"Filter"`
	Tag2Value            string                    `position:"Query" name:"Tag.2.Value"`
	Tag4Key              string                    `position:"Query" name:"Tag.4.Key"`
	Status               string                    `position:"Query" name:"Status"`
}

func (req *DescribeImagesRequest) Invoke(client *sdk.Client) (resp *DescribeImagesResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeImages", "ecs", "")
	resp = &DescribeImagesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeImagesFilter struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

type DescribeImagesResponse struct {
	responses.BaseResponse
	RequestId  common.String
	RegionId   common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	Images     DescribeImagesImageList
}

type DescribeImagesImage struct {
	Progress             common.String
	ImageId              common.String
	ImageName            common.String
	ImageVersion         common.String
	Description          common.String
	Size                 common.Integer
	ImageOwnerAlias      common.String
	IsSupportIoOptimized bool
	IsSupportCloudinit   bool
	OSName               common.String
	Architecture         common.String
	Status               common.String
	ProductCode          common.String
	IsSubscribed         bool
	CreationTime         common.String
	IsSelfShared         common.String
	OSType               common.String
	Platform             common.String
	Usage                common.String
	IsCopied             bool
	DiskDeviceMappings   DescribeImagesDiskDeviceMappingList
	Tags                 DescribeImagesTagList
}

type DescribeImagesDiskDeviceMapping struct {
	SnapshotId      common.String
	Size            common.String
	Device          common.String
	Type            common.String
	Format          common.String
	ImportOSSBucket common.String
	ImportOSSObject common.String
}

type DescribeImagesTag struct {
	TagKey   common.String
	TagValue common.String
}

type DescribeImagesFilterList []DescribeImagesFilter

func (list *DescribeImagesFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImagesFilter)
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

type DescribeImagesImageList []DescribeImagesImage

func (list *DescribeImagesImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImagesImage)
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

type DescribeImagesDiskDeviceMappingList []DescribeImagesDiskDeviceMapping

func (list *DescribeImagesDiskDeviceMappingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImagesDiskDeviceMapping)
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

type DescribeImagesTagList []DescribeImagesTag

func (list *DescribeImagesTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImagesTag)
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
