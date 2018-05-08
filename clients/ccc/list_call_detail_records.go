package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListCallDetailRecordsRequest struct {
	requests.RpcRequest
	InstanceId         string `position:"Query" name:"InstanceId"`
	ContactDisposition string `position:"Query" name:"ContactDisposition"`
	ContactType        string `position:"Query" name:"ContactType"`
	Criteria           string `position:"Query" name:"Criteria"`
	PhoneNumber        string `position:"Query" name:"PhoneNumber"`
	PageSize           int    `position:"Query" name:"PageSize"`
	OrderBy            string `position:"Query" name:"OrderBy"`
	StopTime           int64  `position:"Query" name:"StopTime"`
	StartTime          int64  `position:"Query" name:"StartTime"`
	PageNumber         int    `position:"Query" name:"PageNumber"`
	WithRecording      string `position:"Query" name:"WithRecording"`
}

func (req *ListCallDetailRecordsRequest) Invoke(client *sdk.Client) (resp *ListCallDetailRecordsResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListCallDetailRecords", "ccc", "")
	resp = &ListCallDetailRecordsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListCallDetailRecordsResponse struct {
	responses.BaseResponse
	RequestId         common.String
	Success           bool
	Code              common.String
	Message           common.String
	HttpStatusCode    common.Integer
	CallDetailRecords ListCallDetailRecordsCallDetailRecords
}

type ListCallDetailRecordsCallDetailRecords struct {
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	List       ListCallDetailRecordsCallDetailRecordList
}

type ListCallDetailRecordsCallDetailRecord struct {
	ContactId          common.String
	StartTime          common.Long
	Duration           common.Integer
	Satisfaction       common.Integer
	ContactType        common.String
	ContactDisposition common.String
	CallingNumber      common.String
	CalledNumber       common.String
	AgentNames         common.String
	SkillGroupNames    common.String
	InstanceId         common.String
	ExtraAttr          common.String
	Agents             ListCallDetailRecordsCallDetailAgentList
	Recordings         ListCallDetailRecordsRecordingList
}

type ListCallDetailRecordsCallDetailAgent struct {
	ContactId      common.String
	AgentId        common.String
	AgentName      common.String
	SkillGroupName common.String
	QueueTime      common.Integer
	RingTime       common.Integer
	StartTime      common.Long
	TalkTime       common.Integer
	HoldTime       common.Integer
	WorkTime       common.Integer
}

type ListCallDetailRecordsRecording struct {
	ContactId       common.String
	ContactType     common.String
	AgentId         common.String
	AgentName       common.String
	CallingNumber   common.String
	CalledNumber    common.String
	StartTime       common.Long
	Duration        common.Integer
	FileName        common.String
	FilePath        common.String
	FileDescription common.String
	Channel         common.String
	InstanceId      common.String
}

type ListCallDetailRecordsCallDetailRecordList []ListCallDetailRecordsCallDetailRecord

func (list *ListCallDetailRecordsCallDetailRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListCallDetailRecordsCallDetailRecord)
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

type ListCallDetailRecordsCallDetailAgentList []ListCallDetailRecordsCallDetailAgent

func (list *ListCallDetailRecordsCallDetailAgentList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListCallDetailRecordsCallDetailAgent)
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

type ListCallDetailRecordsRecordingList []ListCallDetailRecordsRecording

func (list *ListCallDetailRecordsRecordingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListCallDetailRecordsRecording)
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
