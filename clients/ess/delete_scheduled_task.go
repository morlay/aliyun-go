package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteScheduledTaskRequest struct {
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ScheduledTaskId      string `position:"Query" name:"ScheduledTaskId"`
}

func (r DeleteScheduledTaskRequest) Invoke(client *sdk.Client) (response *DeleteScheduledTaskResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteScheduledTaskRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ess", "2014-08-28", "DeleteScheduledTask", "ess", "")

	resp := struct {
		*responses.BaseResponse
		DeleteScheduledTaskResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteScheduledTaskResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteScheduledTaskResponse struct {
	RequestId string
}
