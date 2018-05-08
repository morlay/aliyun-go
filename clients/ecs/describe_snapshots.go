package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeSnapshotsRequest struct {
	requests.RpcRequest
	Tag4Value            string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Tag2Key              string `position:"Query" name:"Tag.2.Key"`
	Filter2Value         string `position:"Query" name:"Filter.2.Value"`
	SnapshotIds          string `position:"Query" name:"SnapshotIds"`
	Usage                string `position:"Query" name:"Usage"`
	SnapshotLinkId       string `position:"Query" name:"SnapshotLinkId"`
	SnapshotName         string `position:"Query" name:"SnapshotName"`
	Tag3Key              string `position:"Query" name:"Tag.3.Key"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	Tag1Value            string `position:"Query" name:"Tag.1.Value"`
	Filter1Key           string `position:"Query" name:"Filter.1.Key"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DiskId               string `position:"Query" name:"DiskId"`
	Tag3Value            string `position:"Query" name:"Tag.3.Value"`
	DryRun               string `position:"Query" name:"DryRun"`
	Tag5Key              string `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SourceDiskType       string `position:"Query" name:"SourceDiskType"`
	Filter1Value         string `position:"Query" name:"Filter.1.Value"`
	Filter2Key           string `position:"Query" name:"Filter.2.Key"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tag5Value            string `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string `position:"Query" name:"Tag.1.Key"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	Encrypted            string `position:"Query" name:"Encrypted"`
	SnapshotType         string `position:"Query" name:"SnapshotType"`
	Tag2Value            string `position:"Query" name:"Tag.2.Value"`
	Tag4Key              string `position:"Query" name:"Tag.4.Key"`
	Status               string `position:"Query" name:"Status"`
}

func (req *DescribeSnapshotsRequest) Invoke(client *sdk.Client) (resp *DescribeSnapshotsResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSnapshots", "ecs", "")
	resp = &DescribeSnapshotsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSnapshotsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	Snapshots  DescribeSnapshotsSnapshotList
}

type DescribeSnapshotsSnapshot struct {
	SnapshotId        common.String
	SnapshotName      common.String
	Progress          common.String
	ProductCode       common.String
	SourceDiskId      common.String
	SourceDiskType    common.String
	RetentionDays     common.Integer
	Encrypted         bool
	SourceDiskSize    common.String
	Description       common.String
	CreationTime      common.String
	Status            common.String
	Usage             common.String
	SourceStorageType common.String
	Tags              DescribeSnapshotsTagList
}

type DescribeSnapshotsTag struct {
	TagKey   common.String
	TagValue common.String
}

type DescribeSnapshotsSnapshotList []DescribeSnapshotsSnapshot

func (list *DescribeSnapshotsSnapshotList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSnapshotsSnapshot)
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

type DescribeSnapshotsTagList []DescribeSnapshotsTag

func (list *DescribeSnapshotsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSnapshotsTag)
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
