package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeVpnGatewayRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeVpnGatewayRequest) Invoke(client *sdk.Client) (response *DescribeVpnGatewayResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeVpnGatewayRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVpnGateway", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DescribeVpnGatewayResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeVpnGatewayResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeVpnGatewayResponse struct {
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
