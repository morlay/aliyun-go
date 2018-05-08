package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DetachDiskRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DiskId               string `position:"Query" name:"DiskId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DetachDiskRequest) Invoke(client *sdk.Client) (resp *DetachDiskResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DetachDisk", "ecs", "")
	resp = &DetachDiskResponse{}
	err = client.DoAction(req, resp)
	return
}

type DetachDiskResponse struct {
	responses.BaseResponse
	RequestId common.String
}
