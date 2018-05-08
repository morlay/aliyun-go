package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTasksRequest struct {
	requests.RpcRequest
	TaskType string `position:"Query" name:"TaskType"`
	PageSize int    `position:"Query" name:"PageSize"`
	Page     int    `position:"Query" name:"Page"`
	Keyword  string `position:"Query" name:"Keyword"`
	TaskId   string `position:"Query" name:"TaskId"`
}

func (req *DescribeTasksRequest) Invoke(client *sdk.Client) (resp *DescribeTasksResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "DescribeTasks", "cms", "")
	resp = &DescribeTasksResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeTasksResponse struct {
	responses.BaseResponse
	Code       string
	Message    string
	Success    string
	RequestId  string
	Data       string
	PageNumber int
	PageSize   int
	TotalCount int
}
