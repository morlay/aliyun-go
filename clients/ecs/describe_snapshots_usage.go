package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSnapshotsUsageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeSnapshotsUsageRequest) Invoke(client *sdk.Client) (response *DescribeSnapshotsUsageResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeSnapshotsUsageRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSnapshotsUsage", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeSnapshotsUsageResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeSnapshotsUsageResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeSnapshotsUsageResponse struct {
	RequestId     string
	SnapshotCount int
	SnapshotSize  int64
}
