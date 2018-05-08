package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	JobId      common.String
	SetId      common.String
	SrcUri     common.String
	Status     common.String
	Percent    common.Integer
	CreateTime common.String
	FinishTime common.String
}
