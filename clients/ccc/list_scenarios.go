package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	Scenarios      ListScenariosScenarioList
}

type ListScenariosScenario struct {
	ScenarioId          common.String
	ScenarioName        common.String
	ScenarioDescription common.String
	Type                common.String
	IsTemplate          bool
	Surveys             ListScenariosSurveyList
	Variables           ListScenariosKeyValuePairList
	Strategy            ListScenariosStrategy
}

type ListScenariosSurvey struct {
	SurveyId          common.String
	SurveyName        common.String
	SurveyDescription common.String
	Role              common.String
	Round             common.Integer
	BeebotId          common.String
	Intents           ListScenariosIntentNodeList
}

type ListScenariosIntentNode struct {
	NodeId   common.String
	IntentId common.String
}

type ListScenariosKeyValuePair struct {
	Key   common.String
	Value common.String
}

type ListScenariosStrategy struct {
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
	WorkingTime         ListScenariosTimeFrameList
	RepeatDays          ListScenariosRepeatDayList
}

type ListScenariosTimeFrame struct {
	BeginTime common.String
	EndTime   common.String
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

type ListScenariosRepeatDayList []common.String

func (list *ListScenariosRepeatDayList) UnmarshalJSON(data []byte) error {
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
