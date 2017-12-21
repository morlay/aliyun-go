package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopActionReturningRequest struct {
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (r ShopActionReturningRequest) Invoke(client *sdk.Client) (response *ShopActionReturningResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopActionReturningRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopActionReturning", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopActionReturningResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopActionReturningResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopActionReturningResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
