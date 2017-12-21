package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCustomerGatewayAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CustomerGatewayId    string `position:"Query" name:"CustomerGatewayId"`
}

func (req *ModifyCustomerGatewayAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyCustomerGatewayAttributeResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyCustomerGatewayAttribute", "vpc", "")
	resp = &ModifyCustomerGatewayAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyCustomerGatewayAttributeResponse struct {
	responses.BaseResponse
	RequestId         string
	CustomerGatewayId string
	IpAddress         string
	Name              string
	Description       string
	CreateTime        int64
}
