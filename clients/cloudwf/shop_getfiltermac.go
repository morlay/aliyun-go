package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopGetfiltermacRequest struct {
	requests.RpcRequest
	Sid int64 `position:"Query" name:"Sid"`
}

func (req *ShopGetfiltermacRequest) Invoke(client *sdk.Client) (resp *ShopGetfiltermacResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopGetfiltermac", "", "")
	resp = &ShopGetfiltermacResponse{}
	err = client.DoAction(req, resp)
	return
}

type ShopGetfiltermacResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
