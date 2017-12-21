package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCustomerGatewayAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CustomerGatewayId    string `position:"Query" name:"CustomerGatewayId"`
}

func (r ModifyCustomerGatewayAttributeRequest) Invoke(client *sdk.Client) (response *ModifyCustomerGatewayAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyCustomerGatewayAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyCustomerGatewayAttribute", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		ModifyCustomerGatewayAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyCustomerGatewayAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyCustomerGatewayAttributeResponse struct {
	RequestId         string
	CustomerGatewayId string
	IpAddress         string
	Name              string
	Description       string
	CreateTime        int64
}
