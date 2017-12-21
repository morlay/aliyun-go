package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCustomerGatewayRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CustomerGatewayId    string `position:"Query" name:"CustomerGatewayId"`
}

func (r DeleteCustomerGatewayRequest) Invoke(client *sdk.Client) (response *DeleteCustomerGatewayResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteCustomerGatewayRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteCustomerGateway", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DeleteCustomerGatewayResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteCustomerGatewayResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteCustomerGatewayResponse struct {
	RequestId string
}
