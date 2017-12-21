package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GroupTrendRequest struct {
	requests.RpcRequest
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (req *GroupTrendRequest) Invoke(client *sdk.Client) (resp *GroupTrendResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GroupTrend", "", "")
	resp = &GroupTrendResponse{}
	err = client.DoAction(req, resp)
	return
}

type GroupTrendResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
