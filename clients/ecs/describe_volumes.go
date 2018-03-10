package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeVolumesRequest struct {
	requests.RpcRequest
	Tag4Value                     string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId               int64  `position:"Query" name:"ResourceOwnerId"`
	SnapshotId                    string `position:"Query" name:"SnapshotId"`
	Tag2Key                       string `position:"Query" name:"Tag.2.Key"`
	AutoSnapshotPolicyId          string `position:"Query" name:"AutoSnapshotPolicyId"`
	Tag3Key                       string `position:"Query" name:"Tag.3.Key"`
	PageNumber                    int    `position:"Query" name:"PageNumber"`
	Tag1Value                     string `position:"Query" name:"Tag.1.Value"`
	LockReason                    string `position:"Query" name:"LockReason"`
	PageSize                      int    `position:"Query" name:"PageSize"`
	Tag3Value                     string `position:"Query" name:"Tag.3.Value"`
	Tag5Key                       string `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount          string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                  string `position:"Query" name:"OwnerAccount"`
	EnableAutomatedSnapshotPolicy string `position:"Query" name:"EnableAutomatedSnapshotPolicy"`
	OwnerId                       int64  `position:"Query" name:"OwnerId"`
	Tag5Value                     string `position:"Query" name:"Tag.5.Value"`
	Tag1Key                       string `position:"Query" name:"Tag.1.Key"`
	InstanceId                    string `position:"Query" name:"InstanceId"`
	Tag2Value                     string `position:"Query" name:"Tag.2.Value"`
	ZoneId                        string `position:"Query" name:"ZoneId"`
	Tag4Key                       string `position:"Query" name:"Tag.4.Key"`
	VolumeIds                     string `position:"Query" name:"VolumeIds"`
	Category                      string `position:"Query" name:"Category"`
	Status                        string `position:"Query" name:"Status"`
}

func (req *DescribeVolumesRequest) Invoke(client *sdk.Client) (resp *DescribeVolumesResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeVolumes", "ecs", "")
	resp = &DescribeVolumesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeVolumesResponse struct {
	responses.BaseResponse
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Volumes    DescribeVolumesVolumeList
}

type DescribeVolumesVolume struct {
	VolumeId                      string
	RegionId                      string
	ZoneId                        string
	VolumeName                    string
	Description                   string
	Category                      string
	Size                          int
	SourceSnapshotId              string
	AutoSnapshotPolicyId          string
	SnapshotLinkId                string
	Status                        string
	EnableAutomatedSnapshotPolicy bool
	CreationTime                  string
	VolumeChargeType              string
	MountInstanceNum              int
	Encrypted                     bool
	OperationLocks                DescribeVolumesOperationLockList
	MountInstances                DescribeVolumesMountInstanceList
	Tags                          DescribeVolumesTagList
}

type DescribeVolumesOperationLock struct {
	LockReason string
}

type DescribeVolumesMountInstance struct {
	InstanceId   string
	Device       string
	AttachedTime string
}

type DescribeVolumesTag struct {
	TagKey   string
	TagValue string
}

type DescribeVolumesVolumeList []DescribeVolumesVolume

func (list *DescribeVolumesVolumeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVolumesVolume)
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

type DescribeVolumesOperationLockList []DescribeVolumesOperationLock

func (list *DescribeVolumesOperationLockList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVolumesOperationLock)
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

type DescribeVolumesMountInstanceList []DescribeVolumesMountInstance

func (list *DescribeVolumesMountInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVolumesMountInstance)
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

type DescribeVolumesTagList []DescribeVolumesTag

func (list *DescribeVolumesTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVolumesTag)
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
