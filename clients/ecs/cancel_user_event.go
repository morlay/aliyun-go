package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CancelUserEventRequest struct {
	requests.RpcRequest
	EventId              string `position:"Query" name:"EventId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CancelUserEventRequest) Invoke(client *sdk.Client) (resp *CancelUserEventResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "CancelUserEvent", "ecs", "")
	resp = &CancelUserEventResponse{}
	err = client.DoAction(req, resp)
	return
}

type CancelUserEventResponse struct {
	responses.BaseResponse
	RequestId string
	EventId   string
}
