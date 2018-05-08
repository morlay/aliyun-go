package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code       common.String
	Message    common.String
	Success    common.String
	RequestId  common.String
	Data       common.String
	PageNumber common.Integer
	PageSize   common.Integer
	TotalCount common.Integer
}
