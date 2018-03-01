package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateTagJobRequest struct {
	requests.RpcRequest
	NotifyTopicName string `position:"Query" name:"NotifyTopicName"`
	NotifyEndpoint  string `position:"Query" name:"NotifyEndpoint"`
	Project         string `position:"Query" name:"Project"`
	ExternalID      string `position:"Query" name:"ExternalID"`
	SrcUri          string `position:"Query" name:"SrcUri"`
}

func (req *CreateTagJobRequest) Invoke(client *sdk.Client) (resp *CreateTagJobResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "CreateTagJob", "imm", "")
	resp = &CreateTagJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateTagJobResponse struct {
	responses.BaseResponse
	RequestId  string
	JobId      string
	SetId      string
	SrcUri     string
	Status     string
	Percent    int
	CreateTime string
	FinishTime string
}
