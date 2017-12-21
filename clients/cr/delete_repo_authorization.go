package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteRepoAuthorizationRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	AuthorizeId   int64  `position:"Path" name:"AuthorizeId"`
}

func (req *DeleteRepoAuthorizationRequest) Invoke(client *sdk.Client) (resp *DeleteRepoAuthorizationResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "DeleteRepoAuthorization", "/repos/[RepoNamespace]/[RepoName]/authorizations/[AuthorizeId]", "", "")
	req.Method = "DELETE"

	resp = &DeleteRepoAuthorizationResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteRepoAuthorizationResponse struct {
	responses.BaseResponse
}
