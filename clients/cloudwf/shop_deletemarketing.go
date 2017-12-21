package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopDeletemarketingRequest struct {
	Id  int64 `position:"Query" name:"Id"`
	Sid int64 `position:"Query" name:"Sid"`
}

func (r ShopDeletemarketingRequest) Invoke(client *sdk.Client) (response *ShopDeletemarketingResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopDeletemarketingRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopDeletemarketing", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopDeletemarketingResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopDeletemarketingResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopDeletemarketingResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
