package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RunExecutionPlanRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (r RunExecutionPlanRequest) Invoke(client *sdk.Client) (response *RunExecutionPlanResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RunExecutionPlanRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Emr", "2016-04-08", "RunExecutionPlan", "", "")

	resp := struct {
		*responses.BaseResponse
		RunExecutionPlanResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RunExecutionPlanResponse

	err = client.DoAction(&req, &resp)
	return
}

type RunExecutionPlanResponse struct {
	RequestId               string
	ExecutionPlanInstanceId string
}
