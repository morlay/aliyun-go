package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopGroupCreateRequest struct {
	ShopIds     string `position:"Query" name:"ShopIds"`
	Name        string `position:"Query" name:"Name"`
	Description string `position:"Query" name:"Description"`
	Bid         int64  `position:"Query" name:"Bid"`
}

func (r ShopGroupCreateRequest) Invoke(client *sdk.Client) (response *ShopGroupCreateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopGroupCreateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopGroupCreate", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopGroupCreateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopGroupCreateResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopGroupCreateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
