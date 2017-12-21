package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateRepoWebhookRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

func (req *CreateRepoWebhookRequest) Invoke(client *sdk.Client) (resp *CreateRepoWebhookResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "CreateRepoWebhook", "/repos/[RepoNamespace]/[RepoName]/webhooks", "", "")
	req.Method = "PUT"

	resp = &CreateRepoWebhookResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateRepoWebhookResponse struct {
	responses.BaseResponse
}
