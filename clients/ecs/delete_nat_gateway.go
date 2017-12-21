package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteNatGatewayRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	NatGatewayId         string `position:"Query" name:"NatGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteNatGatewayRequest) Invoke(client *sdk.Client) (response *DeleteNatGatewayResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteNatGatewayRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteNatGateway", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DeleteNatGatewayResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteNatGatewayResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteNatGatewayResponse struct {
	RequestId string
}
