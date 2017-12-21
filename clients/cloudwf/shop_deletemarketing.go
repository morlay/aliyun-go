package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopDeletemarketingRequest struct {
	requests.RpcRequest
	Id  int64 `position:"Query" name:"Id"`
	Sid int64 `position:"Query" name:"Sid"`
}

func (req *ShopDeletemarketingRequest) Invoke(client *sdk.Client) (resp *ShopDeletemarketingResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopDeletemarketing", "", "")
	resp = &ShopDeletemarketingResponse{}
	err = client.DoAction(req, resp)
	return
}

type ShopDeletemarketingResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
