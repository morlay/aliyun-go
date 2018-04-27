package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	Scenario       CreateScenarioScenario
}

type CreateScenarioScenario struct {
	ScenarioId          string
	ScenarioName        string
	ScenarioDescription string
	Type                string
	IsTemplate          bool
	Surveys             CreateScenarioSurveyList
	Variables           CreateScenarioKeyValuePairList
	Strategy            CreateScenarioStrategy
}

type CreateScenarioSurvey struct {
	SurveyId          string
	SurveyName        string
	SurveyDescription string
	Role              string
	Round             int
	BeebotId          string
	Intents           CreateScenarioIntentNodeList
}

type CreateScenarioIntentNode struct {
	NodeId   string
	IntentId string
}

type CreateScenarioKeyValuePair struct {
	Key   string
	Value string
}

type CreateScenarioStrategy struct {
	StrategyId          string
	StrategyName        string
	StrategyDescription string
	Type                string
	StartTime           int64
	EndTime             int64
	RepeatBy            string
	MaxAttemptsPerDay   int
	MinAttemptInterval  int
	Customized          string
	RoutingStrategy     string
	FollowUpStrategy    string
	IsTemplate          bool
	WorkingTime         CreateScenarioTimeFrameList
	RepeatDays          CreateScenarioRepeatDayList
}

type CreateScenarioTimeFrame struct {
	BeginTime string
	EndTime   string
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

type CreateScenarioRepeatDayList []string

func (list *CreateScenarioRepeatDayList) UnmarshalJSON(data []byte) error {
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
