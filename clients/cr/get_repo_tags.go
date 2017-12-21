package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRepoTagsRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	Page          int    `position:"Query" name:"Page"`
	PageSize      int    `position:"Query" name:"PageSize"`
}

func (req *GetRepoTagsRequest) Invoke(client *sdk.Client) (resp *GetRepoTagsResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "GetRepoTags", "/repos/[RepoNamespace]/[RepoName]/tags", "", "")
	req.Method = "GET"

	resp = &GetRepoTagsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetRepoTagsResponse struct {
	responses.BaseResponse
}
