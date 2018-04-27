package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyFlowProjectRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	Description     string `position:"Query" name:"Description"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *ModifyFlowProjectRequest) Invoke(client *sdk.Client) (resp *ModifyFlowProjectResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyFlowProject", "", "")
	resp = &ModifyFlowProjectResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyFlowProjectResponse struct {
	responses.BaseResponse
	RequestId string
	Data      bool
}
