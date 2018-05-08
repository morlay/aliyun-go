package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SubmitAIVideoPornRecogJobRequest struct {
	requests.RpcRequest
	UserData               string `position:"Query" name:"UserData"`
	ResourceOwnerId        string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	AIVideoPornRecogConfig string `position:"Query" name:"AIVideoPornRecogConfig"`
	OwnerId                string `position:"Query" name:"OwnerId"`
	MediaId                string `position:"Query" name:"MediaId"`
}

func (req *SubmitAIVideoPornRecogJobRequest) Invoke(client *sdk.Client) (resp *SubmitAIVideoPornRecogJobResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "SubmitAIVideoPornRecogJob", "vod", "")
	resp = &SubmitAIVideoPornRecogJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitAIVideoPornRecogJobResponse struct {
	responses.BaseResponse
	RequestId           common.String
	AIVideoPornRecogJob SubmitAIVideoPornRecogJobAIVideoPornRecogJob
}

type SubmitAIVideoPornRecogJobAIVideoPornRecogJob struct {
	JobId        common.String
	MediaId      common.String
	Status       common.String
	Code         common.String
	Message      common.String
	CreationTime common.String
	Data         common.String
}
