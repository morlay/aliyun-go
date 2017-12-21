package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsPublishCreateRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	AppName      string `position:"Query" name:"AppName"`
	Topic        string `position:"Query" name:"Topic"`
	ProducerId   string `position:"Query" name:"ProducerId"`
}

func (r OnsPublishCreateRequest) Invoke(client *sdk.Client) (response *OnsPublishCreateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsPublishCreateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsPublishCreate", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsPublishCreateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsPublishCreateResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsPublishCreateResponse struct {
	RequestId string
	HelpUrl   string
}
