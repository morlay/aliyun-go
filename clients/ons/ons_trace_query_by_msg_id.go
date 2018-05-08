package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsTraceQueryByMsgIdRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	MsgId        string `position:"Query" name:"MsgId"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
}

func (req *OnsTraceQueryByMsgIdRequest) Invoke(client *sdk.Client) (resp *OnsTraceQueryByMsgIdResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsTraceQueryByMsgId", "", "")
	resp = &OnsTraceQueryByMsgIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsTraceQueryByMsgIdResponse struct {
	responses.BaseResponse
	RequestId common.String
	HelpUrl   common.String
	QueryId   common.String
}
