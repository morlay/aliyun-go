package ons

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsTopicCreateRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	Cluster      string `position:"Query" name:"Cluster"`
	QueueNum     int    `position:"Query" name:"QueueNum"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	AppName      string `position:"Query" name:"AppName"`
	Qps          int64  `position:"Query" name:"Qps"`
	Topic        string `position:"Query" name:"Topic"`
	Remark       string `position:"Query" name:"Remark"`
	Appkey       string `position:"Query" name:"Appkey"`
	Order        string `position:"Query" name:"Order"`
	Status       int    `position:"Query" name:"Status"`
}

func (r OnsTopicCreateRequest) Invoke(client *sdk.Client) (response *OnsTopicCreateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsTopicCreateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsTopicCreate", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsTopicCreateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsTopicCreateResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsTopicCreateResponse struct {
	RequestId string
	HelpUrl   string
}
