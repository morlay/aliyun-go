package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId            string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	TerminateSubscription string `position:"Query" name:"TerminateSubscription"`
	Force                 string `position:"Query" name:"Force"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteInstanceRequest) Invoke(client *sdk.Client) (resp *DeleteInstanceResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteInstance", "ecs", "")
	resp = &DeleteInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteInstanceResponse struct {
	responses.BaseResponse
	RequestId string
}
