package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId       common.String
	TaskId          common.String
	Status          common.String
	SrcUri          common.String
	TgtUri          common.String
	Style           common.String
	NotifyTopicName common.String
	NotifyEndpoint  common.String
	ExternalID      common.String
	CreateTime      common.String
	FinishTime      common.String
	Percent         common.Integer
}
