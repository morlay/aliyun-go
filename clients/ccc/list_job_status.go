package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListJobStatusRequest struct {
	requests.RpcRequest
	ContactName   string `position:"Query" name:"ContactName"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	TimeAlignment string `position:"Query" name:"TimeAlignment"`
	GroupId       string `position:"Query" name:"GroupId"`
	PhoneNumber   string `position:"Query" name:"PhoneNumber"`
	PageSize      int    `position:"Query" name:"PageSize"`
	EndTime       int64  `position:"Query" name:"EndTime"`
	StartTime     int64  `position:"Query" name:"StartTime"`
	ScenarioId    string `position:"Query" name:"ScenarioId"`
	PageNumber    int    `position:"Query" name:"PageNumber"`
}

func (req *ListJobStatusRequest) Invoke(client *sdk.Client) (resp *ListJobStatusResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListJobStatus", "ccc", "")
	resp = &ListJobStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListJobStatusResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	Jobs           ListJobStatusJobs
}

type ListJobStatusJobs struct {
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	List       ListJobStatusJobList
}

type ListJobStatusJob struct {
	JobId          common.String
	GroupId        common.String
	ScenarioId     common.String
	StrategyId     common.String
	Priority       common.Integer
	Status         common.String
	ReferenceId    common.String
	FailureReason  common.String
	Contacts       ListJobStatusContactList
	Extras         ListJobStatusKeyValuePairList
	Tasks          ListJobStatusTaskList
	Summary        ListJobStatusSummaryItemList
	CallingNumbers ListJobStatusCallingNumberList
}

type ListJobStatusContact struct {
	ContactId   common.String
	ContactName common.String
	Honorific   common.String
	Role        common.String
	PhoneNumber common.String
	State       common.String
	ReferenceId common.String
	JobId       common.String
}

type ListJobStatusKeyValuePair struct {
	Key   common.String
	Value common.String
}

type ListJobStatusTask struct {
	TaskId        common.String
	JobId         common.String
	ScenarioId    common.String
	ChatbotId     common.String
	PlanedTime    common.Long
	ActualTime    common.Long
	CallingNumber common.String
	CalledNumber  common.String
	CallId        common.String
	Status        common.String
	Brief         common.String
	Duration      common.Integer
	Contact1      ListJobStatusContact1
}

type ListJobStatusContact1 struct {
	ContactId   common.String
	ContactName common.String
	Honorific   common.String
	Role        common.String
	PhoneNumber common.String
	State       common.String
	ReferenceId common.String
	JobId       common.String
}

type ListJobStatusSummaryItem struct {
	SummaryId            common.String
	GroupId              common.String
	JobId                common.String
	TaskId               common.String
	ConversationDetailId common.String
	Category             common.String
	SummaryName          common.String
	Content              common.String
}

type ListJobStatusJobList []ListJobStatusJob

func (list *ListJobStatusJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobStatusJob)
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

type ListJobStatusContactList []ListJobStatusContact

func (list *ListJobStatusContactList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobStatusContact)
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

type ListJobStatusKeyValuePairList []ListJobStatusKeyValuePair

func (list *ListJobStatusKeyValuePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobStatusKeyValuePair)
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

type ListJobStatusTaskList []ListJobStatusTask

func (list *ListJobStatusTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobStatusTask)
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

type ListJobStatusSummaryItemList []ListJobStatusSummaryItem

func (list *ListJobStatusSummaryItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobStatusSummaryItem)
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

type ListJobStatusCallingNumberList []common.String

func (list *ListJobStatusCallingNumberList) UnmarshalJSON(data []byte) error {
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
