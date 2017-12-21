package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMqttManualUpdateRuleRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	OwnerId      string `position:"Query" name:"OwnerId"`
	AdminKey     string `position:"Query" name:"AdminKey"`
}

func (req *OnsMqttManualUpdateRuleRequest) Invoke(client *sdk.Client) (resp *OnsMqttManualUpdateRuleResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttManualUpdateRule", "", "")
	resp = &OnsMqttManualUpdateRuleResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsMqttManualUpdateRuleResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
}
