package kms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ScheduleKeyDeletionRequest struct {
	requests.RpcRequest
	KeyId               string `position:"Query" name:"KeyId"`
	PendingWindowInDays int    `position:"Query" name:"PendingWindowInDays"`
	STSToken            string `position:"Query" name:"STSToken"`
}

func (req *ScheduleKeyDeletionRequest) Invoke(client *sdk.Client) (resp *ScheduleKeyDeletionResponse, err error) {
	req.InitWithApiInfo("Kms", "2016-01-20", "ScheduleKeyDeletion", "kms", "")
	resp = &ScheduleKeyDeletionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ScheduleKeyDeletionResponse struct {
	responses.BaseResponse
	RequestId string
}
