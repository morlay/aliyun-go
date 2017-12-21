package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId      string
	VpnGatewayId   string
	VpcId          string
	VSwitchId      string
	InternetIp     string
	IntranetIp     string
	CreateTime     int64
	EndTime        int64
	Spec           string
	Name           string
	Description    string
	Status         string
	BusinessStatus string
}
