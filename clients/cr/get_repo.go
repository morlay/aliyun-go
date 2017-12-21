package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRepoRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

func (req *GetRepoRequest) Invoke(client *sdk.Client) (resp *GetRepoResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "GetRepo", "/repos/[RepoNamespace]/[RepoName]", "", "")
	req.Method = "GET"

	resp = &GetRepoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetRepoResponse struct {
	responses.BaseResponse
}
