package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListClusterOperationHostTaskRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	OperationId     string `position:"Query" name:"OperationId"`
	HostId          string `position:"Query" name:"HostId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
	Status          string `position:"Query" name:"Status"`
}

func (req *ListClusterOperationHostTaskRequest) Invoke(client *sdk.Client) (resp *ListClusterOperationHostTaskResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterOperationHostTask", "", "")
	resp = &ListClusterOperationHostTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterOperationHostTaskResponse struct {
	responses.BaseResponse
	RequestId                    string
	TotalCount                   int
	PageNumber                   int
	PageSize                     int
	ClusterOperationHostTaskList ListClusterOperationHostTaskClusterOperationHostTaskList
}

type ListClusterOperationHostTaskClusterOperationHostTask struct {
	TaskId     string
	TaskName   string
	Status     string
	Percentage string
}

type ListClusterOperationHostTaskClusterOperationHostTaskList []ListClusterOperationHostTaskClusterOperationHostTask

func (list *ListClusterOperationHostTaskClusterOperationHostTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterOperationHostTaskClusterOperationHostTask)
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
