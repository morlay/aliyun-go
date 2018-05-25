package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateRepoRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

func (req *UpdateRepoRequest) Invoke(client *sdk.Client) (resp *UpdateRepoResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "UpdateRepo", "/repos/[RepoNamespace]/[RepoName]", "", "")
	req.Method = "POST"

	resp = &UpdateRepoResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateRepoResponse struct {
	responses.BaseResponse
}
