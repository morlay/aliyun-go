package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetPhotoProcessTaskRequest struct {
	requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	TaskId  string `position:"Query" name:"TaskId"`
}

func (req *GetPhotoProcessTaskRequest) Invoke(client *sdk.Client) (resp *GetPhotoProcessTaskResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "GetPhotoProcessTask", "imm", "")
	resp = &GetPhotoProcessTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetPhotoProcessTaskResponse struct {
	responses.BaseResponse
	RequestId       string
	TaskId          string
	Status          string
	SrcUri          string
	TgtUri          string
	Style           string
	NotifyTopicName string
	NotifyEndpoint  string
	ExternalID      string
	CreateTime      string
	FinishTime      string
	Percent         int
}
