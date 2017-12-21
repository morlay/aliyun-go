package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateRepoAuthorizationRequest struct {
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

func (r CreateRepoAuthorizationRequest) Invoke(client *sdk.Client) (response *CreateRepoAuthorizationResponse, err error) {
	req := struct {
		*requests.RoaRequest
		CreateRepoAuthorizationRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "CreateRepoAuthorization", "/repos/[RepoNamespace]/[RepoName]/authorizations", "", "")
	req.Method = "PUT"

	resp := struct {
		*responses.BaseResponse
		CreateRepoAuthorizationResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateRepoAuthorizationResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateRepoAuthorizationResponse struct {
}
