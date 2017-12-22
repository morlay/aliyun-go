package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitAIVideoSummaryJobRequest struct {
	requests.RpcRequest
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	AIVideoSummaryConfig string `position:"Query" name:"AIVideoSummaryConfig"`
}

func (req *SubmitAIVideoSummaryJobRequest) Invoke(client *sdk.Client) (resp *SubmitAIVideoSummaryJobResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "SubmitAIVideoSummaryJob", "vod", "")
	resp = &SubmitAIVideoSummaryJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitAIVideoSummaryJobResponse struct {
	responses.BaseResponse
	RequestId         string
	AIVideoSummaryJob SubmitAIVideoSummaryJobAIVideoSummaryJob
}

type SubmitAIVideoSummaryJobAIVideoSummaryJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
