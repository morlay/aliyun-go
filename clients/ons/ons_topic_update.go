package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsTopicUpdateRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Perm         int    `position:"Query" name:"Perm"`
	Topic        string `position:"Query" name:"Topic"`
}

func (r OnsTopicUpdateRequest) Invoke(client *sdk.Client) (response *OnsTopicUpdateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsTopicUpdateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsTopicUpdate", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsTopicUpdateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsTopicUpdateResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsTopicUpdateResponse struct {
	RequestId string
	HelpUrl   string
}
