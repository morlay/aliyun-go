package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeSnapshotsUsageRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeSnapshotsUsageRequest) Invoke(client *sdk.Client) (resp *DescribeSnapshotsUsageResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSnapshotsUsage", "ecs", "")
	resp = &DescribeSnapshotsUsageResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSnapshotsUsageResponse struct {
	responses.BaseResponse
	RequestId     common.String
	SnapshotCount common.Integer
	SnapshotSize  common.Long
}
