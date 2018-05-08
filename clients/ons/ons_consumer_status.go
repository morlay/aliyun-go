package ons

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	HelpUrl   common.String
	Data      OnsConsumerStatusData
}

type OnsConsumerStatusData struct {
	Online                     bool
	TotalDiff                  common.Long
	ConsumeTps                 common.Float
	LastTimestamp              common.Long
	DelayTime                  common.Long
	ConsumeModel               common.String
	SubscriptionSame           bool
	RebalanceOK                bool
	ConnectionSet              OnsConsumerStatusConnectionDoList
	DetailInTopicList          OnsConsumerStatusDetailInTopicDoList
	ConsumerConnectionInfoList OnsConsumerStatusConsumerConnectionInfoDoList
}

type OnsConsumerStatusConnectionDo struct {
	ClientId   common.String
	ClientAddr common.String
	Language   common.String
	Version    common.String
}

type OnsConsumerStatusDetailInTopicDo struct {
	Topic         common.String
	TotalDiff     common.Long
	LastTimestamp common.Long
	DelayTime     common.Long
}

type OnsConsumerStatusConsumerConnectionInfoDo struct {
	ClientId        common.String
	Connection      common.String
	Language        common.String
	Version         common.String
	ConsumeModel    common.String
	ConsumeType     common.String
	ThreadCount     common.Integer
	StartTimeStamp  common.Long
	LastTimeStamp   common.Long
	SubscriptionSet OnsConsumerStatusSubscriptionDataList
	RunningDataList OnsConsumerStatusConsumerRunningDataDoList
	Jstack          OnsConsumerStatusThreadTrackDoList
}

type OnsConsumerStatusSubscriptionData struct {
	Topic      common.String
	SubString  common.String
	SubVersion common.Long
	TagsSet    OnsConsumerStatusTagsSetList
}

type OnsConsumerStatusConsumerRunningDataDo struct {
	ConsumerId         common.String
	Topic              common.String
	Rt                 common.Float
	OkTps              common.Float
	FailedTps          common.Float
	FailedCountPerHour common.Long
}

type OnsConsumerStatusThreadTrackDo struct {
	Thread    common.String
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

type OnsConsumerStatusTagsSetList []common.String

func (list *OnsConsumerStatusTagsSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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

type OnsConsumerStatusTrackListList []common.String

func (list *OnsConsumerStatusTrackListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
