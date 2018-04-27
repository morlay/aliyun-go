package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateUserEventRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	PlanTime             string `position:"Query" name:"PlanTime"`
	ExpireTime           string `position:"Query" name:"ExpireTime"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	EventType            string `position:"Query" name:"EventType"`
}

func (req *CreateUserEventRequest) Invoke(client *sdk.Client) (resp *CreateUserEventResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateUserEvent", "ecs", "")
	resp = &CreateUserEventResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateUserEventResponse struct {
	responses.BaseResponse
	RequestId string
	EventId   string
}
