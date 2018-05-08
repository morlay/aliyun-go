package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateJobExecutionPlanFolderRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	ParentId        int64  `position:"Query" name:"ParentId"`
}

func (req *CreateJobExecutionPlanFolderRequest) Invoke(client *sdk.Client) (resp *CreateJobExecutionPlanFolderResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateJobExecutionPlanFolder", "", "")
	resp = &CreateJobExecutionPlanFolderResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateJobExecutionPlanFolderResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   common.String
	ErrCode   common.String
	ErrMsg    common.String
	FolderId  common.String
	FolderId1 common.String
}
