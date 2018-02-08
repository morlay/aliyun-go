package lubanruler

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateLubanrulerTaskRequest struct {
	requests.RpcRequest
	TaskDO string `position:"Body" name:"TaskDO"`
	Token  string `position:"Body" name:"Token"`
}

func (req *CreateLubanrulerTaskRequest) Invoke(client *sdk.Client) (resp *CreateLubanrulerTaskResponse, err error) {
	req.InitWithApiInfo("Lubanruler", "2017-12-28", "CreateLubanrulerTask", "", "")
	resp = &CreateLubanrulerTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateLubanrulerTaskResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
}
