package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMessageSendRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	ProducerId   string `position:"Query" name:"ProducerId"`
	Tag          string `position:"Query" name:"Tag"`
	Message      string `position:"Query" name:"Message"`
	Key          string `position:"Query" name:"Key"`
}

func (r OnsMessageSendRequest) Invoke(client *sdk.Client) (response *OnsMessageSendResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsMessageSendRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMessageSend", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsMessageSendResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsMessageSendResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsMessageSendResponse struct {
	RequestId string
	HelpUrl   string
	Data      string
}
