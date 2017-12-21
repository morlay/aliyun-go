package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryApplyStatusRequest struct {
	requests.RpcRequest
	ApplyId int64 `position:"Query" name:"ApplyId"`
}

func (req *QueryApplyStatusRequest) Invoke(client *sdk.Client) (resp *QueryApplyStatusResponse, err error) {
	req.InitWithApiInfo("Iot", "2017-04-20", "QueryApplyStatus", "", "")
	resp = &QueryApplyStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryApplyStatusResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
	Finish       bool
}
