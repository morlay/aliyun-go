package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMessagePushRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	ClientId     string `position:"Query" name:"ClientId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	MsgId        string `position:"Query" name:"MsgId"`
	Topic        string `position:"Query" name:"Topic"`
}

func (req *OnsMessagePushRequest) Invoke(client *sdk.Client) (resp *OnsMessagePushResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMessagePush", "", "")
	resp = &OnsMessagePushResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsMessagePushResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
}
