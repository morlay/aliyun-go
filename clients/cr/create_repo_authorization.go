package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateRepoAuthorizationRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

func (req *CreateRepoAuthorizationRequest) Invoke(client *sdk.Client) (resp *CreateRepoAuthorizationResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "CreateRepoAuthorization", "/repos/[RepoNamespace]/[RepoName]/authorizations", "", "")
	req.Method = "PUT"

	resp = &CreateRepoAuthorizationResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateRepoAuthorizationResponse struct {
	responses.BaseResponse
}
