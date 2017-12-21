package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRepoBuildLogsRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	BuildId       string `position:"Path" name:"BuildId"`
}

func (req *GetRepoBuildLogsRequest) Invoke(client *sdk.Client) (resp *GetRepoBuildLogsResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "GetRepoBuildLogs", "/repos/[RepoNamespace]/[RepoName]/build/[BuildId]/logs", "", "")
	req.Method = "GET"

	resp = &GetRepoBuildLogsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetRepoBuildLogsResponse struct {
	responses.BaseResponse
}
