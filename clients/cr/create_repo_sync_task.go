package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateRepoSyncTaskRequest struct {
	requests.RoaRequest
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

func (req *CreateRepoSyncTaskRequest) Invoke(client *sdk.Client) (resp *CreateRepoSyncTaskResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "CreateRepoSyncTask", "/repos/[RepoNamespace]/[RepoName]/syncTasks", "", "")
	req.Method = "PUT"

	resp = &CreateRepoSyncTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateRepoSyncTaskResponse struct {
	responses.BaseResponse
}
