package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteRepoRequest struct {
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

func (r DeleteRepoRequest) Invoke(client *sdk.Client) (response *DeleteRepoResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DeleteRepoRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "DeleteRepo", "/repos/[RepoNamespace]/[RepoName]", "", "")
	req.Method = "DELETE"

	resp := struct {
		*responses.BaseResponse
		DeleteRepoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteRepoResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteRepoResponse struct {
}
