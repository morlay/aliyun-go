package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsMqttQueryTraceByTraceIdRequest struct {
	requests.RpcRequest
	TraceId      string `position:"Query" name:"TraceId"`
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
}

func (req *OnsMqttQueryTraceByTraceIdRequest) Invoke(client *sdk.Client) (resp *OnsMqttQueryTraceByTraceIdResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttQueryTraceByTraceId", "", "")
	resp = &OnsMqttQueryTraceByTraceIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsMqttQueryTraceByTraceIdResponse struct {
	responses.BaseResponse
	RequestId common.String
	HelpUrl   common.String
	PushCount common.Long
}
