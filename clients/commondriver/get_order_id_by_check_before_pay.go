package commondriver

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetOrderIdByCheckBeforePayRequest struct {
	requests.RpcRequest
	OrderId string `position:"Query" name:"OrderId"`
}

func (req *GetOrderIdByCheckBeforePayRequest) Invoke(client *sdk.Client) (resp *GetOrderIdByCheckBeforePayResponse, err error) {
	req.InitWithApiInfo("Commondriver", "2015-12-29", "GetOrderIdByCheckBeforePay", "", "")
	resp = &GetOrderIdByCheckBeforePayResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetOrderIdByCheckBeforePayResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Data      bool
}
