package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	NextMarker common.String
	Tasks      ListConvertOfficeFormatTasksTasksItemList
}

type ListConvertOfficeFormatTasksTasksItem struct {
	TaskId          common.String
	Status          common.String
	Percent         common.Integer
	PageCount       common.Integer
	SrcUri          common.String
	TgtType         common.String
	TgtUri          common.String
	ImageSpec       common.String
	NotifyTopicName common.String
	NotifyEndpoint  common.String
	ExternalID      common.String
	CreateTime      common.String
	FinishTime      common.String
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
