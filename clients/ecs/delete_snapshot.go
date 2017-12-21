package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteSnapshotRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SnapshotId           string `position:"Query" name:"SnapshotId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Force                string `position:"Query" name:"Force"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteSnapshotRequest) Invoke(client *sdk.Client) (response *DeleteSnapshotResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteSnapshotRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteSnapshot", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DeleteSnapshotResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteSnapshotResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteSnapshotResponse struct {
	RequestId string
}
