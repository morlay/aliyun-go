package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteTasksRequest struct {
	requests.RpcRequest
	TaskIds string `position:"Query" name:"TaskIds"`
}

func (req *DeleteTasksRequest) Invoke(client *sdk.Client) (resp *DeleteTasksResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "DeleteTasks", "cms", "")
	resp = &DeleteTasksResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteTasksResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	Success   common.String
	RequestId common.String
}
