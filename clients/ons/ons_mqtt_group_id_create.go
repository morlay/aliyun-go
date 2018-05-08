package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsMqttGroupIdCreateRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	GroupId      string `position:"Query" name:"GroupId"`
	Topic        string `position:"Query" name:"Topic"`
}

func (req *OnsMqttGroupIdCreateRequest) Invoke(client *sdk.Client) (resp *OnsMqttGroupIdCreateResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsMqttGroupIdCreate", "", "")
	resp = &OnsMqttGroupIdCreateResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsMqttGroupIdCreateResponse struct {
	responses.BaseResponse
	RequestId common.String
	HelpUrl   common.String
}
