package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteDiskRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DiskId               string `position:"Query" name:"DiskId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteDiskRequest) Invoke(client *sdk.Client) (resp *DeleteDiskResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteDisk", "ecs", "")
	resp = &DeleteDiskResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteDiskResponse struct {
	responses.BaseResponse
	RequestId common.String
}
