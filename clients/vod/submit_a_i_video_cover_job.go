package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitAIVideoCoverJobRequest struct {
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	AIVideoCoverConfig   string `position:"Query" name:"AIVideoCoverConfig"`
}

func (r SubmitAIVideoCoverJobRequest) Invoke(client *sdk.Client) (response *SubmitAIVideoCoverJobResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SubmitAIVideoCoverJobRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "SubmitAIVideoCoverJob", "", "")

	resp := struct {
		*responses.BaseResponse
		SubmitAIVideoCoverJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SubmitAIVideoCoverJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type SubmitAIVideoCoverJobResponse struct {
	RequestId       string
	AIVideoCoverJob SubmitAIVideoCoverJobAIVideoCoverJob
}

type SubmitAIVideoCoverJobAIVideoCoverJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
