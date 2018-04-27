package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListRecordingsRequest struct {
	requests.RpcRequest
	AgentId     string `position:"Query" name:"AgentId"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	Criteria    string `position:"Query" name:"Criteria"`
	PhoneNumber string `position:"Query" name:"PhoneNumber"`
	PageSize    int    `position:"Query" name:"PageSize"`
	StopTime    int64  `position:"Query" name:"StopTime"`
	StartTime   int64  `position:"Query" name:"StartTime"`
	PageNumber  int    `position:"Query" name:"PageNumber"`
}

func (req *ListRecordingsRequest) Invoke(client *sdk.Client) (resp *ListRecordingsResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListRecordings", "ccc", "")
	resp = &ListRecordingsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListRecordingsResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	Recordings     ListRecordingsRecordings
}

type ListRecordingsRecordings struct {
	TotalCount int
	PageNumber int
	PageSize   int
	List       ListRecordingsRecordingList
}

type ListRecordingsRecording struct {
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

type ListRecordingsRecordingList []ListRecordingsRecording

func (list *ListRecordingsRecordingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRecordingsRecording)
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
