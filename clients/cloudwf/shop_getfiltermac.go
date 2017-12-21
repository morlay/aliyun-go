package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopGetfiltermacRequest struct {
	Sid int64 `position:"Query" name:"Sid"`
}

func (r ShopGetfiltermacRequest) Invoke(client *sdk.Client) (response *ShopGetfiltermacResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopGetfiltermacRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopGetfiltermac", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopGetfiltermacResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopGetfiltermacResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopGetfiltermacResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
