package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRepoWebhookRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

func (req *GetRepoWebhookRequest) Invoke(client *sdk.Client) (resp *GetRepoWebhookResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "GetRepoWebhook", "/repos/[RepoNamespace]/[RepoName]/webhooks", "", "")
	req.Method = "GET"

	resp = &GetRepoWebhookResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetRepoWebhookResponse struct {
	responses.BaseResponse
}
