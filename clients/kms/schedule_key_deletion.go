package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ScheduleKeyDeletionRequest struct {
	KeyId               string `position:"Query" name:"KeyId"`
	PendingWindowInDays int    `position:"Query" name:"PendingWindowInDays"`
	STSToken            string `position:"Query" name:"STSToken"`
}

func (r ScheduleKeyDeletionRequest) Invoke(client *sdk.Client) (response *ScheduleKeyDeletionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ScheduleKeyDeletionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Kms", "2016-01-20", "ScheduleKeyDeletion", "kms", "")

	resp := struct {
		*responses.BaseResponse
		ScheduleKeyDeletionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ScheduleKeyDeletionResponse

	err = client.DoAction(&req, &resp)
	return
}

type ScheduleKeyDeletionResponse struct {
	RequestId string
}
