package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopActionCustomeRequest struct {
	requests.RpcRequest
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (req *ShopActionCustomeRequest) Invoke(client *sdk.Client) (resp *ShopActionCustomeResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopActionCustome", "", "")
	resp = &ShopActionCustomeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ShopActionCustomeResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
