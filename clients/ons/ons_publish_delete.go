package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsPublishDeleteRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	ProducerId   string `position:"Query" name:"ProducerId"`
}

func (r OnsPublishDeleteRequest) Invoke(client *sdk.Client) (response *OnsPublishDeleteResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsPublishDeleteRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsPublishDelete", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsPublishDeleteResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsPublishDeleteResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsPublishDeleteResponse struct {
	RequestId string
	HelpUrl   string
}
