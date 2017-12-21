package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopSetfiltermacRequest struct {
	requests.RpcRequest
	Mac string `position:"Query" name:"Mac"`
	Sid int64  `position:"Query" name:"Sid"`
}

func (req *ShopSetfiltermacRequest) Invoke(client *sdk.Client) (resp *ShopSetfiltermacResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopSetfiltermac", "", "")
	resp = &ShopSetfiltermacResponse{}
	err = client.DoAction(req, resp)
	return
}

type ShopSetfiltermacResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
