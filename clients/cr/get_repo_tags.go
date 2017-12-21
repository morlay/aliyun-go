package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRepoTagsRequest struct {
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	Page          int    `position:"Query" name:"Page"`
	PageSize      int    `position:"Query" name:"PageSize"`
}

func (r GetRepoTagsRequest) Invoke(client *sdk.Client) (response *GetRepoTagsResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetRepoTagsRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "GetRepoTags", "/repos/[RepoNamespace]/[RepoName]/tags", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetRepoTagsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetRepoTagsResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetRepoTagsResponse struct {
}
