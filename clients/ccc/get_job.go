package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	Job            GetJobJob
}

type GetJobJob struct {
	JobId          common.String
	GroupId        common.String
	ScenarioId     common.String
	StrategyId     common.String
	Priority       common.Integer
	SystemPriority common.Integer
	Status         common.String
	ReferenceId    common.String
	FailureReason  common.String
	Contacts       GetJobContactList
	Extras         GetJobKeyValuePairList
	Tasks          GetJobTaskList
	Summary        GetJobSummaryItem3List
	CallingNumbers GetJobCallingNumberList
}

type GetJobContact struct {
	ContactId   common.String
	ContactName common.String
	Honorific   common.String
	Role        common.String
	PhoneNumber common.String
	State       common.String
	ReferenceId common.String
}

type GetJobKeyValuePair struct {
	Key   common.String
	Value common.String
}

type GetJobTask struct {
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
	Conversation  GetJobConversationDetailList
	Contact2      GetJobContact2
}

type GetJobConversationDetail struct {
	Timestamp common.Long
	Speaker   common.String
	Script    common.String
	Summary1  GetJobSummaryItemList
}

type GetJobSummaryItem struct {
	Category    common.String
	SummaryName common.String
	Content     common.String
}

type GetJobContact2 struct {
	ContactId   common.String
	ContactName common.String
	Honorific   common.String
	Role        common.String
	PhoneNumber common.String
	State       common.String
	ReferenceId common.String
}

type GetJobSummaryItem3 struct {
	Category    common.String
	SummaryName common.String
	Content     common.String
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

type GetJobCallingNumberList []common.String

func (list *GetJobCallingNumberList) UnmarshalJSON(data []byte) error {
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
