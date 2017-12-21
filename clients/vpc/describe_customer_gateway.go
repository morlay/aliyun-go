package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCustomerGatewayRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CustomerGatewayId    string `position:"Query" name:"CustomerGatewayId"`
}

func (r DescribeCustomerGatewayRequest) Invoke(client *sdk.Client) (response *DescribeCustomerGatewayResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeCustomerGatewayRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeCustomerGateway", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DescribeCustomerGatewayResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeCustomerGatewayResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeCustomerGatewayResponse struct {
	RequestId         string
	CustomerGatewayId string
	IpAddress         string
	Name              string
	Description       string
	CreateTime        int64
}
