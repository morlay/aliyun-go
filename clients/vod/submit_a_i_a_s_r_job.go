package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SubmitAIASRJobRequest struct {
	requests.RpcRequest
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AIASRConfig          string `position:"Query" name:"AIASRConfig"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
}

func (req *SubmitAIASRJobRequest) Invoke(client *sdk.Client) (resp *SubmitAIASRJobResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "SubmitAIASRJob", "vod", "")
	resp = &SubmitAIASRJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitAIASRJobResponse struct {
	responses.BaseResponse
	RequestId common.String
	AIASRJob  SubmitAIASRJobAIASRJob
}

type SubmitAIASRJobAIASRJob struct {
	JobId        common.String
	MediaId      common.String
	Status       common.String
	Code         common.String
	Message      common.String
	CreationTime common.String
	Data         common.String
}
