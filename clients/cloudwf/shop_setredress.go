package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopSetredressRequest struct {
	requests.RpcRequest
	Workday     string `position:"Query" name:"Workday"`
	Filterclose int    `position:"Query" name:"Filterclose"`
	Minstoptime int    `position:"Query" name:"Minstoptime"`
	Holiday     string `position:"Query" name:"Holiday"`
	Hnum        string `position:"Query" name:"Hnum"`
	Sid         int64  `position:"Query" name:"Sid"`
	Clerk       int    `position:"Query" name:"Clerk"`
	Filterstate int    `position:"Query" name:"Filterstate"`
	Wnum        string `position:"Query" name:"Wnum"`
	State       int    `position:"Query" name:"State"`
	Crowdfixed  int    `position:"Query" name:"Crowdfixed"`
	Maxstoptime int    `position:"Query" name:"Maxstoptime"`
}

func (req *ShopSetredressRequest) Invoke(client *sdk.Client) (resp *ShopSetredressResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopSetredress", "", "")
	resp = &ShopSetredressResponse{}
	err = client.DoAction(req, resp)
	return
}

type ShopSetredressResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
