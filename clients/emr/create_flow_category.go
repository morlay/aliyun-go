package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateFlowCategoryRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	Type            string `position:"Query" name:"Type"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	ParentId        string `position:"Query" name:"ParentId"`
}

func (req *CreateFlowCategoryRequest) Invoke(client *sdk.Client) (resp *CreateFlowCategoryResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateFlowCategory", "", "")
	resp = &CreateFlowCategoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateFlowCategoryResponse struct {
	responses.BaseResponse
	RequestId common.String
	Id        common.String
}
