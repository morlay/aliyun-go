package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	HelpUrl   common.String
	Data      OnsConsumerTimeSpanData
}

type OnsConsumerTimeSpanData struct {
	Topic            common.String
	MinTimeStamp     common.Long
	MaxTimeStamp     common.Long
	ConsumeTimeStamp common.Long
}
