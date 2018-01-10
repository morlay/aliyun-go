package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeImagesRequest struct {
	requests.RpcRequest
	Tag4Value            string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	SnapshotId           string `position:"Query" name:"SnapshotId"`
	Tag2Key              string `position:"Query" name:"Tag.2.Key"`
	Filter2Value         string `position:"Query" name:"Filter.2.Value"`
	Usage                string `position:"Query" name:"Usage"`
	Tag3Key              string `position:"Query" name:"Tag.3.Key"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	ImageOwnerAlias      string `position:"Query" name:"ImageOwnerAlias"`
	Tag1Value            string `position:"Query" name:"Tag.1.Value"`
	IsSupportIoOptimized string `position:"Query" name:"IsSupportIoOptimized"`
	Filter1Key           string `position:"Query" name:"Filter.1.Key"`
	ImageName            string `position:"Query" name:"ImageName"`
	IsSupportCloudinit   string `position:"Query" name:"IsSupportCloudinit"`
	PageSize             int    `position:"Query" name:"PageSize"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	Tag3Value            string `position:"Query" name:"Tag.3.Value"`
	Architecture         string `position:"Query" name:"Architecture"`
	DryRun               string `position:"Query" name:"DryRun"`
	Tag5Key              string `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ShowExpired          string `position:"Query" name:"ShowExpired"`
	Filter1Value         string `position:"Query" name:"Filter.1.Value"`
	OSType               string `position:"Query" name:"OSType"`
	Filter2Key           string `position:"Query" name:"Filter.2.Key"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tag5Value            string `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string `position:"Query" name:"Tag.1.Key"`
	Tag2Value            string `position:"Query" name:"Tag.2.Value"`
	Tag4Key              string `position:"Query" name:"Tag.4.Key"`
	Status               string `position:"Query" name:"Status"`
}

func (req *DescribeImagesRequest) Invoke(client *sdk.Client) (resp *DescribeImagesResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeImages", "ecs", "")
	resp = &DescribeImagesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeImagesResponse struct {
	responses.BaseResponse
	RequestId  string
	RegionId   string
	TotalCount int
	PageNumber int
	PageSize   int
	Images     DescribeImagesImageList
}

type DescribeImagesImage struct {
	Progress             string
	ImageId              string
	ImageName            string
	ImageVersion         string
	Description          string
	Size                 int
	ImageOwnerAlias      string
	IsSupportIoOptimized bool
	IsSupportCloudinit   bool
	OSName               string
	Architecture         string
	Status               string
	ProductCode          string
	IsSubscribed         bool
	CreationTime         string
	IsSelfShared         string
	OSType               string
	Platform             string
	Usage                string
	IsCopied             bool
	DiskDeviceMappings   DescribeImagesDiskDeviceMappingList
	Tags                 DescribeImagesTagList
}

type DescribeImagesDiskDeviceMapping struct {
	SnapshotId      string
	Size            string
	Device          string
	Type            string
	Format          string
	ImportOSSBucket string
	ImportOSSObject string
}

type DescribeImagesTag struct {
	TagKey   string
	TagValue string
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
