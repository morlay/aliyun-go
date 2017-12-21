package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMqttBuyProduceRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Data         string `position:"Query" name:"Data"`
}

func (req *OnsMqttBuyProduceRequest) Invoke(client *sdk.Client) (resp *OnsMqttBuyProduceResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttBuyProduce", "", "")
	resp = &OnsMqttBuyProduceResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsMqttBuyProduceResponse struct {
	responses.BaseResponse
	Success   bool
	RequestId string
	Code      string
	Message   string
	Data      string
}
