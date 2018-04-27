package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListClusterOperationHostTaskForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	OperationId     string `position:"Query" name:"OperationId"`
	HostId          string `position:"Query" name:"HostId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	UserId          string `position:"Query" name:"UserId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
	Status          string `position:"Query" name:"Status"`
}

func (req *ListClusterOperationHostTaskForAdminRequest) Invoke(client *sdk.Client) (resp *ListClusterOperationHostTaskForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterOperationHostTaskForAdmin", "", "")
	resp = &ListClusterOperationHostTaskForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterOperationHostTaskForAdminResponse struct {
	responses.BaseResponse
	RequestId                    string
	TotalCount                   int
	PageNumber                   int
	PageSize                     int
	ClusterOperationHostTaskList ListClusterOperationHostTaskForAdminClusterOperationHostTaskList
}

type ListClusterOperationHostTaskForAdminClusterOperationHostTask struct {
	TaskId     string
	TaskName   string
	Status     string
	Percentage string
}

type ListClusterOperationHostTaskForAdminClusterOperationHostTaskList []ListClusterOperationHostTaskForAdminClusterOperationHostTask

func (list *ListClusterOperationHostTaskForAdminClusterOperationHostTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterOperationHostTaskForAdminClusterOperationHostTask)
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
