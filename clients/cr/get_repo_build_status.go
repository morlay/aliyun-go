package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRepoBuildStatusRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	BuildId       string `position:"Path" name:"BuildId"`
}

func (req *GetRepoBuildStatusRequest) Invoke(client *sdk.Client) (resp *GetRepoBuildStatusResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "GetRepoBuildStatus", "/repos/[RepoNamespace]/[RepoName]/build/[BuildId]/status", "", "")
	req.Method = "GET"

	resp = &GetRepoBuildStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetRepoBuildStatusResponse struct {
	responses.BaseResponse
}
