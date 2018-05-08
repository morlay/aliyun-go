package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	Recordings     ListRecordingsRecordings
}

type ListRecordingsRecordings struct {
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	List       ListRecordingsRecordingList
}

type ListRecordingsRecording struct {
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
