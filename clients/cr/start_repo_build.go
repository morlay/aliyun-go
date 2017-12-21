package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StartRepoBuildRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

func (req *StartRepoBuildRequest) Invoke(client *sdk.Client) (resp *StartRepoBuildResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "StartRepoBuild", "/repos/[RepoNamespace]/[RepoName]/build", "", "")
	req.Method = "PUT"

	resp = &StartRepoBuildResponse{}
	err = client.DoAction(req, resp)
	return
}

type StartRepoBuildResponse struct {
	responses.BaseResponse
}
