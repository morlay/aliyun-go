package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsMqttBuyCheckRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Data         string `position:"Query" name:"Data"`
}

func (req *OnsMqttBuyCheckRequest) Invoke(client *sdk.Client) (resp *OnsMqttBuyCheckResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttBuyCheck", "", "")
	resp = &OnsMqttBuyCheckResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsMqttBuyCheckResponse struct {
	responses.BaseResponse
	Success   bool
	RequestId common.String
	Code      common.String
	Message   common.String
	Data      common.String
}
