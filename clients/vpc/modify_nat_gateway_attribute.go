package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyNatGatewayAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	NatGatewayId         string `position:"Query" name:"NatGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyNatGatewayAttributeRequest) Invoke(client *sdk.Client) (response *ModifyNatGatewayAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyNatGatewayAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyNatGatewayAttribute", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		ModifyNatGatewayAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyNatGatewayAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyNatGatewayAttributeResponse struct {
	RequestId string
}
