package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteScheduledTaskRequest struct {
	requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ScheduledTaskId      string `position:"Query" name:"ScheduledTaskId"`
}

func (req *DeleteScheduledTaskRequest) Invoke(client *sdk.Client) (resp *DeleteScheduledTaskResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DeleteScheduledTask", "ess", "")
	resp = &DeleteScheduledTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteScheduledTaskResponse struct {
	responses.BaseResponse
	RequestId common.String
}
