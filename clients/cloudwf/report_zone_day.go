package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportZoneDayRequest struct {
	requests.RpcRequest
	BeginDate string `position:"Query" name:"BeginDate"`
	EndDate   string `position:"Query" name:"EndDate"`
	Agsid     int64  `position:"Query" name:"Agsid"`
}

func (req *ReportZoneDayRequest) Invoke(client *sdk.Client) (resp *ReportZoneDayResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ReportZoneDay", "", "")
	resp = &ReportZoneDayResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReportZoneDayResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
