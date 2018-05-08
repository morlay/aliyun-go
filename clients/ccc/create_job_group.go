package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	JobGroup       CreateJobGroupJobGroup
}

type CreateJobGroupJobGroup struct {
	JobGroupId          common.String
	JobGroupName        common.String
	JobGroupDescription common.String
	ScenarioId          common.String
	JobFilePath         common.String
	CreationTime        common.Long
	CallingNumbers      CreateJobGroupCallingNumberList
	Strategy            CreateJobGroupStrategy
}

type CreateJobGroupStrategy struct {
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
	WorkingTime         CreateJobGroupTimeFrameList
	RepeatDays          CreateJobGroupRepeatDayList
}

type CreateJobGroupTimeFrame struct {
	From common.String
	To   common.String
}

type CreateJobGroupCallingNumberList []common.String

func (list *CreateJobGroupCallingNumberList) UnmarshalJSON(data []byte) error {
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

type CreateJobGroupRepeatDayList []common.String

func (list *CreateJobGroupRepeatDayList) UnmarshalJSON(data []byte) error {
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
