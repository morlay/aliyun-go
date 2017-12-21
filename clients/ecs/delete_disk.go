package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDiskRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DiskId               string `position:"Query" name:"DiskId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteDiskRequest) Invoke(client *sdk.Client) (response *DeleteDiskResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteDiskRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteDisk", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DeleteDiskResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteDiskResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteDiskResponse struct {
	RequestId string
}
