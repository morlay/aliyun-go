package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateOrderRequest struct {
	requests.RpcRequest
	Data  string `position:"Body" name:"Data"`
	CsbId int64  `position:"Query" name:"CsbId"`
}

func (req *UpdateOrderRequest) Invoke(client *sdk.Client) (resp *UpdateOrderResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "UpdateOrder", "CSB", "")
	resp = &UpdateOrderResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateOrderResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
}
