package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportZoneRealtimeRequest struct {
	Agsid int64 `position:"Query" name:"Agsid"`
}

func (r ReportZoneRealtimeRequest) Invoke(client *sdk.Client) (response *ReportZoneRealtimeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReportZoneRealtimeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ReportZoneRealtime", "", "")

	resp := struct {
		*responses.BaseResponse
		ReportZoneRealtimeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReportZoneRealtimeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReportZoneRealtimeResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
