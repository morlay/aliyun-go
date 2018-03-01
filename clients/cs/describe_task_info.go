package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTaskInfoRequest struct {
	requests.RoaRequest
	TaskId string `position:"Path" name:"TaskId"`
}

func (req *DescribeTaskInfoRequest) Invoke(client *sdk.Client) (resp *DescribeTaskInfoResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeTaskInfo", "/tasks/[TaskId]", "", "")
	req.Method = "GET"

	resp = &DescribeTaskInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeTaskInfoResponse struct {
	responses.BaseResponse
}
