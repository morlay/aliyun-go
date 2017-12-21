package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsTraceQueryByMsgIdRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	MsgId        string `position:"Query" name:"MsgId"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
}

func (r OnsTraceQueryByMsgIdRequest) Invoke(client *sdk.Client) (response *OnsTraceQueryByMsgIdResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsTraceQueryByMsgIdRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsTraceQueryByMsgId", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsTraceQueryByMsgIdResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsTraceQueryByMsgIdResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsTraceQueryByMsgIdResponse struct {
	RequestId string
	HelpUrl   string
	QueryId   string
}
