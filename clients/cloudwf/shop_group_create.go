package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopGroupCreateRequest struct {
	requests.RpcRequest
	ShopIds     string `position:"Query" name:"ShopIds"`
	Name        string `position:"Query" name:"Name"`
	Description string `position:"Query" name:"Description"`
	Bid         int64  `position:"Query" name:"Bid"`
}

func (req *ShopGroupCreateRequest) Invoke(client *sdk.Client) (resp *ShopGroupCreateResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopGroupCreate", "", "")
	resp = &ShopGroupCreateResponse{}
	err = client.DoAction(req, resp)
	return
}

type ShopGroupCreateResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
