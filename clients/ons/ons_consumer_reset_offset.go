package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsConsumerResetOffsetRequest struct {
	requests.RpcRequest
	PreventCache   int64  `position:"Query" name:"PreventCache"`
	OnsRegionId    string `position:"Query" name:"OnsRegionId"`
	OnsPlatform    string `position:"Query" name:"OnsPlatform"`
	ConsumerId     string `position:"Query" name:"ConsumerId"`
	Topic          string `position:"Query" name:"Topic"`
	ResetTimestamp int64  `position:"Query" name:"ResetTimestamp"`
	Type           int    `position:"Query" name:"Type"`
}

func (req *OnsConsumerResetOffsetRequest) Invoke(client *sdk.Client) (resp *OnsConsumerResetOffsetResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsConsumerResetOffset", "", "")
	resp = &OnsConsumerResetOffsetResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsConsumerResetOffsetResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
}
