package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StopTasksRequest struct {
	requests.RpcRequest
	TaskIds string `position:"Query" name:"TaskIds"`
}

func (req *StopTasksRequest) Invoke(client *sdk.Client) (resp *StopTasksResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "StopTasks", "cms", "")
	resp = &StopTasksResponse{}
	err = client.DoAction(req, resp)
	return
}

type StopTasksResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	Success   string
	RequestId string
}
