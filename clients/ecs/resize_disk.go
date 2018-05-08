package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ResizeDiskRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	NewSize              int    `position:"Query" name:"NewSize"`
	DiskId               string `position:"Query" name:"DiskId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ResizeDiskRequest) Invoke(client *sdk.Client) (resp *ResizeDiskResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ResizeDisk", "ecs", "")
	resp = &ResizeDiskResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResizeDiskResponse struct {
	responses.BaseResponse
	RequestId common.String
}
