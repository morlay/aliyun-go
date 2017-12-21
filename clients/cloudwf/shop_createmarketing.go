package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopCreatemarketingRequest struct {
	Etime string `position:"Query" name:"Etime"`
	Name  string `position:"Query" name:"Name"`
	Stime string `position:"Query" name:"Stime"`
	Sid   int64  `position:"Query" name:"Sid"`
}

func (r ShopCreatemarketingRequest) Invoke(client *sdk.Client) (response *ShopCreatemarketingResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopCreatemarketingRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopCreatemarketing", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopCreatemarketingResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopCreatemarketingResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopCreatemarketingResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
