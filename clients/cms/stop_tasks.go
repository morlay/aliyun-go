package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code      common.String
	Message   common.String
	Success   common.String
	RequestId common.String
}
