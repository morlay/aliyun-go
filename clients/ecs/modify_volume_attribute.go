package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyVolumeAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	VolumeName           string `position:"Query" name:"VolumeName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VolumeId             string `position:"Query" name:"VolumeId"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyVolumeAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyVolumeAttributeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyVolumeAttribute", "ecs", "")
	resp = &ModifyVolumeAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyVolumeAttributeResponse struct {
	responses.BaseResponse
	RequestId string
}
