package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopActionCustomeRequest struct {
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (r ShopActionCustomeRequest) Invoke(client *sdk.Client) (response *ShopActionCustomeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopActionCustomeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopActionCustome", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopActionCustomeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopActionCustomeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopActionCustomeResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
