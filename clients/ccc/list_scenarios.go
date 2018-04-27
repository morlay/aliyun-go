package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListScenariosRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *ListScenariosRequest) Invoke(client *sdk.Client) (resp *ListScenariosResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListScenarios", "ccc", "")
	resp = &ListScenariosResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListScenariosResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	Scenarios      ListScenariosScenarioList
}

type ListScenariosScenario struct {
	ScenarioId          string
	ScenarioName        string
	ScenarioDescription string
	Type                string
	IsTemplate          bool
	Surveys             ListScenariosSurveyList
	Variables           ListScenariosKeyValuePairList
	Strategy            ListScenariosStrategy
}

type ListScenariosSurvey struct {
	SurveyId          string
	SurveyName        string
	SurveyDescription string
	Role              string
	Round             int
	BeebotId          string
	Intents           ListScenariosIntentNodeList
}

type ListScenariosIntentNode struct {
	NodeId   string
	IntentId string
}

type ListScenariosKeyValuePair struct {
	Key   string
	Value string
}

type ListScenariosStrategy struct {
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
	WorkingTime         ListScenariosTimeFrameList
	RepeatDays          ListScenariosRepeatDayList
}

type ListScenariosTimeFrame struct {
	BeginTime string
	EndTime   string
}

type ListScenariosScenarioList []ListScenariosScenario

func (list *ListScenariosScenarioList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListScenariosScenario)
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

type ListScenariosSurveyList []ListScenariosSurvey

func (list *ListScenariosSurveyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListScenariosSurvey)
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

type ListScenariosKeyValuePairList []ListScenariosKeyValuePair

func (list *ListScenariosKeyValuePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListScenariosKeyValuePair)
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

type ListScenariosIntentNodeList []ListScenariosIntentNode

func (list *ListScenariosIntentNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListScenariosIntentNode)
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

type ListScenariosTimeFrameList []ListScenariosTimeFrame

func (list *ListScenariosTimeFrameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListScenariosTimeFrame)
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

type ListScenariosRepeatDayList []string

func (list *ListScenariosRepeatDayList) UnmarshalJSON(data []byte) error {
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
