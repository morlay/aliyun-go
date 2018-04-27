package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListOpsOperationTaskRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
	OperationId     int64 `position:"Query" name:"OperationId"`
	PageNumber      int64 `position:"Query" name:"PageNumber"`
}

func (req *ListOpsOperationTaskRequest) Invoke(client *sdk.Client) (resp *ListOpsOperationTaskResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListOpsOperationTask", "", "")
	resp = &ListOpsOperationTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListOpsOperationTaskResponse struct {
	responses.BaseResponse
	RequestId  string
	Total      int
	PageNumber int
	PageSize   int
	TaskList   ListOpsOperationTaskTaskListItemList
}

type ListOpsOperationTaskTaskListItem struct {
	Id                int64
	OpsOperationId    int64
	TaskId            int64
	Status            string
	RegionId          string
	UserId            string
	ClusterId         int64
	ExternalClusterId string
	HostId            int64
	StartTime         string
	EndTime           string
	CommandName       string
	HostName          string
}

type ListOpsOperationTaskTaskListItemList []ListOpsOperationTaskTaskListItem

func (list *ListOpsOperationTaskTaskListItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListOpsOperationTaskTaskListItem)
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
