package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsSubscriptionCreateRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Topic        string `position:"Query" name:"Topic"`
}

func (req *OnsSubscriptionCreateRequest) Invoke(client *sdk.Client) (resp *OnsSubscriptionCreateResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsSubscriptionCreate", "", "")
	resp = &OnsSubscriptionCreateResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsSubscriptionCreateResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
}
