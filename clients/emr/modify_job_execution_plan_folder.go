package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Success   common.String
	ErrCode   common.String
	ErrMsg    common.String
	FolderId  common.String
}
