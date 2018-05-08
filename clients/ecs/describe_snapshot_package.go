package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId        common.String
	TotalCount       common.Integer
	PageNumber       common.Integer
	PageSize         common.Integer
	SnapshotPackages DescribeSnapshotPackageSnapshotPackageList
}

type DescribeSnapshotPackageSnapshotPackage struct {
	StartTime    common.String
	EndTime      common.String
	InitCapacity common.Long
	DisplayName  common.String
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
