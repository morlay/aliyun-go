package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateScheduledTaskRequest struct {
	requests.RpcRequest
	LaunchTime           string `position:"Query" name:"LaunchTime"`
	ScheduledAction      string `position:"Query" name:"ScheduledAction"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RecurrenceValue      string `position:"Query" name:"RecurrenceValue"`
	LaunchExpirationTime int    `position:"Query" name:"LaunchExpirationTime"`
	RecurrenceEndTime    string `position:"Query" name:"RecurrenceEndTime"`
	ScheduledTaskName    string `position:"Query" name:"ScheduledTaskName"`
	TaskEnabled          string `position:"Query" name:"TaskEnabled"`
	RecurrenceType       string `position:"Query" name:"RecurrenceType"`
}

func (req *CreateScheduledTaskRequest) Invoke(client *sdk.Client) (resp *CreateScheduledTaskResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "CreateScheduledTask", "ess", "")
	resp = &CreateScheduledTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateScheduledTaskResponse struct {
	responses.BaseResponse
	ScheduledTaskId string
	RequestId       string
}
