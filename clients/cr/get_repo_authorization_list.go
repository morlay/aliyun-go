package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRepoAuthorizationListRequest struct {
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

func (r GetRepoAuthorizationListRequest) Invoke(client *sdk.Client) (response *GetRepoAuthorizationListResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetRepoAuthorizationListRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "GetRepoAuthorizationList", "/repos/[RepoNamespace]/[RepoName]/authorizations", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetRepoAuthorizationListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetRepoAuthorizationListResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetRepoAuthorizationListResponse struct {
}
