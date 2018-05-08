package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId          common.String
	AIVideoCategoryJob SubmitAIVideoCategoryJobAIVideoCategoryJob
}

type SubmitAIVideoCategoryJobAIVideoCategoryJob struct {
	JobId        common.String
	MediaId      common.String
	Status       common.String
	Code         common.String
	Message      common.String
	CreationTime common.String
	Data         common.String
}
