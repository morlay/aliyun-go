package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListRecordingsByContactIdRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	ContactId  string `position:"Query" name:"ContactId"`
}

func (req *ListRecordingsByContactIdRequest) Invoke(client *sdk.Client) (resp *ListRecordingsByContactIdResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "ListRecordingsByContactId", "ccc", "")
	resp = &ListRecordingsByContactIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListRecordingsByContactIdResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	Recordings     ListRecordingsByContactIdRecordingList
}

type ListRecordingsByContactIdRecording struct {
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
