package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SubmitAIVideoCensorJobRequest struct {
	requests.RpcRequest
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	AIVideoCensorConfig  string `position:"Query" name:"AIVideoCensorConfig"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
}

func (req *SubmitAIVideoCensorJobRequest) Invoke(client *sdk.Client) (resp *SubmitAIVideoCensorJobResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "SubmitAIVideoCensorJob", "vod", "")
	resp = &SubmitAIVideoCensorJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitAIVideoCensorJobResponse struct {
	responses.BaseResponse
	RequestId        common.String
	AIVideoCensorJob SubmitAIVideoCensorJobAIVideoCensorJob
}

type SubmitAIVideoCensorJobAIVideoCensorJob struct {
	JobId        common.String
	MediaId      common.String
	Status       common.String
	Code         common.String
	Message      common.String
	CreationTime common.String
	Data         common.String
}
