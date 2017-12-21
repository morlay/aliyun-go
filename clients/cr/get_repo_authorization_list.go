package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRepoAuthorizationListRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

func (req *GetRepoAuthorizationListRequest) Invoke(client *sdk.Client) (resp *GetRepoAuthorizationListResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "GetRepoAuthorizationList", "/repos/[RepoNamespace]/[RepoName]/authorizations", "", "")
	req.Method = "GET"

	resp = &GetRepoAuthorizationListResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetRepoAuthorizationListResponse struct {
	responses.BaseResponse
}
