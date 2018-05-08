package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId        common.String
	Total            common.Integer
	PageNumber       common.Integer
	PageSize         common.Integer
	OpsOperationList ListOpsOperationOpsOperationListItemList
}

type ListOpsOperationOpsOperationListItem struct {
	Id              common.Long
	StartTime       common.String
	EndTime         common.String
	OpsCommandId    common.Long
	OpsCommandName  common.String
	Status          common.String
	TotalTaskNum    common.Long
	FailedTaskNum   common.Long
	FinishedTaskNum common.Long
	ClusterId       common.String
	RegionId        common.String
	Params          common.String
	Remark          common.String
	RunningTime     common.Long
	Category        common.String
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
