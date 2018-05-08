package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteCustomerGatewayRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CustomerGatewayId    string `position:"Query" name:"CustomerGatewayId"`
}

func (req *DeleteCustomerGatewayRequest) Invoke(client *sdk.Client) (resp *DeleteCustomerGatewayResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteCustomerGateway", "vpc", "")
	resp = &DeleteCustomerGatewayResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCustomerGatewayResponse struct {
	responses.BaseResponse
	RequestId common.String
}
