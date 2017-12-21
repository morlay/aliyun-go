package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type HeadquartersTrendRequest struct {
	requests.RpcRequest
	Bid int64 `position:"Query" name:"Bid"`
}

func (req *HeadquartersTrendRequest) Invoke(client *sdk.Client) (resp *HeadquartersTrendResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "HeadquartersTrend", "", "")
	resp = &HeadquartersTrendResponse{}
	err = client.DoAction(req, resp)
	return
}

type HeadquartersTrendResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
