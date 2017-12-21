package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GroupOverviewRequest struct {
	requests.RpcRequest
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (req *GroupOverviewRequest) Invoke(client *sdk.Client) (resp *GroupOverviewResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GroupOverview", "", "")
	resp = &GroupOverviewResponse{}
	err = client.DoAction(req, resp)
	return
}

type GroupOverviewResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
