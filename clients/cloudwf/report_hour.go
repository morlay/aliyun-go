package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportHourRequest struct {
	requests.RpcRequest
	BeginDate string `position:"Query" name:"BeginDate"`
	EndDate   string `position:"Query" name:"EndDate"`
	Agsid     int64  `position:"Query" name:"Agsid"`
}

func (req *ReportHourRequest) Invoke(client *sdk.Client) (resp *ReportHourResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ReportHour", "", "")
	resp = &ReportHourResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReportHourResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
