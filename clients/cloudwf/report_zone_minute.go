package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportZoneMinuteRequest struct {
	requests.RpcRequest
	BeginDate string `position:"Query" name:"BeginDate"`
	EndDate   string `position:"Query" name:"EndDate"`
	Agsid     int64  `position:"Query" name:"Agsid"`
}

func (req *ReportZoneMinuteRequest) Invoke(client *sdk.Client) (resp *ReportZoneMinuteResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ReportZoneMinute", "", "")
	resp = &ReportZoneMinuteResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReportZoneMinuteResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
