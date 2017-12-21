package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsConsumerResetOffsetRequest struct {
	PreventCache   int64  `position:"Query" name:"PreventCache"`
	OnsRegionId    string `position:"Query" name:"OnsRegionId"`
	OnsPlatform    string `position:"Query" name:"OnsPlatform"`
	ConsumerId     string `position:"Query" name:"ConsumerId"`
	Topic          string `position:"Query" name:"Topic"`
	ResetTimestamp int64  `position:"Query" name:"ResetTimestamp"`
	Type           int    `position:"Query" name:"Type"`
}

func (r OnsConsumerResetOffsetRequest) Invoke(client *sdk.Client) (response *OnsConsumerResetOffsetResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsConsumerResetOffsetRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsConsumerResetOffset", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsConsumerResetOffsetResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsConsumerResetOffsetResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsConsumerResetOffsetResponse struct {
	RequestId string
	HelpUrl   string
}
