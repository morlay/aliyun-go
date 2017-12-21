package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitAIVideoPornRecogJobRequest struct {
	UserData               string `position:"Query" name:"UserData"`
	ResourceOwnerId        string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	AIVideoPornRecogConfig string `position:"Query" name:"AIVideoPornRecogConfig"`
	OwnerId                string `position:"Query" name:"OwnerId"`
	MediaId                string `position:"Query" name:"MediaId"`
}

func (r SubmitAIVideoPornRecogJobRequest) Invoke(client *sdk.Client) (response *SubmitAIVideoPornRecogJobResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SubmitAIVideoPornRecogJobRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "SubmitAIVideoPornRecogJob", "", "")

	resp := struct {
		*responses.BaseResponse
		SubmitAIVideoPornRecogJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SubmitAIVideoPornRecogJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type SubmitAIVideoPornRecogJobResponse struct {
	RequestId           string
	AIVideoPornRecogJob SubmitAIVideoPornRecogJobAIVideoPornRecogJob
}

type SubmitAIVideoPornRecogJobAIVideoPornRecogJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
