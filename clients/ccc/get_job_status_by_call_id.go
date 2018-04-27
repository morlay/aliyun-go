package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetJobStatusByCallIdRequest struct {
	requests.RpcRequest
	CallId     string `position:"Query" name:"CallId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *GetJobStatusByCallIdRequest) Invoke(client *sdk.Client) (resp *GetJobStatusByCallIdResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "GetJobStatusByCallId", "ccc", "")
	resp = &GetJobStatusByCallIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetJobStatusByCallIdResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	Job            GetJobStatusByCallIdJob
}

type GetJobStatusByCallIdJob struct {
	JobId          string
	GroupId        string
	ScenarioId     string
	StrategyId     string
	Priority       int
	Status         string
	ReferenceId    string
	FailureReason  string
	Contacts       GetJobStatusByCallIdContactList
	Extras         GetJobStatusByCallIdKeyValuePairList
	Tasks          GetJobStatusByCallIdTaskList
	Summary        GetJobStatusByCallIdSummaryItem3List
	CallingNumbers GetJobStatusByCallIdCallingNumberList
}

type GetJobStatusByCallIdContact struct {
	ContactId   string
	ContactName string
	Honorific   string
	Role        string
	PhoneNumber string
	State       string
	ReferenceId string
	JobId       string
}

type GetJobStatusByCallIdKeyValuePair struct {
	Key   string
	Value string
}

type GetJobStatusByCallIdTask struct {
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
	Conversation  GetJobStatusByCallIdConversationDetailList
	Contact2      GetJobStatusByCallIdContact2
}

type GetJobStatusByCallIdConversationDetail struct {
	ConversationDetailId string
	TaskId               string
	Timestamp            int64
	Speaker              string
	Script               string
	Summary1             GetJobStatusByCallIdSummaryItemList
}

type GetJobStatusByCallIdSummaryItem struct {
	SummaryId   string
	Category    string
	SummaryName string
	Content     string
}

type GetJobStatusByCallIdContact2 struct {
	ContactId   string
	ContactName string
	Honorific   string
	Role        string
	PhoneNumber string
	State       string
	ReferenceId string
	JobId       string
}

type GetJobStatusByCallIdSummaryItem3 struct {
	SummaryId   string
	Category    string
	SummaryName string
	Content     string
}

type GetJobStatusByCallIdContactList []GetJobStatusByCallIdContact

func (list *GetJobStatusByCallIdContactList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobStatusByCallIdContact)
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

type GetJobStatusByCallIdKeyValuePairList []GetJobStatusByCallIdKeyValuePair

func (list *GetJobStatusByCallIdKeyValuePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobStatusByCallIdKeyValuePair)
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

type GetJobStatusByCallIdTaskList []GetJobStatusByCallIdTask

func (list *GetJobStatusByCallIdTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobStatusByCallIdTask)
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

type GetJobStatusByCallIdSummaryItem3List []GetJobStatusByCallIdSummaryItem3

func (list *GetJobStatusByCallIdSummaryItem3List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobStatusByCallIdSummaryItem3)
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

type GetJobStatusByCallIdCallingNumberList []string

func (list *GetJobStatusByCallIdCallingNumberList) UnmarshalJSON(data []byte) error {
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

type GetJobStatusByCallIdConversationDetailList []GetJobStatusByCallIdConversationDetail

func (list *GetJobStatusByCallIdConversationDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobStatusByCallIdConversationDetail)
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

type GetJobStatusByCallIdSummaryItemList []GetJobStatusByCallIdSummaryItem

func (list *GetJobStatusByCallIdSummaryItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobStatusByCallIdSummaryItem)
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
