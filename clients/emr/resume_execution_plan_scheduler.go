package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ResumeExecutionPlanSchedulerRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *ResumeExecutionPlanSchedulerRequest) Invoke(client *sdk.Client) (resp *ResumeExecutionPlanSchedulerResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ResumeExecutionPlanScheduler", "", "")
	resp = &ResumeExecutionPlanSchedulerResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResumeExecutionPlanSchedulerResponse struct {
	responses.BaseResponse
	RequestId common.String
}
