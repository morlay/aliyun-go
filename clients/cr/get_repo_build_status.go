package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRepoBuildStatusRequest struct {
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	BuildId       string `position:"Path" name:"BuildId"`
}

func (r GetRepoBuildStatusRequest) Invoke(client *sdk.Client) (response *GetRepoBuildStatusResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetRepoBuildStatusRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "GetRepoBuildStatus", "/repos/[RepoNamespace]/[RepoName]/build/[BuildId]/status", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetRepoBuildStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetRepoBuildStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetRepoBuildStatusResponse struct {
}
