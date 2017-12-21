package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type KillExecutionJobInstanceRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	JobInstanceId   string `position:"Query" name:"JobInstanceId"`
}

func (r KillExecutionJobInstanceRequest) Invoke(client *sdk.Client) (response *KillExecutionJobInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		KillExecutionJobInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Emr", "2016-04-08", "KillExecutionJobInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		KillExecutionJobInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.KillExecutionJobInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type KillExecutionJobInstanceResponse struct {
	RequestId string
}
