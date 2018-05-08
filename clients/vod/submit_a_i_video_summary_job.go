package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId         common.String
	AIVideoSummaryJob SubmitAIVideoSummaryJobAIVideoSummaryJob
}

type SubmitAIVideoSummaryJobAIVideoSummaryJob struct {
	JobId        common.String
	MediaId      common.String
	Status       common.String
	Code         common.String
	Message      common.String
	CreationTime common.String
	Data         common.String
}
