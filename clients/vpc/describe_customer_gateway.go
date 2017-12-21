package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCustomerGatewayRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CustomerGatewayId    string `position:"Query" name:"CustomerGatewayId"`
}

func (req *DescribeCustomerGatewayRequest) Invoke(client *sdk.Client) (resp *DescribeCustomerGatewayResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeCustomerGateway", "vpc", "")
	resp = &DescribeCustomerGatewayResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCustomerGatewayResponse struct {
	responses.BaseResponse
	RequestId         string
	CustomerGatewayId string
	IpAddress         string
	Name              string
	Description       string
	CreateTime        int64
}
