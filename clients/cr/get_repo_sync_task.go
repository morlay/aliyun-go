package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRepoSyncTaskRequest struct {
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	SyncTaskId    string `position:"Path" name:"SyncTaskId"`
}

func (r GetRepoSyncTaskRequest) Invoke(client *sdk.Client) (response *GetRepoSyncTaskResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetRepoSyncTaskRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "GetRepoSyncTask", "/repos/[RepoNamespace]/[RepoName]/syncTasks/[SyncTaskId]", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetRepoSyncTaskResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetRepoSyncTaskResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetRepoSyncTaskResponse struct {
}
