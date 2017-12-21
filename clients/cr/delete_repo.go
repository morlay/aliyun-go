package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteRepoRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

func (req *DeleteRepoRequest) Invoke(client *sdk.Client) (resp *DeleteRepoResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "DeleteRepo", "/repos/[RepoNamespace]/[RepoName]", "", "")
	req.Method = "DELETE"

	resp = &DeleteRepoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteRepoResponse struct {
	responses.BaseResponse
}
