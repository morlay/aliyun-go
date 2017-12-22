package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitAIVideoCategoryJobRequest struct {
	requests.RpcRequest
	AIVideoCategoryConfig string `position:"Query" name:"AIVideoCategoryConfig"`
	UserData              string `position:"Query" name:"UserData"`
	ResourceOwnerId       string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               string `position:"Query" name:"OwnerId"`
	MediaId               string `position:"Query" name:"MediaId"`
}

func (req *SubmitAIVideoCategoryJobRequest) Invoke(client *sdk.Client) (resp *SubmitAIVideoCategoryJobResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "SubmitAIVideoCategoryJob", "vod", "")
	resp = &SubmitAIVideoCategoryJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitAIVideoCategoryJobResponse struct {
	responses.BaseResponse
	RequestId          string
	AIVideoCategoryJob SubmitAIVideoCategoryJobAIVideoCategoryJob
}

type SubmitAIVideoCategoryJobAIVideoCategoryJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
