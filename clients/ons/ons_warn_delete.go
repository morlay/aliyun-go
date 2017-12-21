package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsWarnDeleteRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Topic        string `position:"Query" name:"Topic"`
}

func (r OnsWarnDeleteRequest) Invoke(client *sdk.Client) (response *OnsWarnDeleteResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsWarnDeleteRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsWarnDelete", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsWarnDeleteResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsWarnDeleteResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsWarnDeleteResponse struct {
	RequestId string
	HelpUrl   string
}
