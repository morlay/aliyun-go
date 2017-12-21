package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopGroupDeleteRequest struct {
	Gid int64 `position:"Query" name:"Gid"`
}

func (r ShopGroupDeleteRequest) Invoke(client *sdk.Client) (response *ShopGroupDeleteResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopGroupDeleteRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopGroupDelete", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopGroupDeleteResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopGroupDeleteResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopGroupDeleteResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
