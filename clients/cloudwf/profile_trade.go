package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProfileTradeRequest struct {
	requests.RpcRequest
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (req *ProfileTradeRequest) Invoke(client *sdk.Client) (resp *ProfileTradeResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileTrade", "", "")
	resp = &ProfileTradeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ProfileTradeResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
