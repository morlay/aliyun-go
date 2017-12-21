package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateRepoWebhookRequest struct {
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

func (r CreateRepoWebhookRequest) Invoke(client *sdk.Client) (response *CreateRepoWebhookResponse, err error) {
	req := struct {
		*requests.RoaRequest
		CreateRepoWebhookRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "CreateRepoWebhook", "/repos/[RepoNamespace]/[RepoName]/webhooks", "", "")
	req.Method = "PUT"

	resp := struct {
		*responses.BaseResponse
		CreateRepoWebhookResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateRepoWebhookResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateRepoWebhookResponse struct {
}
