package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteFlowCategoryRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *DeleteFlowCategoryRequest) Invoke(client *sdk.Client) (resp *DeleteFlowCategoryResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DeleteFlowCategory", "", "")
	resp = &DeleteFlowCategoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteFlowCategoryResponse struct {
	responses.BaseResponse
	RequestId common.String
	Data      bool
}
