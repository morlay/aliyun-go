package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyVpnGatewayAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyVpnGatewayAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyVpnGatewayAttributeResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVpnGatewayAttribute", "vpc", "")
	resp = &ModifyVpnGatewayAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyVpnGatewayAttributeResponse struct {
	responses.BaseResponse
	RequestId      common.String
	VpnGatewayId   common.String
	VpcId          common.String
	VSwitchId      common.String
	InternetIp     common.String
	IntranetIp     common.String
	CreateTime     common.Long
	EndTime        common.Long
	Spec           common.String
	Name           common.String
	Description    common.String
	Status         common.String
	BusinessStatus common.String
}
