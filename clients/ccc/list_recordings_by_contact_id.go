package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListRecordingsByContactIdRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	ContactId  string `position:"Query" name:"ContactId"`
}

func (req *ListRecordingsByContactIdRequest) Invoke(client *sdk.Client) (resp *ListRecordingsByContactIdResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListRecordingsByContactId", "CCC", "")
	resp = &ListRecordingsByContactIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListRecordingsByContactIdResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	Recordings     ListRecordingsByContactIdRecordingList
}

type ListRecordingsByContactIdRecording struct {
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

type ListRecordingsByContactIdRecordingList []ListRecordingsByContactIdRecording

func (list *ListRecordingsByContactIdRecordingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRecordingsByContactIdRecording)
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
