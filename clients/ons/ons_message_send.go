package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMessageSendRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	ProducerId   string `position:"Query" name:"ProducerId"`
	Tag          string `position:"Query" name:"Tag"`
	Message      string `position:"Query" name:"Message"`
	Key          string `position:"Query" name:"Key"`
}

func (req *OnsMessageSendRequest) Invoke(client *sdk.Client) (resp *OnsMessageSendResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMessageSend", "", "")
	resp = &OnsMessageSendResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsMessageSendResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
	Data      string
}
