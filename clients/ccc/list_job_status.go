package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	Jobs           ListJobStatusJobs
}

type ListJobStatusJobs struct {
	TotalCount int
	PageNumber int
	PageSize   int
	List       ListJobStatusJobList
}

type ListJobStatusJob struct {
	JobId          string
	GroupId        string
	ScenarioId     string
	StrategyId     string
	Priority       int
	Status         string
	ReferenceId    string
	FailureReason  string
	Contacts       ListJobStatusContactList
	Extras         ListJobStatusKeyValuePairList
	Tasks          ListJobStatusTaskList
	Summary        ListJobStatusSummaryItemList
	CallingNumbers ListJobStatusCallingNumberList
}

type ListJobStatusContact struct {
	ContactId   string
	ContactName string
	Honorific   string
	Role        string
	PhoneNumber string
	State       string
	ReferenceId string
	JobId       string
}

type ListJobStatusKeyValuePair struct {
	Key   string
	Value string
}

type ListJobStatusTask struct {
	TaskId        string
	JobId         string
	ScenarioId    string
	ChatbotId     string
	PlanedTime    int64
	ActualTime    int64
	CallingNumber string
	CalledNumber  string
	CallId        string
	Status        string
	Brief         string
	Duration      int
	Contact1      ListJobStatusContact1
}

type ListJobStatusContact1 struct {
	ContactId   string
	ContactName string
	Honorific   string
	Role        string
	PhoneNumber string
	State       string
	ReferenceId string
	JobId       string
}

type ListJobStatusSummaryItem struct {
	SummaryId            string
	GroupId              string
	JobId                string
	TaskId               string
	ConversationDetailId string
	Category             string
	SummaryName          string
	Content              string
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

type ListJobStatusCallingNumberList []string

func (list *ListJobStatusCallingNumberList) UnmarshalJSON(data []byte) error {
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
