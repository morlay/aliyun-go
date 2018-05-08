package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyScheduledTaskRequest struct {
	requests.RpcRequest
	LaunchTime           string `position:"Query" name:"LaunchTime"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
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
	ScheduledTaskId      string `position:"Query" name:"ScheduledTaskId"`
	RecurrenceType       string `position:"Query" name:"RecurrenceType"`
}

func (req *ModifyScheduledTaskRequest) Invoke(client *sdk.Client) (resp *ModifyScheduledTaskResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "ModifyScheduledTask", "ess", "")
	resp = &ModifyScheduledTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyScheduledTaskResponse struct {
	responses.BaseResponse
	RequestId common.String
}
