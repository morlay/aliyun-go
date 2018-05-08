package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateScenarioRequest struct {
	requests.RpcRequest
	InstanceId   string                         `position:"Query" name:"InstanceId"`
	SurveysJsons *CreateScenarioSurveysJsonList `position:"Query" type:"Repeated" name:"SurveysJson"`
	StrategyJson string                         `position:"Query" name:"StrategyJson"`
	Name         string                         `position:"Query" name:"Name"`
	Description  string                         `position:"Query" name:"Description"`
}

func (req *CreateScenarioRequest) Invoke(client *sdk.Client) (resp *CreateScenarioResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "CreateScenario", "ccc", "")
	resp = &CreateScenarioResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateScenarioResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	Scenario       CreateScenarioScenario
}

type CreateScenarioScenario struct {
	ScenarioId          common.String
	ScenarioName        common.String
	ScenarioDescription common.String
	Type                common.String
	IsTemplate          bool
	Surveys             CreateScenarioSurveyList
	Variables           CreateScenarioKeyValuePairList
	Strategy            CreateScenarioStrategy
}

type CreateScenarioSurvey struct {
	SurveyId          common.String
	SurveyName        common.String
	SurveyDescription common.String
	Role              common.String
	Round             common.Integer
	BeebotId          common.String
	Intents           CreateScenarioIntentNodeList
}

type CreateScenarioIntentNode struct {
	NodeId   common.String
	IntentId common.String
}

type CreateScenarioKeyValuePair struct {
	Key   common.String
	Value common.String
}

type CreateScenarioStrategy struct {
	StrategyId          common.String
	StrategyName        common.String
	StrategyDescription common.String
	Type                common.String
	StartTime           common.Long
	EndTime             common.Long
	RepeatBy            common.String
	MaxAttemptsPerDay   common.Integer
	MinAttemptInterval  common.Integer
	Customized          common.String
	RoutingStrategy     common.String
	FollowUpStrategy    common.String
	IsTemplate          bool
	WorkingTime         CreateScenarioTimeFrameList
	RepeatDays          CreateScenarioRepeatDayList
}

type CreateScenarioTimeFrame struct {
	BeginTime common.String
	EndTime   common.String
}

type CreateScenarioSurveysJsonList []string

func (list *CreateScenarioSurveysJsonList) UnmarshalJSON(data []byte) error {
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

type CreateScenarioSurveyList []CreateScenarioSurvey

func (list *CreateScenarioSurveyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateScenarioSurvey)
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

type CreateScenarioKeyValuePairList []CreateScenarioKeyValuePair

func (list *CreateScenarioKeyValuePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateScenarioKeyValuePair)
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

type CreateScenarioIntentNodeList []CreateScenarioIntentNode

func (list *CreateScenarioIntentNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateScenarioIntentNode)
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

type CreateScenarioTimeFrameList []CreateScenarioTimeFrame

func (list *CreateScenarioTimeFrameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateScenarioTimeFrame)
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

type CreateScenarioRepeatDayList []common.String

func (list *CreateScenarioRepeatDayList) UnmarshalJSON(data []byte) error {
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
