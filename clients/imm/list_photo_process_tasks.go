package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	NextMarker common.String
	Tasks      ListPhotoProcessTasksTasksItemList
}

type ListPhotoProcessTasksTasksItem struct {
	TaskId          common.String
	Status          common.String
	SrcUri          common.String
	TgtUri          common.String
	Style           common.String
	NotifyTopicName common.String
	NotifyEndpoint  common.String
	ExternalID      common.String
	CreateTime      common.String
	FinishTime      common.String
	Percent         common.Integer
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
