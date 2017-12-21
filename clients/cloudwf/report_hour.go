package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportHourRequest struct {
	BeginDate string `position:"Query" name:"BeginDate"`
	EndDate   string `position:"Query" name:"EndDate"`
	Agsid     int64  `position:"Query" name:"Agsid"`
}

func (r ReportHourRequest) Invoke(client *sdk.Client) (response *ReportHourResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReportHourRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ReportHour", "", "")

	resp := struct {
		*responses.BaseResponse
		ReportHourResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReportHourResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReportHourResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
