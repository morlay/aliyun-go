package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateOrderListRequest struct {
	requests.RpcRequest
	Data string `position:"Body" name:"Data"`
}

func (req *UpdateOrderListRequest) Invoke(client *sdk.Client) (resp *UpdateOrderListResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "UpdateOrderList", "CSB", "")
	resp = &UpdateOrderListResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateOrderListResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      UpdateOrderListData
}

type UpdateOrderListData struct {
	UpdateCount common.Integer
}
