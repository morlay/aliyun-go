package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopActionReturningRequest struct {
	requests.RpcRequest
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (req *ShopActionReturningRequest) Invoke(client *sdk.Client) (resp *ShopActionReturningResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopActionReturning", "", "")
	resp = &ShopActionReturningResponse{}
	err = client.DoAction(req, resp)
	return
}

type ShopActionReturningResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
