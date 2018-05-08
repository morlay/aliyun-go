package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DetachVolumeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VolumeId             string `position:"Query" name:"VolumeId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DetachVolumeRequest) Invoke(client *sdk.Client) (resp *DetachVolumeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DetachVolume", "ecs", "")
	resp = &DetachVolumeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DetachVolumeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
