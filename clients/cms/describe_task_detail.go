package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTaskDetailRequest struct {
	requests.RpcRequest
	TaskId string `position:"Query" name:"TaskId"`
}

func (req *DescribeTaskDetailRequest) Invoke(client *sdk.Client) (resp *DescribeTaskDetailResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "DescribeTaskDetail", "cms", "")
	resp = &DescribeTaskDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeTaskDetailResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	Success   string
	RequestId string
	Data      string
}
