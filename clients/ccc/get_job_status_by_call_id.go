package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	Job            GetJobStatusByCallIdJob
}

type GetJobStatusByCallIdJob struct {
	JobId          common.String
	GroupId        common.String
	ScenarioId     common.String
	StrategyId     common.String
	Priority       common.Integer
	Status         common.String
	ReferenceId    common.String
	FailureReason  common.String
	Contacts       GetJobStatusByCallIdContactList
	Extras         GetJobStatusByCallIdKeyValuePairList
	Tasks          GetJobStatusByCallIdTaskList
	Summary        GetJobStatusByCallIdSummaryItem3List
	CallingNumbers GetJobStatusByCallIdCallingNumberList
}

type GetJobStatusByCallIdContact struct {
	ContactId   common.String
	ContactName common.String
	Honorific   common.String
	Role        common.String
	PhoneNumber common.String
	State       common.String
	ReferenceId common.String
	JobId       common.String
}

type GetJobStatusByCallIdKeyValuePair struct {
	Key   common.String
	Value common.String
}

type GetJobStatusByCallIdTask struct {
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
	Conversation  GetJobStatusByCallIdConversationDetailList
	Contact2      GetJobStatusByCallIdContact2
}

type GetJobStatusByCallIdConversationDetail struct {
	ConversationDetailId common.String
	TaskId               common.String
	Timestamp            common.Long
	Speaker              common.String
	Script               common.String
	Summary1             GetJobStatusByCallIdSummaryItemList
}

type GetJobStatusByCallIdSummaryItem struct {
	SummaryId   common.String
	Category    common.String
	SummaryName common.String
	Content     common.String
}

type GetJobStatusByCallIdContact2 struct {
	ContactId   common.String
	ContactName common.String
	Honorific   common.String
	Role        common.String
	PhoneNumber common.String
	State       common.String
	ReferenceId common.String
	JobId       common.String
}

type GetJobStatusByCallIdSummaryItem3 struct {
	SummaryId   common.String
	Category    common.String
	SummaryName common.String
	Content     common.String
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

type GetJobStatusByCallIdCallingNumberList []common.String

func (list *GetJobStatusByCallIdCallingNumberList) UnmarshalJSON(data []byte) error {
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
