package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsWarnCreateRequest struct {
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

func (r OnsWarnCreateRequest) Invoke(client *sdk.Client) (response *OnsWarnCreateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsWarnCreateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsWarnCreate", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsWarnCreateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsWarnCreateResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsWarnCreateResponse struct {
	RequestId string
	HelpUrl   string
}
