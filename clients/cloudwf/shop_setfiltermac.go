package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopSetfiltermacRequest struct {
	Mac string `position:"Query" name:"Mac"`
	Sid int64  `position:"Query" name:"Sid"`
}

func (r ShopSetfiltermacRequest) Invoke(client *sdk.Client) (response *ShopSetfiltermacResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopSetfiltermacRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopSetfiltermac", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopSetfiltermacResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopSetfiltermacResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopSetfiltermacResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
