package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportRealtimeRequest struct {
	Agsid int64 `position:"Query" name:"Agsid"`
}

func (r ReportRealtimeRequest) Invoke(client *sdk.Client) (response *ReportRealtimeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReportRealtimeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ReportRealtime", "", "")

	resp := struct {
		*responses.BaseResponse
		ReportRealtimeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReportRealtimeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReportRealtimeResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
