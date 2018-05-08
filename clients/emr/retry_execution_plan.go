package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RetryExecutionPlanRequest struct {
	requests.RpcRequest
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	ExecutionPlanWorkNodeIds string `position:"Query" name:"ExecutionPlanWorkNodeIds"`
	Id                       string `position:"Query" name:"Id"`
}

func (req *RetryExecutionPlanRequest) Invoke(client *sdk.Client) (resp *RetryExecutionPlanResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "RetryExecutionPlan", "", "")
	resp = &RetryExecutionPlanResponse{}
	err = client.DoAction(req, resp)
	return
}

type RetryExecutionPlanResponse struct {
	responses.BaseResponse
	RequestId               common.String
	ExecutionPlanInstanceId common.String
}
