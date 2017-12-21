package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopInfoRequest struct {
	requests.RpcRequest
	Sid int64 `position:"Query" name:"Sid"`
}

func (req *ShopInfoRequest) Invoke(client *sdk.Client) (resp *ShopInfoResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopInfo", "", "")
	resp = &ShopInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type ShopInfoResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
