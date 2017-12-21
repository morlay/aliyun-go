package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopInfoRequest struct {
	Sid int64 `position:"Query" name:"Sid"`
}

func (r ShopInfoRequest) Invoke(client *sdk.Client) (response *ShopInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopInfoResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
