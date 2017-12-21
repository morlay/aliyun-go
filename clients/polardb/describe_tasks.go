package polardb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTasksRequest struct {
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

func (r DescribeTasksRequest) Invoke(client *sdk.Client) (response *DescribeTasksResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeTasksRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "DescribeTasks", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeTasksResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeTasksResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeTasksResponse struct {
	RequestId        string
	StartTime        string
	EndTime          string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Tasks            DescribeTasksTaskList
}

type DescribeTasksTask struct {
	TaskId             string
	BeginTime          string
	FinishTime         string
	ExpectedFinishTime string
	TaskAction         string
	Progress           int
	DBName             string
	ProgressInfo       string
	TaskErrorCode      string
	TaskErrorMessage   string
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
