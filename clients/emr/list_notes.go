package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListNotesRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
}

func (req *ListNotesRequest) Invoke(client *sdk.Client) (resp *ListNotesResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListNotes", "", "")
	resp = &ListNotesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListNotesResponse struct {
	responses.BaseResponse
	RequestId string
	Notes     ListNotesNoteInfoList
}

type ListNotesNoteInfo struct {
	Id   string
	Name string
	Type string
}

type ListNotesNoteInfoList []ListNotesNoteInfo

func (list *ListNotesNoteInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListNotesNoteInfo)
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
