package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsTrendTopicInputTpsRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	Period       int64  `position:"Query" name:"Period"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
	Type         int    `position:"Query" name:"Type"`
}

func (req *OnsTrendTopicInputTpsRequest) Invoke(client *sdk.Client) (resp *OnsTrendTopicInputTpsResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsTrendTopicInputTps", "", "")
	resp = &OnsTrendTopicInputTpsResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsTrendTopicInputTpsResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
	Data      OnsTrendTopicInputTpsData
}

type OnsTrendTopicInputTpsData struct {
	Title   string
	XUnit   string
	YUnit   string
	Records OnsTrendTopicInputTpsStatsDataDoList
}

type OnsTrendTopicInputTpsStatsDataDo struct {
	X int64
	Y float32
}

type OnsTrendTopicInputTpsStatsDataDoList []OnsTrendTopicInputTpsStatsDataDo

func (list *OnsTrendTopicInputTpsStatsDataDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTrendTopicInputTpsStatsDataDo)
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
