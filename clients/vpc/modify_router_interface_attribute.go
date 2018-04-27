package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyRouterInterfaceAttributeRequest struct {
	requests.RpcRequest
	OppositeRouterId         string `position:"Query" name:"OppositeRouterId"`
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	DeleteHealthCheckIp      string `position:"Query" name:"DeleteHealthCheckIp"`
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

func (req *ModifyRouterInterfaceAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyRouterInterfaceAttributeResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyRouterInterfaceAttribute", "vpc", "")
	resp = &ModifyRouterInterfaceAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyRouterInterfaceAttributeResponse struct {
	responses.BaseResponse
	RequestId string
}
