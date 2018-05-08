package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeTasksRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	Status               string `position:"Query" name:"Status"`
}

func (req *DescribeTasksRequest) Invoke(client *sdk.Client) (resp *DescribeTasksResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeTasks", "polardb", "")
	resp = &DescribeTasksResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeTasksResponse struct {
	responses.BaseResponse
	RequestId        common.String
	StartTime        common.String
	EndTime          common.String
	TotalRecordCount common.Integer
	PageNumber       common.Integer
	PageRecordCount  common.Integer
	Tasks            DescribeTasksTaskList
}

type DescribeTasksTask struct {
	TaskId             common.String
	BeginTime          common.String
	FinishTime         common.String
	ExpectedFinishTime common.String
	TaskAction         common.String
	Progress           common.Integer
	DBName             common.String
	ProgressInfo       common.String
	TaskErrorCode      common.String
	TaskErrorMessage   common.String
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
