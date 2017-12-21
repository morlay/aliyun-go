package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateRepoSyncTaskRequest struct {
	RepoNamespace string `position:"Path" name:"RepoNamespace"`
	RepoName      string `position:"Path" name:"RepoName"`
}

func (r CreateRepoSyncTaskRequest) Invoke(client *sdk.Client) (response *CreateRepoSyncTaskResponse, err error) {
	req := struct {
		*requests.RoaRequest
		CreateRepoSyncTaskRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "CreateRepoSyncTask", "/repos/[RepoNamespace]/[RepoName]/syncTasks", "", "")
	req.Method = "PUT"

	resp := struct {
		*responses.BaseResponse
		CreateRepoSyncTaskResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateRepoSyncTaskResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateRepoSyncTaskResponse struct {
}
