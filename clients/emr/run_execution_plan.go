package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RunExecutionPlanRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Arguments       string `position:"Query" name:"Arguments"`
	Id              string `position:"Query" name:"Id"`
}

func (req *RunExecutionPlanRequest) Invoke(client *sdk.Client) (resp *RunExecutionPlanResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "RunExecutionPlan", "", "")
	resp = &RunExecutionPlanResponse{}
	err = client.DoAction(req, resp)
	return
}

type RunExecutionPlanResponse struct {
	responses.BaseResponse
	RequestId               common.String
	ExecutionPlanInstanceId common.String
}
