package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteVolumeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VolumeId             string `position:"Query" name:"VolumeId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteVolumeRequest) Invoke(client *sdk.Client) (resp *DeleteVolumeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteVolume", "ecs", "")
	resp = &DeleteVolumeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteVolumeResponse struct {
	responses.BaseResponse
	RequestId string
}
