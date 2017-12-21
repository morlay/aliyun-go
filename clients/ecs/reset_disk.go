package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResetDiskRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SnapshotId           string `position:"Query" name:"SnapshotId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DiskId               string `position:"Query" name:"DiskId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ResetDiskRequest) Invoke(client *sdk.Client) (response *ResetDiskResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ResetDiskRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ResetDisk", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ResetDiskResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ResetDiskResponse

	err = client.DoAction(&req, &resp)
	return
}

type ResetDiskResponse struct {
	RequestId string
}
