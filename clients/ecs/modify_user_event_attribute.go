package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyUserEventAttributeRequest struct {
	requests.RpcRequest
	EventId              string `position:"Query" name:"EventId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	NewPlanTime          string `position:"Query" name:"NewPlanTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NewExpireTime        string `position:"Query" name:"NewExpireTime"`
}

func (req *ModifyUserEventAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyUserEventAttributeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyUserEventAttribute", "ecs", "")
	resp = &ModifyUserEventAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyUserEventAttributeResponse struct {
	responses.BaseResponse
	RequestId string
	EventId   string
}
