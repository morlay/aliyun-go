package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyFlowCategoryRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	ParentId        string `position:"Query" name:"ParentId"`
}

func (req *ModifyFlowCategoryRequest) Invoke(client *sdk.Client) (resp *ModifyFlowCategoryResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyFlowCategory", "", "")
	resp = &ModifyFlowCategoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyFlowCategoryResponse struct {
	responses.BaseResponse
	RequestId string
	Data      bool
}
