package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRepoBuildLogsRequest struct {
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	BuildId       string `position:"Path" name:"BuildId"`
}

func (r GetRepoBuildLogsRequest) Invoke(client *sdk.Client) (response *GetRepoBuildLogsResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetRepoBuildLogsRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "GetRepoBuildLogs", "/repos/[RepoNamespace]/[RepoName]/build/[BuildId]/logs", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetRepoBuildLogsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetRepoBuildLogsResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetRepoBuildLogsResponse struct {
}
