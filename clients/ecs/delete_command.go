package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCommandRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	CommandId            string `position:"Query" name:"CommandId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteCommandRequest) Invoke(client *sdk.Client) (resp *DeleteCommandResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteCommand", "ecs", "")
	resp = &DeleteCommandResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCommandResponse struct {
	responses.BaseResponse
	RequestId string
}
