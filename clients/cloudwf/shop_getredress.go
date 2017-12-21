package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopGetredressRequest struct {
	Sid int64 `position:"Query" name:"Sid"`
}

func (r ShopGetredressRequest) Invoke(client *sdk.Client) (response *ShopGetredressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopGetredressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopGetredress", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopGetredressResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopGetredressResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopGetredressResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
