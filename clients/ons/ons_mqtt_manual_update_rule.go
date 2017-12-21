package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsMqttManualUpdateRuleRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	OwnerId      string `position:"Query" name:"OwnerId"`
	AdminKey     string `position:"Query" name:"AdminKey"`
}

func (r OnsMqttManualUpdateRuleRequest) Invoke(client *sdk.Client) (response *OnsMqttManualUpdateRuleResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsMqttManualUpdateRuleRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttManualUpdateRule", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsMqttManualUpdateRuleResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsMqttManualUpdateRuleResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsMqttManualUpdateRuleResponse struct {
	RequestId string
	HelpUrl   string
}
