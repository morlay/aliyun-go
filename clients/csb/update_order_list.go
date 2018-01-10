package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code      int
	Message   string
	RequestId string
	Data      UpdateOrderListData
}

type UpdateOrderListData struct {
	UpdateCount int
}
