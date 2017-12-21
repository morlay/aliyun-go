package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsConsumerTimeSpanRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Topic        string `position:"Query" name:"Topic"`
}

func (req *OnsConsumerTimeSpanRequest) Invoke(client *sdk.Client) (resp *OnsConsumerTimeSpanResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsConsumerTimeSpan", "", "")
	resp = &OnsConsumerTimeSpanResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsConsumerTimeSpanResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
	Data      OnsConsumerTimeSpanData
}

type OnsConsumerTimeSpanData struct {
	Topic            string
	MinTimeStamp     int64
	MaxTimeStamp     int64
	ConsumeTimeStamp int64
}
