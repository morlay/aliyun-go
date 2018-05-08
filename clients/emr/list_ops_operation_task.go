package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	Total      common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	TaskList   ListOpsOperationTaskTaskListItemList
}

type ListOpsOperationTaskTaskListItem struct {
	Id                common.Long
	OpsOperationId    common.Long
	TaskId            common.Long
	Status            common.String
	RegionId          common.String
	UserId            common.String
	ClusterId         common.Long
	ExternalClusterId common.String
	HostId            common.Long
	StartTime         common.String
	EndTime           common.String
	CommandName       common.String
	HostName          common.String
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
