package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyUserBusinessBehaviorRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	StatusValue          string `position:"Query" name:"StatusValue"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	StatusKey            string `position:"Query" name:"StatusKey"`
}

func (req *ModifyUserBusinessBehaviorRequest) Invoke(client *sdk.Client) (resp *ModifyUserBusinessBehaviorResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyUserBusinessBehavior", "ecs", "")
	resp = &ModifyUserBusinessBehaviorResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyUserBusinessBehaviorResponse struct {
	responses.BaseResponse
	RequestId string
}
