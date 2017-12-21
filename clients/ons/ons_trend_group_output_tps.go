package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsTrendGroupOutputTpsRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	Period       int64  `position:"Query" name:"Period"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Topic        string `position:"Query" name:"Topic"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
	Type         int    `position:"Query" name:"Type"`
}

func (req *OnsTrendGroupOutputTpsRequest) Invoke(client *sdk.Client) (resp *OnsTrendGroupOutputTpsResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsTrendGroupOutputTps", "", "")
	resp = &OnsTrendGroupOutputTpsResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsTrendGroupOutputTpsResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
	Data      OnsTrendGroupOutputTpsData
}

type OnsTrendGroupOutputTpsData struct {
	Title   string
	XUnit   string
	YUnit   string
	Records OnsTrendGroupOutputTpsStatsDataDoList
}

type OnsTrendGroupOutputTpsStatsDataDo struct {
	X int64
	Y float32
}

type OnsTrendGroupOutputTpsStatsDataDoList []OnsTrendGroupOutputTpsStatsDataDo

func (list *OnsTrendGroupOutputTpsStatsDataDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTrendGroupOutputTpsStatsDataDo)
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
