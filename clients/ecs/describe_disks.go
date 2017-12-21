package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDisksRequest struct {
	Tag4Value                     string                                 `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId               int64                                  `position:"Query" name:"ResourceOwnerId"`
	SnapshotId                    string                                 `position:"Query" name:"SnapshotId"`
	Tag2Key                       string                                 `position:"Query" name:"Tag.2.Key"`
	Filter2Value                  string                                 `position:"Query" name:"Filter.2.Value"`
	AutoSnapshotPolicyId          string                                 `position:"Query" name:"AutoSnapshotPolicyId"`
	Tag3Key                       string                                 `position:"Query" name:"Tag.3.Key"`
	PageNumber                    int                                    `position:"Query" name:"PageNumber"`
	DiskName                      string                                 `position:"Query" name:"DiskName"`
	Tag1Value                     string                                 `position:"Query" name:"Tag.1.Value"`
	DeleteAutoSnapshot            string                                 `position:"Query" name:"DeleteAutoSnapshot"`
	ResourceGroupId               string                                 `position:"Query" name:"ResourceGroupId"`
	DiskChargeType                string                                 `position:"Query" name:"DiskChargeType"`
	LockReason                    string                                 `position:"Query" name:"LockReason"`
	Filter1Key                    string                                 `position:"Query" name:"Filter.1.Key"`
	PageSize                      int                                    `position:"Query" name:"PageSize"`
	DiskIds                       string                                 `position:"Query" name:"DiskIds"`
	DeleteWithInstance            string                                 `position:"Query" name:"DeleteWithInstance"`
	Tag3Value                     string                                 `position:"Query" name:"Tag.3.Value"`
	EnableAutoSnapshot            string                                 `position:"Query" name:"EnableAutoSnapshot"`
	Tag5Key                       string                                 `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount          string                                 `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                  string                                 `position:"Query" name:"OwnerAccount"`
	Filter1Value                  string                                 `position:"Query" name:"Filter.1.Value"`
	Portable                      string                                 `position:"Query" name:"Portable"`
	EnableAutomatedSnapshotPolicy string                                 `position:"Query" name:"EnableAutomatedSnapshotPolicy"`
	Filter2Key                    string                                 `position:"Query" name:"Filter.2.Key"`
	OwnerId                       int64                                  `position:"Query" name:"OwnerId"`
	DiskType                      string                                 `position:"Query" name:"DiskType"`
	Tag5Value                     string                                 `position:"Query" name:"Tag.5.Value"`
	Tag1Key                       string                                 `position:"Query" name:"Tag.1.Key"`
	AdditionalAttributess         *DescribeDisksAdditionalAttributesList `position:"Query" type:"Repeated" name:"AdditionalAttributes"`
	EnableShared                  string                                 `position:"Query" name:"EnableShared"`
	InstanceId                    string                                 `position:"Query" name:"InstanceId"`
	Encrypted                     string                                 `position:"Query" name:"Encrypted"`
	Tag2Value                     string                                 `position:"Query" name:"Tag.2.Value"`
	ZoneId                        string                                 `position:"Query" name:"ZoneId"`
	Tag4Key                       string                                 `position:"Query" name:"Tag.4.Key"`
	Category                      string                                 `position:"Query" name:"Category"`
	Status                        string                                 `position:"Query" name:"Status"`
}

func (r DescribeDisksRequest) Invoke(client *sdk.Client) (response *DescribeDisksResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDisksRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeDisks", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDisksResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDisksResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDisksResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Disks      DescribeDisksDiskList
}

type DescribeDisksDisk struct {
	DiskId                        string
	RegionId                      string
	ZoneId                        string
	DiskName                      string
	Description                   string
	Type                          string
	Category                      string
	Size                          int
	ImageId                       string
	SourceSnapshotId              string
	AutoSnapshotPolicyId          string
	ProductCode                   string
	Portable                      bool
	Status                        string
	InstanceId                    string
	Device                        string
	DeleteWithInstance            bool
	DeleteAutoSnapshot            bool
	EnableAutoSnapshot            bool
	EnableAutomatedSnapshotPolicy bool
	CreationTime                  string
	AttachedTime                  string
	DetachedTime                  string
	DiskChargeType                string
	ExpiredTime                   string
	ResourceGroupId               string
	Encrypted                     bool
	IOPS                          int
	IOPSRead                      int
	IOPSWrite                     int
	OperationLocks                DescribeDisksOperationLockList
	Tags                          DescribeDisksTagList
}

type DescribeDisksOperationLock struct {
	LockReason string
}

type DescribeDisksTag struct {
	TagKey   string
	TagValue string
}

type DescribeDisksAdditionalAttributesList []int64

func (list *DescribeDisksAdditionalAttributesList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type DescribeDisksDiskList []DescribeDisksDisk

func (list *DescribeDisksDiskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDisksDisk)
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

type DescribeDisksOperationLockList []DescribeDisksOperationLock

func (list *DescribeDisksOperationLockList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDisksOperationLock)
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

type DescribeDisksTagList []DescribeDisksTag

func (list *DescribeDisksTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDisksTag)
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
