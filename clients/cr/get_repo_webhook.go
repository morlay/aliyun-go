package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRepoWebhookRequest struct {
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

func (r GetRepoWebhookRequest) Invoke(client *sdk.Client) (response *GetRepoWebhookResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetRepoWebhookRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "GetRepoWebhook", "/repos/[RepoNamespace]/[RepoName]/webhooks", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetRepoWebhookResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetRepoWebhookResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetRepoWebhookResponse struct {
}
