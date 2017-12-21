package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRepoRequest struct {
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

func (r GetRepoRequest) Invoke(client *sdk.Client) (response *GetRepoResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetRepoRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "GetRepo", "/repos/[RepoNamespace]/[RepoName]", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetRepoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetRepoResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetRepoResponse struct {
}
