package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyNatGatewayAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	NatGatewayId         string `position:"Query" name:"NatGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyNatGatewayAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyNatGatewayAttributeResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyNatGatewayAttribute", "vpc", "")
	resp = &ModifyNatGatewayAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyNatGatewayAttributeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
