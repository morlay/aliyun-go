package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopGetredressRequest struct {
	requests.RpcRequest
	Sid int64 `position:"Query" name:"Sid"`
}

func (req *ShopGetredressRequest) Invoke(client *sdk.Client) (resp *ShopGetredressResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopGetredress", "", "")
	resp = &ShopGetredressResponse{}
	err = client.DoAction(req, resp)
	return
}

type ShopGetredressResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
