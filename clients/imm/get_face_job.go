package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	JobId      common.String
	SetId      common.String
	SrcUri     common.String
	Status     common.String
	Percent    common.Integer
	CreateTime common.String
	FinishTime common.String
}
