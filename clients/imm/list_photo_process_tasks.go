package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListPhotoProcessTasksRequest struct {
	requests.RpcRequest
	MaxKeys int    `position:"Query" name:"MaxKeys"`
	Marker  string `position:"Query" name:"Marker"`
	Project string `position:"Query" name:"Project"`
}

func (req *ListPhotoProcessTasksRequest) Invoke(client *sdk.Client) (resp *ListPhotoProcessTasksResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "ListPhotoProcessTasks", "imm", "")
	resp = &ListPhotoProcessTasksResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListPhotoProcessTasksResponse struct {
	responses.BaseResponse
	RequestId  string
	NextMarker string
	Tasks      ListPhotoProcessTasksTasksItemList
}

type ListPhotoProcessTasksTasksItem struct {
	TaskId          string
	Status          string
	SrcUri          string
	TgtUri          string
	Style           string
	NotifyTopicName string
	NotifyEndpoint  string
	ExternalID      string
	CreateTime      string
	FinishTime      string
	Percent         int
}

type ListPhotoProcessTasksTasksItemList []ListPhotoProcessTasksTasksItem

func (list *ListPhotoProcessTasksTasksItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPhotoProcessTasksTasksItem)
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
