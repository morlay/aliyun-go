package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsWarnCreateRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	BlockTime    string `position:"Query" name:"BlockTime"`
	Level        string `position:"Query" name:"Level"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	DelayTime    string `position:"Query" name:"DelayTime"`
	Topic        string `position:"Query" name:"Topic"`
	Threshold    string `position:"Query" name:"Threshold"`
	AlertTime    string `position:"Query" name:"AlertTime"`
	Contacts     string `position:"Query" name:"Contacts"`
}

func (req *OnsWarnCreateRequest) Invoke(client *sdk.Client) (resp *OnsWarnCreateResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsWarnCreate", "", "")
	resp = &OnsWarnCreateResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsWarnCreateResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
}
