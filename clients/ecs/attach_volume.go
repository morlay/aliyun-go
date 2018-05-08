package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AttachVolumeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VolumeId             string `position:"Query" name:"VolumeId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *AttachVolumeRequest) Invoke(client *sdk.Client) (resp *AttachVolumeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "AttachVolume", "ecs", "")
	resp = &AttachVolumeResponse{}
	err = client.DoAction(req, resp)
	return
}

type AttachVolumeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
