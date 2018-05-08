package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CancelTaskRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TaskId               string `position:"Query" name:"TaskId"`
}

func (req *CancelTaskRequest) Invoke(client *sdk.Client) (resp *CancelTaskResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "CancelTask", "ecs", "")
	resp = &CancelTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type CancelTaskResponse struct {
	responses.BaseResponse
	RequestId common.String
}
