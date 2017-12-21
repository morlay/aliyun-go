package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsSubscriptionDeleteRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Topic        string `position:"Query" name:"Topic"`
}

func (r OnsSubscriptionDeleteRequest) Invoke(client *sdk.Client) (response *OnsSubscriptionDeleteResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsSubscriptionDeleteRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsSubscriptionDelete", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsSubscriptionDeleteResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsSubscriptionDeleteResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsSubscriptionDeleteResponse struct {
	RequestId string
	HelpUrl   string
}
