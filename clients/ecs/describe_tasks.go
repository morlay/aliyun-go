package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTasksRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TaskIds              string `position:"Query" name:"TaskIds"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	TaskStatus           string `position:"Query" name:"TaskStatus"`
	PageSize             int    `position:"Query" name:"PageSize"`
	TaskAction           string `position:"Query" name:"TaskAction"`
}

func (req *DescribeTasksRequest) Invoke(client *sdk.Client) (resp *DescribeTasksResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeTasks", "ecs", "")
	resp = &DescribeTasksResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeTasksResponse struct {
	responses.BaseResponse
	RequestId  string
	RegionId   string
	TotalCount int
	PageNumber int
	PageSize   int
	TaskSet    DescribeTasksTaskList
}

type DescribeTasksTask struct {
	TaskId        string
	TaskAction    string
	TaskStatus    string
	SupportCancel string
	CreationTime  string
	FinishedTime  string
}

type DescribeTasksTaskList []DescribeTasksTask

func (list *DescribeTasksTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTasksTask)
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
