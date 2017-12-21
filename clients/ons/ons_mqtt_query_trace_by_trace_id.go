package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMqttQueryTraceByTraceIdRequest struct {
	TraceId      string `position:"Query" name:"TraceId"`
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
}

func (r OnsMqttQueryTraceByTraceIdRequest) Invoke(client *sdk.Client) (response *OnsMqttQueryTraceByTraceIdResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsMqttQueryTraceByTraceIdRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttQueryTraceByTraceId", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsMqttQueryTraceByTraceIdResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsMqttQueryTraceByTraceIdResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsMqttQueryTraceByTraceIdResponse struct {
	RequestId string
	HelpUrl   string
	PushCount int64
}
