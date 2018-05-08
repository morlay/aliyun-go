package commondriver

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetOrderIdByQueryPurchaseRequest struct {
	requests.RpcRequest
	OrderId string `position:"Query" name:"OrderId"`
}

func (req *GetOrderIdByQueryPurchaseRequest) Invoke(client *sdk.Client) (resp *GetOrderIdByQueryPurchaseResponse, err error) {
	req.InitWithApiInfo("Commondriver", "2015-12-29", "GetOrderIdByQueryPurchase", "", "")
	resp = &GetOrderIdByQueryPurchaseResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetOrderIdByQueryPurchaseResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Data      bool
}
