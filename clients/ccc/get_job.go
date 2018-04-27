package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetJobRequest struct {
	requests.RpcRequest
	JobId      string `position:"Query" name:"JobId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *GetJobRequest) Invoke(client *sdk.Client) (resp *GetJobResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "GetJob", "ccc", "")
	resp = &GetJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetJobResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	Job            GetJobJob
}

type GetJobJob struct {
	JobId          string
	GroupId        string
	ScenarioId     string
	StrategyId     string
	Priority       int
	SystemPriority int
	Status         string
	ReferenceId    string
	FailureReason  string
	Contacts       GetJobContactList
	Extras         GetJobKeyValuePairList
	Tasks          GetJobTaskList
	Summary        GetJobSummaryItem3List
	CallingNumbers GetJobCallingNumberList
}

type GetJobContact struct {
	ContactId   string
	ContactName string
	Honorific   string
	Role        string
	PhoneNumber string
	State       string
	ReferenceId string
}

type GetJobKeyValuePair struct {
	Key   string
	Value string
}

type GetJobTask struct {
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
	Conversation  GetJobConversationDetailList
	Contact2      GetJobContact2
}

type GetJobConversationDetail struct {
	Timestamp int64
	Speaker   string
	Script    string
	Summary1  GetJobSummaryItemList
}

type GetJobSummaryItem struct {
	Category    string
	SummaryName string
	Content     string
}

type GetJobContact2 struct {
	ContactId   string
	ContactName string
	Honorific   string
	Role        string
	PhoneNumber string
	State       string
	ReferenceId string
}

type GetJobSummaryItem3 struct {
	Category    string
	SummaryName string
	Content     string
}

type GetJobContactList []GetJobContact

func (list *GetJobContactList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobContact)
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

type GetJobKeyValuePairList []GetJobKeyValuePair

func (list *GetJobKeyValuePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobKeyValuePair)
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

type GetJobTaskList []GetJobTask

func (list *GetJobTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobTask)
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

type GetJobSummaryItem3List []GetJobSummaryItem3

func (list *GetJobSummaryItem3List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobSummaryItem3)
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

type GetJobCallingNumberList []string

func (list *GetJobCallingNumberList) UnmarshalJSON(data []byte) error {
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

type GetJobConversationDetailList []GetJobConversationDetail

func (list *GetJobConversationDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobConversationDetail)
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

type GetJobSummaryItemList []GetJobSummaryItem

func (list *GetJobSummaryItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobSummaryItem)
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
