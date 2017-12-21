package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportZoneHourRequest struct {
	requests.RpcRequest
	BeginDate string `position:"Query" name:"BeginDate"`
	EndDate   string `position:"Query" name:"EndDate"`
	Agsid     int64  `position:"Query" name:"Agsid"`
}

func (req *ReportZoneHourRequest) Invoke(client *sdk.Client) (resp *ReportZoneHourResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ReportZoneHour", "", "")
	resp = &ReportZoneHourResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReportZoneHourResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
