package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportRealtimeRequest struct {
	requests.RpcRequest
	Agsid int64 `position:"Query" name:"Agsid"`
}

func (req *ReportRealtimeRequest) Invoke(client *sdk.Client) (resp *ReportRealtimeResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ReportRealtime", "", "")
	resp = &ReportRealtimeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReportRealtimeResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
