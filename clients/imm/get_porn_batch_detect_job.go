package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId       common.String
	JobId           common.String
	SrcUri          common.String
	TgtUri          common.String
	NotifyTopicName common.String
	NotifyEndpoint  common.String
	ExternalID      common.String
	Status          common.String
	CreateTime      common.String
	FinishTime      common.String
	Percent         common.Integer
}
