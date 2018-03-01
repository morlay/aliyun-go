package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListConvertOfficeFormatTasksRequest struct {
	requests.RpcRequest
	MaxKeys int    `position:"Query" name:"MaxKeys"`
	Marker  string `position:"Query" name:"Marker"`
	Project string `position:"Query" name:"Project"`
}

func (req *ListConvertOfficeFormatTasksRequest) Invoke(client *sdk.Client) (resp *ListConvertOfficeFormatTasksResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "ListConvertOfficeFormatTasks", "imm", "")
	resp = &ListConvertOfficeFormatTasksResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListConvertOfficeFormatTasksResponse struct {
	responses.BaseResponse
	RequestId  string
	NextMarker string
	Tasks      ListConvertOfficeFormatTasksTasksItemList
}

type ListConvertOfficeFormatTasksTasksItem struct {
	TaskId          string
	Status          string
	Percent         int
	PageCount       int
	SrcUri          string
	TgtType         string
	TgtUri          string
	ImageSpec       string
	NotifyTopicName string
	NotifyEndpoint  string
	ExternalID      string
	CreateTime      string
	FinishTime      string
}

type ListConvertOfficeFormatTasksTasksItemList []ListConvertOfficeFormatTasksTasksItem

func (list *ListConvertOfficeFormatTasksTasksItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListConvertOfficeFormatTasksTasksItem)
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
