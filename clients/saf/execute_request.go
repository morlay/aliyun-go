package saf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ExecuteRequestRequest struct {
	requests.RpcRequest
	ServiceParameters string `position:"Query" name:"ServiceParameters"`
	Service           string `position:"Query" name:"Service"`
}

func (req *ExecuteRequestRequest) Invoke(client *sdk.Client) (resp *ExecuteRequestResponse, err error) {
	req.InitWithApiInfo("saf", "2017-03-31", "ExecuteRequest", "saf", "")
	resp = &ExecuteRequestResponse{}
	err = client.DoAction(req, resp)
	return
}

type ExecuteRequestResponse struct {
	responses.BaseResponse
	Code      int
	Data      string
	Message   string
	RequestId string
}
