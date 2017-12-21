package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type TerminatePhysicalConnectionRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	PhysicalConnectionId string `position:"Query" name:"PhysicalConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *TerminatePhysicalConnectionRequest) Invoke(client *sdk.Client) (resp *TerminatePhysicalConnectionResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "TerminatePhysicalConnection", "ecs", "")
	resp = &TerminatePhysicalConnectionResponse{}
	err = client.DoAction(req, resp)
	return
}

type TerminatePhysicalConnectionResponse struct {
	responses.BaseResponse
	RequestId string
}
