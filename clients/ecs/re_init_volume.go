package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReInitVolumeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Password             string `position:"Query" name:"Password"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VolumeId             string `position:"Query" name:"VolumeId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ReInitVolumeRequest) Invoke(client *sdk.Client) (resp *ReInitVolumeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ReInitVolume", "ecs", "")
	resp = &ReInitVolumeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReInitVolumeResponse struct {
	responses.BaseResponse
	RequestId string
}
