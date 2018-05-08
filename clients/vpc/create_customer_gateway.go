package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateCustomerGatewayRequest struct {
	requests.RpcRequest
	IpAddress            string `position:"Query" name:"IpAddress"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CreateCustomerGatewayRequest) Invoke(client *sdk.Client) (resp *CreateCustomerGatewayResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateCustomerGateway", "vpc", "")
	resp = &CreateCustomerGatewayResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateCustomerGatewayResponse struct {
	responses.BaseResponse
	RequestId         common.String
	CustomerGatewayId common.String
	IpAddress         common.String
	Name              common.String
	Description       common.String
	CreateTime        common.Long
}
