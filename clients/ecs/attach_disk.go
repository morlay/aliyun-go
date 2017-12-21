package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AttachDiskRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DiskId               string `position:"Query" name:"DiskId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Device               string `position:"Query" name:"Device"`
	DeleteWithInstance   string `position:"Query" name:"DeleteWithInstance"`
}

func (req *AttachDiskRequest) Invoke(client *sdk.Client) (resp *AttachDiskResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "AttachDisk", "ecs", "")
	resp = &AttachDiskResponse{}
	err = client.DoAction(req, resp)
	return
}

type AttachDiskResponse struct {
	responses.BaseResponse
	RequestId string
}
