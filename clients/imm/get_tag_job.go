package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetTagJobRequest struct {
	requests.RpcRequest
	JobId   string `position:"Query" name:"JobId"`
	Project string `position:"Query" name:"Project"`
}

func (req *GetTagJobRequest) Invoke(client *sdk.Client) (resp *GetTagJobResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "GetTagJob", "imm", "")
	resp = &GetTagJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetTagJobResponse struct {
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
