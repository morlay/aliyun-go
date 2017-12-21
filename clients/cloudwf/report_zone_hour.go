package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportZoneHourRequest struct {
	BeginDate string `position:"Query" name:"BeginDate"`
	EndDate   string `position:"Query" name:"EndDate"`
	Agsid     int64  `position:"Query" name:"Agsid"`
}

func (r ReportZoneHourRequest) Invoke(client *sdk.Client) (response *ReportZoneHourResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReportZoneHourRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ReportZoneHour", "", "")

	resp := struct {
		*responses.BaseResponse
		ReportZoneHourResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReportZoneHourResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReportZoneHourResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
