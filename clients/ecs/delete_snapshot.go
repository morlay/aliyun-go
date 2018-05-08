package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteSnapshotRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SnapshotId           string `position:"Query" name:"SnapshotId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Force                string `position:"Query" name:"Force"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteSnapshotRequest) Invoke(client *sdk.Client) (resp *DeleteSnapshotResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteSnapshot", "ecs", "")
	resp = &DeleteSnapshotResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteSnapshotResponse struct {
	responses.BaseResponse
	RequestId common.String
}
