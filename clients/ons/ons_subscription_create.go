package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsSubscriptionCreateRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Topic        string `position:"Query" name:"Topic"`
}

func (r OnsSubscriptionCreateRequest) Invoke(client *sdk.Client) (response *OnsSubscriptionCreateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsSubscriptionCreateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsSubscriptionCreate", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsSubscriptionCreateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsSubscriptionCreateResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsSubscriptionCreateResponse struct {
	RequestId string
	HelpUrl   string
}
