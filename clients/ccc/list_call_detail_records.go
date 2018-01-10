package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListCallDetailRecordsRequest struct {
	requests.RpcRequest
	InstanceId  string `position:"Query" name:"InstanceId"`
	Criteria    string `position:"Query" name:"Criteria"`
	PhoneNumber string `position:"Query" name:"PhoneNumber"`
	PageSize    int    `position:"Query" name:"PageSize"`
	StartTime   int64  `position:"Query" name:"StartTime"`
	StopTime    int64  `position:"Query" name:"StopTime"`
	PageNumber  int    `position:"Query" name:"PageNumber"`
}

func (req *ListCallDetailRecordsRequest) Invoke(client *sdk.Client) (resp *ListCallDetailRecordsResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListCallDetailRecords", "ccc", "")
	resp = &ListCallDetailRecordsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListCallDetailRecordsResponse struct {
	responses.BaseResponse
	RequestId         string
	Success           bool
	Code              string
	Message           string
	HttpStatusCode    int
	CallDetailRecords ListCallDetailRecordsCallDetailRecords
}

type ListCallDetailRecordsCallDetailRecords struct {
	TotalCount int
	PageNumber int
	PageSize   int
	List       ListCallDetailRecordsCallDetailRecordList
}

type ListCallDetailRecordsCallDetailRecord struct {
	ContactId          string
	StartTime          int64
	Duration           int
	ContactType        string
	ContactDisposition string
	CallingNumber      string
	CalledNumber       string
	AgentNames         string
	SkillGroupNames    string
	InstanceId         string
	Agents             ListCallDetailRecordsCallDetailAgentList
	Recordings         ListCallDetailRecordsRecordingList
}

type ListCallDetailRecordsCallDetailAgent struct {
	ContactId      string
	AgentId        string
	AgentName      string
	SkillGroupName string
	QueueTime      int
	RingTime       int
	StartTime      int64
	TalkTime       int
	HoldTime       int
	WorkTime       int
}

type ListCallDetailRecordsRecording struct {
	ContactId       string
	ContactType     string
	AgentId         string
	AgentName       string
	CallingNumber   string
	CalledNumber    string
	StartTime       int64
	Duration        int
	FileName        string
	FilePath        string
	FileDescription string
	Channel         string
	InstanceId      string
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
