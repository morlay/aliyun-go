package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreatePornBatchDetectJobRequest struct {
	requests.RpcRequest
	NotifyTopicName string `position:"Query" name:"NotifyTopicName"`
	NotifyEndpoint  string `position:"Query" name:"NotifyEndpoint"`
	Project         string `position:"Query" name:"Project"`
	ExternalID      string `position:"Query" name:"ExternalID"`
	SrcUri          string `position:"Query" name:"SrcUri"`
	TgtUri          string `position:"Query" name:"TgtUri"`
}

func (req *CreatePornBatchDetectJobRequest) Invoke(client *sdk.Client) (resp *CreatePornBatchDetectJobResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "CreatePornBatchDetectJob", "imm", "")
	resp = &CreatePornBatchDetectJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreatePornBatchDetectJobResponse struct {
	responses.BaseResponse
	RequestId  string
	JobId      string
	TgtLoc     string
	Status     string
	CreateTime string
	Percent    int
}
