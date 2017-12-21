package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportZoneMinuteRequest struct {
	BeginDate string `position:"Query" name:"BeginDate"`
	EndDate   string `position:"Query" name:"EndDate"`
	Agsid     int64  `position:"Query" name:"Agsid"`
}

func (r ReportZoneMinuteRequest) Invoke(client *sdk.Client) (response *ReportZoneMinuteResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReportZoneMinuteRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ReportZoneMinute", "", "")

	resp := struct {
		*responses.BaseResponse
		ReportZoneMinuteResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReportZoneMinuteResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReportZoneMinuteResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
