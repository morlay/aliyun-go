package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsSubscriptionUpdateRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	ReadEnable   string `position:"Query" name:"ReadEnable"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
}

func (r OnsSubscriptionUpdateRequest) Invoke(client *sdk.Client) (response *OnsSubscriptionUpdateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsSubscriptionUpdateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsSubscriptionUpdate", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsSubscriptionUpdateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsSubscriptionUpdateResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsSubscriptionUpdateResponse struct {
	RequestId string
	HelpUrl   string
}
