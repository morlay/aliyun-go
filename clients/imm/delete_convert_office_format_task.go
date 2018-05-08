package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteConvertOfficeFormatTaskRequest struct {
	requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	TaskId  string `position:"Query" name:"TaskId"`
}

func (req *DeleteConvertOfficeFormatTaskRequest) Invoke(client *sdk.Client) (resp *DeleteConvertOfficeFormatTaskResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "DeleteConvertOfficeFormatTask", "imm", "")
	resp = &DeleteConvertOfficeFormatTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteConvertOfficeFormatTaskResponse struct {
	responses.BaseResponse
	RequestId common.String
}
