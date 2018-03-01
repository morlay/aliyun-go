package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetPornBatchDetectJobRequest struct {
	requests.RpcRequest
	JobId   string `position:"Query" name:"JobId"`
	Project string `position:"Query" name:"Project"`
}

func (req *GetPornBatchDetectJobRequest) Invoke(client *sdk.Client) (resp *GetPornBatchDetectJobResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "GetPornBatchDetectJob", "imm", "")
	resp = &GetPornBatchDetectJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetPornBatchDetectJobResponse struct {
	responses.BaseResponse
	RequestId       string
	JobId           string
	SrcUri          string
	TgtUri          string
	NotifyTopicName string
	NotifyEndpoint  string
	ExternalID      string
	Status          string
	CreateTime      string
	FinishTime      string
	Percent         int
}
