package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSnapshotsRequest struct {
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

func (r DescribeSnapshotsRequest) Invoke(client *sdk.Client) (response *DescribeSnapshotsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeSnapshotsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSnapshots", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeSnapshotsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeSnapshotsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeSnapshotsResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Snapshots  DescribeSnapshotsSnapshotList
}

type DescribeSnapshotsSnapshot struct {
	SnapshotId        string
	SnapshotName      string
	Progress          string
	ProductCode       string
	SourceDiskId      string
	SourceDiskType    string
	RetentionDays     int
	Encrypted         bool
	SourceDiskSize    string
	Description       string
	CreationTime      string
	Status            string
	Usage             string
	SourceStorageType string
	Tags              DescribeSnapshotsTagList
}

type DescribeSnapshotsTag struct {
	TagKey   string
	TagValue string
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
