package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CancelTaskRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TaskId               string `position:"Query" name:"TaskId"`
}

func (r CancelTaskRequest) Invoke(client *sdk.Client) (response *CancelTaskResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CancelTaskRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CancelTask", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CancelTaskResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CancelTaskResponse

	err = client.DoAction(&req, &resp)
	return
}

type CancelTaskResponse struct {
	RequestId string
}
