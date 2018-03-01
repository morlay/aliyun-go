package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetFaceJobRequest struct {
	requests.RpcRequest
	JobId   string `position:"Query" name:"JobId"`
	Project string `position:"Query" name:"Project"`
}

func (req *GetFaceJobRequest) Invoke(client *sdk.Client) (resp *GetFaceJobResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "GetFaceJob", "imm", "")
	resp = &GetFaceJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetFaceJobResponse struct {
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
