package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SuspendExecutionPlanSchedulerRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (r SuspendExecutionPlanSchedulerRequest) Invoke(client *sdk.Client) (response *SuspendExecutionPlanSchedulerResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SuspendExecutionPlanSchedulerRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Emr", "2016-04-08", "SuspendExecutionPlanScheduler", "", "")

	resp := struct {
		*responses.BaseResponse
		SuspendExecutionPlanSchedulerResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SuspendExecutionPlanSchedulerResponse

	err = client.DoAction(&req, &resp)
	return
}

type SuspendExecutionPlanSchedulerResponse struct {
	RequestId string
}
