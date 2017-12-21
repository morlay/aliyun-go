package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyRouterInterfaceAttributeRequest struct {
	OppositeRouterId         string `position:"Query" name:"OppositeRouterId"`
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	Description              string `position:"Query" name:"Description"`
	HealthCheckTargetIp      string `position:"Query" name:"HealthCheckTargetIp"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	RouterInterfaceId        string `position:"Query" name:"RouterInterfaceId"`
	OppositeInterfaceOwnerId int64  `position:"Query" name:"OppositeInterfaceOwnerId"`
	HealthCheckSourceIp      string `position:"Query" name:"HealthCheckSourceIp"`
	Name                     string `position:"Query" name:"Name"`
	OppositeRouterType       string `position:"Query" name:"OppositeRouterType"`
	OppositeInterfaceId      string `position:"Query" name:"OppositeInterfaceId"`
}

func (r ModifyRouterInterfaceAttributeRequest) Invoke(client *sdk.Client) (response *ModifyRouterInterfaceAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyRouterInterfaceAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyRouterInterfaceAttribute", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifyRouterInterfaceAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyRouterInterfaceAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyRouterInterfaceAttributeResponse struct {
	RequestId string
}
