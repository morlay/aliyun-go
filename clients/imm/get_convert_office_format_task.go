package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetConvertOfficeFormatTaskRequest struct {
	requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	TaskId  string `position:"Query" name:"TaskId"`
}

func (req *GetConvertOfficeFormatTaskRequest) Invoke(client *sdk.Client) (resp *GetConvertOfficeFormatTaskResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "GetConvertOfficeFormatTask", "imm", "")
	resp = &GetConvertOfficeFormatTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetConvertOfficeFormatTaskResponse struct {
	responses.BaseResponse
	RequestId       string
	TgtType         string
	Status          string
	Percent         int
	PageCount       int
	TaskId          string
	TgtUri          string
	ImageSpec       string
	NotifyTopicName string
	NotifyEndpoint  string
	ExternalID      string
	CreateTime      string
	FinishTime      string
	SrcUri          string
}
