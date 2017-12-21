package rds

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
	PageNumber           int    `position:"Query" name:"PageNumber"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	TaskAction           string `position:"Query" name:"TaskAction"`
	Status               string `position:"Query" name:"Status"`
}

func (req *DescribeTasksRequest) Invoke(client *sdk.Client) (resp *DescribeTasksResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeTasks", "rds", "")
	resp = &DescribeTasksResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeTasksResponse struct {
	responses.BaseResponse
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeTasksTaskProgressInfoList
}

type DescribeTasksTaskProgressInfo struct {
	DBName             string
	BeginTime          string
	ProgressInfo       string
	FinishTime         string
	TaskAction         string
	TaskId             string
	Progress           string
	ExpectedFinishTime string
	Status             string
	TaskErrorCode      string
	TaskErrorMessage   string
}

type DescribeTasksTaskProgressInfoList []DescribeTasksTaskProgressInfo

func (list *DescribeTasksTaskProgressInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTasksTaskProgressInfo)
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
