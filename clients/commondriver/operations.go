package commondriver

import (
	"github.com/morlay/aliyun-go/core"
)

func (c *CommondriverClient) GetOrderIdByCheckBeforePay(req *GetOrderIdByCheckBeforePayArgs) (resp *GetOrderIdByCheckBeforePayResponse, err error) {
	resp = &GetOrderIdByCheckBeforePayResponse{}
	err = c.Request("GetOrderIdByCheckBeforePay", req, resp)
	return
}

type GetOrderIdByCheckBeforePayArgs struct {
	OrderId string
}
type GetOrderIdByCheckBeforePayResponse struct {
	RequestId string
	Success   core.Bool
	Data      core.Bool
}

func (c *CommondriverClient) GetOrderIdByQueryPurchase(req *GetOrderIdByQueryPurchaseArgs) (resp *GetOrderIdByQueryPurchaseResponse, err error) {
	resp = &GetOrderIdByQueryPurchaseResponse{}
	err = c.Request("GetOrderIdByQueryPurchase", req, resp)
	return
}

type GetOrderIdByQueryPurchaseArgs struct {
	OrderId string
}
type GetOrderIdByQueryPurchaseResponse struct {
	RequestId string
	Success   core.Bool
	Data      core.Bool
}
