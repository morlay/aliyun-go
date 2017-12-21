package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsTopicStatusRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
}

func (r OnsTopicStatusRequest) Invoke(client *sdk.Client) (response *OnsTopicStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsTopicStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsTopicStatus", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsTopicStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsTopicStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsTopicStatusResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsTopicStatusData
}

type OnsTopicStatusData struct {
	TotalCount    int64
	LastTimeStamp int64
}
