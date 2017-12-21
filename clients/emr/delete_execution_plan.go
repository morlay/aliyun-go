package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteExecutionPlanRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *DeleteExecutionPlanRequest) Invoke(client *sdk.Client) (resp *DeleteExecutionPlanResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DeleteExecutionPlan", "", "")
	resp = &DeleteExecutionPlanResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteExecutionPlanResponse struct {
	responses.BaseResponse
	RequestId string
}
