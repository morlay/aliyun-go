package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsConsumerTimeSpanRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Topic        string `position:"Query" name:"Topic"`
}

func (r OnsConsumerTimeSpanRequest) Invoke(client *sdk.Client) (response *OnsConsumerTimeSpanResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsConsumerTimeSpanRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsConsumerTimeSpan", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsConsumerTimeSpanResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsConsumerTimeSpanResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsConsumerTimeSpanResponse struct {
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
