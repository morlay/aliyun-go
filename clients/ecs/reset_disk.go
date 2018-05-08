package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ResetDiskRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SnapshotId           string `position:"Query" name:"SnapshotId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DiskId               string `position:"Query" name:"DiskId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ResetDiskRequest) Invoke(client *sdk.Client) (resp *ResetDiskResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ResetDisk", "ecs", "")
	resp = &ResetDiskResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResetDiskResponse struct {
	responses.BaseResponse
	RequestId common.String
}
