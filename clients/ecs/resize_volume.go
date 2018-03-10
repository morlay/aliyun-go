package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResizeVolumeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VolumeId             string `position:"Query" name:"VolumeId"`
	NewSize              int    `position:"Query" name:"NewSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ResizeVolumeRequest) Invoke(client *sdk.Client) (resp *ResizeVolumeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ResizeVolume", "ecs", "")
	resp = &ResizeVolumeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResizeVolumeResponse struct {
	responses.BaseResponse
	RequestId string
}
