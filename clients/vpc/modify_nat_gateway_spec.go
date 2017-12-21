package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyNatGatewaySpecRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	NatGatewayId         string `position:"Query" name:"NatGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Spec                 string `position:"Query" name:"Spec"`
}

func (req *ModifyNatGatewaySpecRequest) Invoke(client *sdk.Client) (resp *ModifyNatGatewaySpecResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyNatGatewaySpec", "vpc", "")
	resp = &ModifyNatGatewaySpecResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyNatGatewaySpecResponse struct {
	responses.BaseResponse
	RequestId string
}
