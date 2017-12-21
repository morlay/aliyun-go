package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId      string
	VpnGatewayId   string
	VpcId          string
	VSwitchId      string
	InternetIp     string
	CreateTime     int64
	EndTime        int64
	Spec           string
	Name           string
	Description    string
	Status         string
	BusinessStatus string
	ChargeType     string
}
