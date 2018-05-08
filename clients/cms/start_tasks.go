package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StartTasksRequest struct {
	requests.RpcRequest
	TaskIds string `position:"Query" name:"TaskIds"`
}

func (req *StartTasksRequest) Invoke(client *sdk.Client) (resp *StartTasksResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "StartTasks", "cms", "")
	resp = &StartTasksResponse{}
	err = client.DoAction(req, resp)
	return
}

type StartTasksResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	Success   string
	RequestId string
}
