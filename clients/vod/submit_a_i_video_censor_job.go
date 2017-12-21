package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitAIVideoCensorJobRequest struct {
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	AIVideoCensorConfig  string `position:"Query" name:"AIVideoCensorConfig"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
}

func (r SubmitAIVideoCensorJobRequest) Invoke(client *sdk.Client) (response *SubmitAIVideoCensorJobResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SubmitAIVideoCensorJobRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "SubmitAIVideoCensorJob", "", "")

	resp := struct {
		*responses.BaseResponse
		SubmitAIVideoCensorJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SubmitAIVideoCensorJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type SubmitAIVideoCensorJobResponse struct {
	RequestId        string
	AIVideoCensorJob SubmitAIVideoCensorJobAIVideoCensorJob
}

type SubmitAIVideoCensorJobAIVideoCensorJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}
