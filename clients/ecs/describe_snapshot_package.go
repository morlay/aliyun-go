package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSnapshotPackageRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeSnapshotPackageRequest) Invoke(client *sdk.Client) (resp *DescribeSnapshotPackageResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSnapshotPackage", "ecs", "")
	resp = &DescribeSnapshotPackageResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSnapshotPackageResponse struct {
	responses.BaseResponse
	RequestId        string
	TotalCount       int
	PageNumber       int
	PageSize         int
	SnapshotPackages DescribeSnapshotPackageSnapshotPackageList
}

type DescribeSnapshotPackageSnapshotPackage struct {
	StartTime    string
	EndTime      string
	InitCapacity int64
	DisplayName  string
}

type DescribeSnapshotPackageSnapshotPackageList []DescribeSnapshotPackageSnapshotPackage

func (list *DescribeSnapshotPackageSnapshotPackageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSnapshotPackageSnapshotPackage)
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
