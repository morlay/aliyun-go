package saf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ExecuteRequestRequest struct {
	requests.RpcRequest
	ServiceParameters string `position:"Query" name:"ServiceParameters"`
	Service           string `position:"Query" name:"Service"`
}

func (req *ExecuteRequestRequest) Invoke(client *sdk.Client) (resp *ExecuteRequestResponse, err error) {
	req.InitWithApiInfo("saf", "2017-03-31", "ExecuteRequest", "", "")
	resp = &ExecuteRequestResponse{}
	err = client.DoAction(req, resp)
	return
}

type ExecuteRequestResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Data      common.String
	Message   common.String
	RequestId common.String
}
