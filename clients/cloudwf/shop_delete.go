package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopDeleteRequest struct {
	Sid int64 `position:"Query" name:"Sid"`
}

func (r ShopDeleteRequest) Invoke(client *sdk.Client) (response *ShopDeleteResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopDeleteRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopDelete", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopDeleteResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopDeleteResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopDeleteResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
