package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteExecutionPlanRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (r DeleteExecutionPlanRequest) Invoke(client *sdk.Client) (response *DeleteExecutionPlanResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteExecutionPlanRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Emr", "2016-04-08", "DeleteExecutionPlan", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteExecutionPlanResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteExecutionPlanResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteExecutionPlanResponse struct {
	RequestId string
}
