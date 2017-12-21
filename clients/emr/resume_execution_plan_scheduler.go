package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResumeExecutionPlanSchedulerRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (r ResumeExecutionPlanSchedulerRequest) Invoke(client *sdk.Client) (response *ResumeExecutionPlanSchedulerResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ResumeExecutionPlanSchedulerRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Emr", "2016-04-08", "ResumeExecutionPlanScheduler", "", "")

	resp := struct {
		*responses.BaseResponse
		ResumeExecutionPlanSchedulerResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ResumeExecutionPlanSchedulerResponse

	err = client.DoAction(&req, &resp)
	return
}

type ResumeExecutionPlanSchedulerResponse struct {
	RequestId string
}
