package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnsConsumerStatusRequest struct {
	requests.RpcRequest
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	NeedJstack   string `position:"Query" name:"NeedJstack"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Detail       string `position:"Query" name:"Detail"`
}

func (req *OnsConsumerStatusRequest) Invoke(client *sdk.Client) (resp *OnsConsumerStatusResponse, err error) {
	req.InitWithApiInfo("Ons", "2017-09-18", "OnsConsumerStatus", "", "")
	resp = &OnsConsumerStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnsConsumerStatusResponse struct {
	responses.BaseResponse
	RequestId string
	HelpUrl   string
	Data      OnsConsumerStatusData
}

type OnsConsumerStatusData struct {
	Online                     bool
	TotalDiff                  int64
	ConsumeTps                 float32
	LastTimestamp              int64
	DelayTime                  int64
	ConsumeModel               string
	SubscriptionSame           bool
	RebalanceOK                bool
	ConnectionSet              OnsConsumerStatusConnectionDoList
	DetailInTopicList          OnsConsumerStatusDetailInTopicDoList
	ConsumerConnectionInfoList OnsConsumerStatusConsumerConnectionInfoDoList
}

type OnsConsumerStatusConnectionDo struct {
	ClientId   string
	ClientAddr string
	Language   string
	Version    string
}

type OnsConsumerStatusDetailInTopicDo struct {
	Topic         string
	TotalDiff     int64
	LastTimestamp int64
	DelayTime     int64
}

type OnsConsumerStatusConsumerConnectionInfoDo struct {
	ClientId        string
	Connection      string
	Language        string
	Version         string
	ConsumeModel    string
	ConsumeType     string
	ThreadCount     int
	StartTimeStamp  int64
	LastTimeStamp   int64
	SubscriptionSet OnsConsumerStatusSubscriptionDataList
	RunningDataList OnsConsumerStatusConsumerRunningDataDoList
	Jstack          OnsConsumerStatusThreadTrackDoList
}

type OnsConsumerStatusSubscriptionData struct {
	Topic      string
	SubString  string
	SubVersion int64
	TagsSet    OnsConsumerStatusTagsSetList
}

type OnsConsumerStatusConsumerRunningDataDo struct {
	ConsumerId         string
	Topic              string
	Rt                 float32
	OkTps              float32
	FailedTps          float32
	FailedCountPerHour int64
}

type OnsConsumerStatusThreadTrackDo struct {
	Thread    string
	TrackList OnsConsumerStatusTrackListList
}

type OnsConsumerStatusConnectionDoList []OnsConsumerStatusConnectionDo

func (list *OnsConsumerStatusConnectionDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerStatusConnectionDo)
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

type OnsConsumerStatusDetailInTopicDoList []OnsConsumerStatusDetailInTopicDo

func (list *OnsConsumerStatusDetailInTopicDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerStatusDetailInTopicDo)
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

type OnsConsumerStatusConsumerConnectionInfoDoList []OnsConsumerStatusConsumerConnectionInfoDo

func (list *OnsConsumerStatusConsumerConnectionInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerStatusConsumerConnectionInfoDo)
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

type OnsConsumerStatusSubscriptionDataList []OnsConsumerStatusSubscriptionData

func (list *OnsConsumerStatusSubscriptionDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerStatusSubscriptionData)
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

type OnsConsumerStatusConsumerRunningDataDoList []OnsConsumerStatusConsumerRunningDataDo

func (list *OnsConsumerStatusConsumerRunningDataDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerStatusConsumerRunningDataDo)
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

type OnsConsumerStatusThreadTrackDoList []OnsConsumerStatusThreadTrackDo

func (list *OnsConsumerStatusThreadTrackDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerStatusThreadTrackDo)
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

type OnsConsumerStatusTagsSetList []string

func (list *OnsConsumerStatusTagsSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type OnsConsumerStatusTrackListList []string

func (list *OnsConsumerStatusTrackListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
