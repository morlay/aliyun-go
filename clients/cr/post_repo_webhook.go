package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PostRepoWebhookRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	WebhookId     int64  `position:"Path" name:"WebhookId"`
}

func (req *PostRepoWebhookRequest) Invoke(client *sdk.Client) (resp *PostRepoWebhookResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "PostRepoWebhook", "/repos/[RepoNamespace]/[RepoName]/webhooks/[WebhookId]", "", "")
	req.Method = "POST"

	resp = &PostRepoWebhookResponse{}
	err = client.DoAction(req, resp)
	return
}

type PostRepoWebhookResponse struct {
	responses.BaseResponse
}
