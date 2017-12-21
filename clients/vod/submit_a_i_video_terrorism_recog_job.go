package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitAIVideoTerrorismRecogJobRequest struct {
	requests.RpcRequest
	UserData                    string `position:"Query" name:"UserData"`
	ResourceOwnerId             string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount        string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string `position:"Query" name:"OwnerAccount"`
	AIVideoTerrorismRecogConfig string `position:"Query" name:"AIVideoTerrorismRecogConfig"`
	OwnerId                     string `position:"Query" name:"OwnerId"`
	MediaId                     string `position:"Query" name:"MediaId"`
}

func (req *SubmitAIVideoTerrorismRecogJobRequest) Invoke(client *sdk.Client) (resp *SubmitAIVideoTerrorismRecogJobResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "SubmitAIVideoTerrorismRecogJob", "", "")
	resp = &SubmitAIVideoTerrorismRecogJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitAIVideoTerrorismRecogJobResponse struct {
	responses.BaseResponse
	RequestId                string
	AIVideoTerrorismRecogJob SubmitAIVideoTerrorismRecogJobAIVideoTerrorismRecogJob
}

type SubmitAIVideoTerrorismRecogJobAIVideoTerrorismRecogJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
