package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateRouterInterfaceRequest struct {
	requests.RpcRequest
	AccessPointId            string `position:"Query" name:"AccessPointId"`
	OppositeRouterId         string `position:"Query" name:"OppositeRouterId"`
	OppositeAccessPointId    string `position:"Query" name:"OppositeAccessPointId"`
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	Role                     string `position:"Query" name:"Role"`
	ClientToken              string `position:"Query" name:"ClientToken"`
	HealthCheckTargetIp      string `position:"Query" name:"HealthCheckTargetIp"`
	Description              string `position:"Query" name:"Description"`
	Spec                     string `position:"Query" name:"Spec"`
	UserCidr                 string `position:"Query" name:"UserCidr"`
	OppositeInterfaceId      string `position:"Query" name:"OppositeInterfaceId"`
	InstanceChargeType       string `position:"Query" name:"InstanceChargeType"`
	Period                   int    `position:"Query" name:"Period"`
	AutoPay                  string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	OppositeRegionId         string `position:"Query" name:"OppositeRegionId"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	OppositeInterfaceOwnerId string `position:"Query" name:"OppositeInterfaceOwnerId"`
	RouterType               string `position:"Query" name:"RouterType"`
	HealthCheckSourceIp      string `position:"Query" name:"HealthCheckSourceIp"`
	RouterId                 string `position:"Query" name:"RouterId"`
	OppositeRouterType       string `position:"Query" name:"OppositeRouterType"`
	Name                     string `position:"Query" name:"Name"`
	PricingCycle             string `position:"Query" name:"PricingCycle"`
}

func (req *CreateRouterInterfaceRequest) Invoke(client *sdk.Client) (resp *CreateRouterInterfaceResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateRouterInterface", "ecs", "")
	resp = &CreateRouterInterfaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateRouterInterfaceResponse struct {
	responses.BaseResponse
	RequestId         string
	RouterInterfaceId string
	OrderId           int64
}
