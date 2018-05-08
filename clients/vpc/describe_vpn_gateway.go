package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeVpnGatewayRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeVpnGatewayRequest) Invoke(client *sdk.Client) (resp *DescribeVpnGatewayResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVpnGateway", "vpc", "")
	resp = &DescribeVpnGatewayResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeVpnGatewayResponse struct {
	responses.BaseResponse
	RequestId         common.String
	VpnGatewayId      common.String
	VpcId             common.String
	VSwitchId         common.String
	InternetIp        common.String
	CreateTime        common.Long
	EndTime           common.Long
	Spec              common.String
	Name              common.String
	Description       common.String
	Status            common.String
	BusinessStatus    common.String
	ChargeType        common.String
	IpsecVpn          common.String
	SslVpn            common.String
	SslMaxConnections common.Long
}
