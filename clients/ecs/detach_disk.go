package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DetachDiskRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DiskId               string `position:"Query" name:"DiskId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DetachDiskRequest) Invoke(client *sdk.Client) (response *DetachDiskResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DetachDiskRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DetachDisk", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DetachDiskResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DetachDiskResponse

	err = client.DoAction(&req, &resp)
	return
}

type DetachDiskResponse struct {
	RequestId string
}
