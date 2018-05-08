package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RollbackVolumeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SnapshotId           string `position:"Query" name:"SnapshotId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VolumeId             string `position:"Query" name:"VolumeId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *RollbackVolumeRequest) Invoke(client *sdk.Client) (resp *RollbackVolumeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "RollbackVolume", "ecs", "")
	resp = &RollbackVolumeResponse{}
	err = client.DoAction(req, resp)
	return
}

type RollbackVolumeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
