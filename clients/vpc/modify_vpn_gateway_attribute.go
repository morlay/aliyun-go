package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyVpnGatewayAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyVpnGatewayAttributeRequest) Invoke(client *sdk.Client) (response *ModifyVpnGatewayAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyVpnGatewayAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVpnGatewayAttribute", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		ModifyVpnGatewayAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyVpnGatewayAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyVpnGatewayAttributeResponse struct {
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
