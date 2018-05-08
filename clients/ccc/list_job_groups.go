package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	JobGroups      ListJobGroupsJobGroups
}

type ListJobGroupsJobGroups struct {
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	List       ListJobGroupsJobGroupList
}

type ListJobGroupsJobGroup struct {
	JobGroupId          common.String
	JobGroupName        common.String
	JobGroupDescription common.String
	ScenarioId          common.String
	JobFilePath         common.String
	CreationTime        common.Long
	CallingNumbers      ListJobGroupsCallingNumberList
	Strategy            ListJobGroupsStrategy
	Progress            ListJobGroupsProgress
}

type ListJobGroupsStrategy struct {
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
	WorkingTime         ListJobGroupsTimeFrameList
	RepeatDays          ListJobGroupsRepeatDayList
}

type ListJobGroupsTimeFrame struct {
	From common.String
	To   common.String
}

type ListJobGroupsProgress struct {
	TotalJobs        common.Integer
	Status           common.String
	TotalNotAnswered common.Integer
	TotalCompleted   common.Integer
	StartTime        common.Long
	Duration         common.Integer
	Categories       ListJobGroupsKeyValuePairList
}

type ListJobGroupsKeyValuePair struct {
	Key   common.String
	Value common.String
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

type ListJobGroupsCallingNumberList []common.String

func (list *ListJobGroupsCallingNumberList) UnmarshalJSON(data []byte) error {
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

type ListJobGroupsRepeatDayList []common.String

func (list *ListJobGroupsRepeatDayList) UnmarshalJSON(data []byte) error {
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
