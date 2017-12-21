package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsTraceQueryByMsgKeyRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
	MsgKey       string `position:"Query" name:"MsgKey"`
}

func (r OnsTraceQueryByMsgKeyRequest) Invoke(client *sdk.Client) (response *OnsTraceQueryByMsgKeyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsTraceQueryByMsgKeyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsTraceQueryByMsgKey", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsTraceQueryByMsgKeyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsTraceQueryByMsgKeyResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsTraceQueryByMsgKeyResponse struct {
	RequestId string
	HelpUrl   string
	QueryId   string
}
