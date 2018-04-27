package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateJobGroupRequest struct {
	requests.RpcRequest
	CallingNumbers *CreateJobGroupCallingNumberList `position:"Query" type:"Repeated" name:"CallingNumber"`
	InstanceId     string                           `position:"Query" name:"InstanceId"`
	StrategyJson   string                           `position:"Query" name:"StrategyJson"`
	Name           string                           `position:"Query" name:"Name"`
	Description    string                           `position:"Query" name:"Description"`
	ScenarioId     string                           `position:"Query" name:"ScenarioId"`
}

func (req *CreateJobGroupRequest) Invoke(client *sdk.Client) (resp *CreateJobGroupResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "CreateJobGroup", "ccc", "")
	resp = &CreateJobGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateJobGroupResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	JobGroup       CreateJobGroupJobGroup
}

type CreateJobGroupJobGroup struct {
	JobGroupId          string
	JobGroupName        string
	JobGroupDescription string
	ScenarioId          string
	JobFilePath         string
	CreationTime        int64
	CallingNumbers      CreateJobGroupCallingNumberList
	Strategy            CreateJobGroupStrategy
}

type CreateJobGroupStrategy struct {
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
	WorkingTime         CreateJobGroupTimeFrameList
	RepeatDays          CreateJobGroupRepeatDayList
}

type CreateJobGroupTimeFrame struct {
	From string
	To   string
}

type CreateJobGroupCallingNumberList []string

func (list *CreateJobGroupCallingNumberList) UnmarshalJSON(data []byte) error {
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

type CreateJobGroupTimeFrameList []CreateJobGroupTimeFrame

func (list *CreateJobGroupTimeFrameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateJobGroupTimeFrame)
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

type CreateJobGroupRepeatDayList []string

func (list *CreateJobGroupRepeatDayList) UnmarshalJSON(data []byte) error {
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
