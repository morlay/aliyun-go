package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListOpsOperationRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
}

func (req *ListOpsOperationRequest) Invoke(client *sdk.Client) (resp *ListOpsOperationResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListOpsOperation", "", "")
	resp = &ListOpsOperationResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListOpsOperationResponse struct {
	responses.BaseResponse
	RequestId        string
	Total            int
	PageNumber       int
	PageSize         int
	OpsOperationList ListOpsOperationOpsOperationListItemList
}

type ListOpsOperationOpsOperationListItem struct {
	Id              int64
	StartTime       string
	EndTime         string
	OpsCommandId    int64
	OpsCommandName  string
	Status          string
	TotalTaskNum    int64
	FailedTaskNum   int64
	FinishedTaskNum int64
	ClusterId       string
	RegionId        string
	Params          string
	Remark          string
	RunningTime     int64
	Category        string
}

type ListOpsOperationOpsOperationListItemList []ListOpsOperationOpsOperationListItem

func (list *ListOpsOperationOpsOperationListItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListOpsOperationOpsOperationListItem)
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
