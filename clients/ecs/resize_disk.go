package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResizeDiskRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	NewSize              int    `position:"Query" name:"NewSize"`
	DiskId               string `position:"Query" name:"DiskId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ResizeDiskRequest) Invoke(client *sdk.Client) (response *ResizeDiskResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ResizeDiskRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ResizeDisk", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ResizeDiskResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ResizeDiskResponse

	err = client.DoAction(&req, &resp)
	return
}

type ResizeDiskResponse struct {
	RequestId string
}
