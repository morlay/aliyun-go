package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteJobExecutionPlanFolderRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
	Id              int64 `position:"Query" name:"Id"`
}

func (req *DeleteJobExecutionPlanFolderRequest) Invoke(client *sdk.Client) (resp *DeleteJobExecutionPlanFolderResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DeleteJobExecutionPlanFolder", "", "")
	resp = &DeleteJobExecutionPlanFolderResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteJobExecutionPlanFolderResponse struct {
	responses.BaseResponse
	RequestId string
	Success   string
	ErrCode   string
	ErrMsg    string
	FolderId  string
}
