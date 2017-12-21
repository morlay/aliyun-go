package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteRepoWebhookRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	WebhookId     int64  `position:"Path" name:"WebhookId"`
}

func (req *DeleteRepoWebhookRequest) Invoke(client *sdk.Client) (resp *DeleteRepoWebhookResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "DeleteRepoWebhook", "/repos/[RepoNamespace]/[RepoName]/webhooks/[WebhookId]", "", "")
	req.Method = "DELETE"

	resp = &DeleteRepoWebhookResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteRepoWebhookResponse struct {
	responses.BaseResponse
}
