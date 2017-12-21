package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyInstanceAutoReleaseTimeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AutoReleaseTime      string `position:"Query" name:"AutoReleaseTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyInstanceAutoReleaseTimeRequest) Invoke(client *sdk.Client) (resp *ModifyInstanceAutoReleaseTimeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceAutoReleaseTime", "ecs", "")
	resp = &ModifyInstanceAutoReleaseTimeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyInstanceAutoReleaseTimeResponse struct {
	responses.BaseResponse
	RequestId string
}
