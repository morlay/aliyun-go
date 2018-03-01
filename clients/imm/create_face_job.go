package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateFaceJobRequest struct {
	requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	SrcUri  string `position:"Query" name:"SrcUri"`
}

func (req *CreateFaceJobRequest) Invoke(client *sdk.Client) (resp *CreateFaceJobResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "CreateFaceJob", "imm", "")
	resp = &CreateFaceJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateFaceJobResponse struct {
	responses.BaseResponse
	RequestId  string
	JobId      string
	SetId      string
	SrcUri     string
	Percent    int
	CreateTime string
	FinishTime string
	Status     string
}
