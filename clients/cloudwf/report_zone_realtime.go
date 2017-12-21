package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportZoneRealtimeRequest struct {
	requests.RpcRequest
	Agsid int64 `position:"Query" name:"Agsid"`
}

func (req *ReportZoneRealtimeRequest) Invoke(client *sdk.Client) (resp *ReportZoneRealtimeResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ReportZoneRealtime", "", "")
	resp = &ReportZoneRealtimeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReportZoneRealtimeResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
