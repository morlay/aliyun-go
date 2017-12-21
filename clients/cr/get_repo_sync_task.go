package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRepoSyncTaskRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
	SyncTaskId    string `position:"Path" name:"SyncTaskId"`
}

func (req *GetRepoSyncTaskRequest) Invoke(client *sdk.Client) (resp *GetRepoSyncTaskResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "GetRepoSyncTask", "/repos/[RepoNamespace]/[RepoName]/syncTasks/[SyncTaskId]", "", "")
	req.Method = "GET"

	resp = &GetRepoSyncTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetRepoSyncTaskResponse struct {
	responses.BaseResponse
}
