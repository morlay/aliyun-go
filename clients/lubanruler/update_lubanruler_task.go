package lubanruler

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateLubanrulerTaskRequest struct {
	requests.RpcRequest
	AoneInfo string `position:"Body" name:"AoneInfo"`
	Token    string `position:"Body" name:"Token"`
}

func (req *UpdateLubanrulerTaskRequest) Invoke(client *sdk.Client) (resp *UpdateLubanrulerTaskResponse, err error) {
	req.InitWithApiInfo("Lubanruler", "2017-12-28", "UpdateLubanrulerTask", "", "")
	resp = &UpdateLubanrulerTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateLubanrulerTaskResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
}
