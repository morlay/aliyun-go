package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EnablePhysicalConnectionRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	PhysicalConnectionId string `position:"Query" name:"PhysicalConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *EnablePhysicalConnectionRequest) Invoke(client *sdk.Client) (resp *EnablePhysicalConnectionResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "EnablePhysicalConnection", "ecs", "")
	resp = &EnablePhysicalConnectionResponse{}
	err = client.DoAction(req, resp)
	return
}

type EnablePhysicalConnectionResponse struct {
	responses.BaseResponse
	RequestId string
}
