package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListJobGroupsRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	PageSize   int    `position:"Query" name:"PageSize"`
	EndTime    int64  `position:"Query" name:"EndTime"`
	StartTime  int64  `position:"Query" name:"StartTime"`
	PageNumber int    `position:"Query" name:"PageNumber"`
}

func (req *ListJobGroupsRequest) Invoke(client *sdk.Client) (resp *ListJobGroupsResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListJobGroups", "ccc", "")
	resp = &ListJobGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListJobGroupsResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	JobGroups      ListJobGroupsJobGroups
}

type ListJobGroupsJobGroups struct {
	TotalCount int
	PageNumber int
	PageSize   int
	List       ListJobGroupsJobGroupList
}

type ListJobGroupsJobGroup struct {
	JobGroupId          string
	JobGroupName        string
	JobGroupDescription string
	ScenarioId          string
	JobFilePath         string
	CreationTime        int64
	CallingNumbers      ListJobGroupsCallingNumberList
	Strategy            ListJobGroupsStrategy
	Progress            ListJobGroupsProgress
}

type ListJobGroupsStrategy struct {
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
	WorkingTime         ListJobGroupsTimeFrameList
	RepeatDays          ListJobGroupsRepeatDayList
}

type ListJobGroupsTimeFrame struct {
	From string
	To   string
}

type ListJobGroupsProgress struct {
	TotalJobs        int
	Status           string
	TotalNotAnswered int
	TotalCompleted   int
	StartTime        int64
	Duration         int
	Categories       ListJobGroupsKeyValuePairList
}

type ListJobGroupsKeyValuePair struct {
	Key   string
	Value string
}

type ListJobGroupsJobGroupList []ListJobGroupsJobGroup

func (list *ListJobGroupsJobGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobGroupsJobGroup)
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

type ListJobGroupsCallingNumberList []string

func (list *ListJobGroupsCallingNumberList) UnmarshalJSON(data []byte) error {
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

type ListJobGroupsTimeFrameList []ListJobGroupsTimeFrame

func (list *ListJobGroupsTimeFrameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobGroupsTimeFrame)
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

type ListJobGroupsRepeatDayList []string

func (list *ListJobGroupsRepeatDayList) UnmarshalJSON(data []byte) error {
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

type ListJobGroupsKeyValuePairList []ListJobGroupsKeyValuePair

func (list *ListJobGroupsKeyValuePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobGroupsKeyValuePair)
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
