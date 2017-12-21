package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteRepoAuthorizationRequest struct {
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	AuthorizeId   int64  `position:"Path" name:"AuthorizeId"`
}

func (r DeleteRepoAuthorizationRequest) Invoke(client *sdk.Client) (response *DeleteRepoAuthorizationResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DeleteRepoAuthorizationRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "DeleteRepoAuthorization", "/repos/[RepoNamespace]/[RepoName]/authorizations/[AuthorizeId]", "", "")
	req.Method = "DELETE"

	resp := struct {
		*responses.BaseResponse
		DeleteRepoAuthorizationResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteRepoAuthorizationResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteRepoAuthorizationResponse struct {
}
