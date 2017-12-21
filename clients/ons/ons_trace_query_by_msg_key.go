package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsTraceQueryByMsgKeyRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
	MsgKey       string `position:"Query" name:"MsgKey"`
}

func (req *OnsTraceQueryByMsgKeyRequest) Invoke(client *sdk.Client) (resp *OnsTraceQueryByMsgKeyResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsTraceQueryByMsgKey", "", "")
	resp = &OnsTraceQueryByMsgKeyResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsTraceQueryByMsgKeyResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
	QueryId   string
}
