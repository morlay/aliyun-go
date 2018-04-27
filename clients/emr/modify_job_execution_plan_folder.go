package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyJobExecutionPlanFolderRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	Id              int64  `position:"Query" name:"Id"`
	ParentId        int64  `position:"Query" name:"ParentId"`
}

func (req *ModifyJobExecutionPlanFolderRequest) Invoke(client *sdk.Client) (resp *ModifyJobExecutionPlanFolderResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyJobExecutionPlanFolder", "", "")
	resp = &ModifyJobExecutionPlanFolderResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyJobExecutionPlanFolderResponse struct {
	responses.BaseResponse
	RequestId string
	Success   string
	ErrCode   string
	ErrMsg    string
	FolderId  string
}
