package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SuspendExecutionPlanSchedulerRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *SuspendExecutionPlanSchedulerRequest) Invoke(client *sdk.Client) (resp *SuspendExecutionPlanSchedulerResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "SuspendExecutionPlanScheduler", "", "")
	resp = &SuspendExecutionPlanSchedulerResponse{}
	err = client.DoAction(req, resp)
	return
}

type SuspendExecutionPlanSchedulerResponse struct {
	responses.BaseResponse
	RequestId string
}
