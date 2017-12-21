package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportDayRequest struct {
	BeginDate string `position:"Query" name:"BeginDate"`
	EndDate   string `position:"Query" name:"EndDate"`
	Agsid     int64  `position:"Query" name:"Agsid"`
}

func (r ReportDayRequest) Invoke(client *sdk.Client) (response *ReportDayResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReportDayRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ReportDay", "", "")

	resp := struct {
		*responses.BaseResponse
		ReportDayResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReportDayResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReportDayResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
