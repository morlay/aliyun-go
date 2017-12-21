package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteNatGatewayRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	NatGatewayId         string `position:"Query" name:"NatGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteNatGatewayRequest) Invoke(client *sdk.Client) (resp *DeleteNatGatewayResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteNatGateway", "ecs", "")
	resp = &DeleteNatGatewayResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteNatGatewayResponse struct {
	responses.BaseResponse
	RequestId string
}
