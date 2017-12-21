package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsConsumerAccumulateRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Detail       string `position:"Query" name:"Detail"`
}

func (r OnsConsumerAccumulateRequest) Invoke(client *sdk.Client) (response *OnsConsumerAccumulateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OnsConsumerAccumulateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsConsumerAccumulate", "", "")

	resp := struct {
		*responses.BaseResponse
		OnsConsumerAccumulateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OnsConsumerAccumulateResponse

	err = client.DoAction(&req, &resp)
	return
}

type OnsConsumerAccumulateResponse struct {
	RequestId string
	HelpUrl   string
	Data      OnsConsumerAccumulateData
}

type OnsConsumerAccumulateData struct {
	Online            bool
	TotalDiff         int64
	ConsumeTps        float32
	LastTimestamp     int64
	DelayTime         int64
	DetailInTopicList OnsConsumerAccumulateDetailInTopicDoList
}

type OnsConsumerAccumulateDetailInTopicDo struct {
	Topic         string
	TotalDiff     int64
	LastTimestamp int64
	DelayTime     int64
}

type OnsConsumerAccumulateDetailInTopicDoList []OnsConsumerAccumulateDetailInTopicDo

func (list *OnsConsumerAccumulateDetailInTopicDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerAccumulateDetailInTopicDo)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
