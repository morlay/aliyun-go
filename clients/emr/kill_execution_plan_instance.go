package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type KillExecutionPlanInstanceRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (r KillExecutionPlanInstanceRequest) Invoke(client *sdk.Client) (response *KillExecutionPlanInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		KillExecutionPlanInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Emr", "2016-04-08", "KillExecutionPlanInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		KillExecutionPlanInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.KillExecutionPlanInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type KillExecutionPlanInstanceResponse struct {
	RequestId string
}
