package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListOfficeConversionTaskRequest struct {
	requests.RpcRequest
	MaxKeys int    `position:"Query" name:"MaxKeys"`
	Marker  string `position:"Query" name:"Marker"`
	Project string `position:"Query" name:"Project"`
}

func (req *ListOfficeConversionTaskRequest) Invoke(client *sdk.Client) (resp *ListOfficeConversionTaskResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "ListOfficeConversionTask", "imm", "")
	resp = &ListOfficeConversionTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListOfficeConversionTaskResponse struct {
	responses.BaseResponse
	RequestId  common.String
	NextMarker common.String
	Tasks      ListOfficeConversionTaskTasksItemList
}

type ListOfficeConversionTaskTasksItem struct {
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

type ListOfficeConversionTaskTasksItemList []ListOfficeConversionTaskTasksItem

func (list *ListOfficeConversionTaskTasksItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListOfficeConversionTaskTasksItem)
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
