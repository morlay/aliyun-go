package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteRepoWebhookRequest struct {
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	WebhookId     int64  `position:"Path" name:"WebhookId"`
}

func (r DeleteRepoWebhookRequest) Invoke(client *sdk.Client) (response *DeleteRepoWebhookResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DeleteRepoWebhookRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "DeleteRepoWebhook", "/repos/[RepoNamespace]/[RepoName]/webhooks/[WebhookId]", "", "")
	req.Method = "DELETE"

	resp := struct {
		*responses.BaseResponse
		DeleteRepoWebhookResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteRepoWebhookResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteRepoWebhookResponse struct {
}
