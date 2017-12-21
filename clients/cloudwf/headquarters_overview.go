package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type HeadquartersOverviewRequest struct {
	requests.RpcRequest
	Bid int64 `position:"Query" name:"Bid"`
}

func (req *HeadquartersOverviewRequest) Invoke(client *sdk.Client) (resp *HeadquartersOverviewResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "HeadquartersOverview", "", "")
	resp = &HeadquartersOverviewResponse{}
	err = client.DoAction(req, resp)
	return
}

type HeadquartersOverviewResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
