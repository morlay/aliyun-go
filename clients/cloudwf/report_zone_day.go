package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportZoneDayRequest struct {
	BeginDate string `position:"Query" name:"BeginDate"`
	EndDate   string `position:"Query" name:"EndDate"`
	Agsid     int64  `position:"Query" name:"Agsid"`
}

func (r ReportZoneDayRequest) Invoke(client *sdk.Client) (response *ReportZoneDayResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReportZoneDayRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ReportZoneDay", "", "")

	resp := struct {
		*responses.BaseResponse
		ReportZoneDayResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReportZoneDayResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReportZoneDayResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
