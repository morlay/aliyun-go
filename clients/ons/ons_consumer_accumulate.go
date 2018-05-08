package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OnsConsumerAccumulateRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Detail       string `position:"Query" name:"Detail"`
}

func (req *OnsConsumerAccumulateRequest) Invoke(client *sdk.Client) (resp *OnsConsumerAccumulateResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsConsumerAccumulate", "", "")
	resp = &OnsConsumerAccumulateResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsConsumerAccumulateResponse struct {
	responses.BaseResponse
	RequestId common.String
	HelpUrl   common.String
	Data      OnsConsumerAccumulateData
}

type OnsConsumerAccumulateData struct {
	Online            bool
	TotalDiff         common.Long
	ConsumeTps        common.Float
	LastTimestamp     common.Long
	DelayTime         common.Long
	DetailInTopicList OnsConsumerAccumulateDetailInTopicDoList
}

type OnsConsumerAccumulateDetailInTopicDo struct {
	Topic         common.String
	TotalDiff     common.Long
	LastTimestamp common.Long
	DelayTime     common.Long
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
