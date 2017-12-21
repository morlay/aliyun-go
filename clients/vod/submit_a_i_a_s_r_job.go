package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitAIASRJobRequest struct {
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AIASRConfig          string `position:"Query" name:"AIASRConfig"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
}

func (r SubmitAIASRJobRequest) Invoke(client *sdk.Client) (response *SubmitAIASRJobResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SubmitAIASRJobRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "SubmitAIASRJob", "", "")

	resp := struct {
		*responses.BaseResponse
		SubmitAIASRJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SubmitAIASRJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type SubmitAIASRJobResponse struct {
	RequestId string
	AIASRJob  SubmitAIASRJobAIASRJob
}

type SubmitAIASRJobAIASRJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
