package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsTopicDeleteRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	Cluster      string `position:"Query" name:"Cluster"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
}

func (r OnsTopicDeleteRequest) Invoke(client *sdk.Client) (response *OnsTopicDeleteResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsTopicDeleteRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsTopicDelete", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsTopicDeleteResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsTopicDeleteResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsTopicDeleteResponse struct {
	RequestId string
	HelpUrl   string
}
