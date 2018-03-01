package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeletePhotoProcessTaskRequest struct {
	requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	TaskId  string `position:"Query" name:"TaskId"`
}

func (req *DeletePhotoProcessTaskRequest) Invoke(client *sdk.Client) (resp *DeletePhotoProcessTaskResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "DeletePhotoProcessTask", "imm", "")
	resp = &DeletePhotoProcessTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeletePhotoProcessTaskResponse struct {
	responses.BaseResponse
	RequestId string
}
