package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopCreatemarketingRequest struct {
	requests.RpcRequest
	Etime string `position:"Query" name:"Etime"`
	Name  string `position:"Query" name:"Name"`
	Stime string `position:"Query" name:"Stime"`
	Sid   int64  `position:"Query" name:"Sid"`
}

func (req *ShopCreatemarketingRequest) Invoke(client *sdk.Client) (resp *ShopCreatemarketingResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopCreatemarketing", "", "")
	resp = &ShopCreatemarketingResponse{}
	err = client.DoAction(req, resp)
	return
}

type ShopCreatemarketingResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
