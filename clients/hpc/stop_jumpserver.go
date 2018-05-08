package hpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type StopJumpserverRequest struct {
	requests.RpcRequest
	TOKEN      string `position:"Query" name:"TOKEN"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Force      int    `position:"Query" name:"Force"`
}

func (req *StopJumpserverRequest) Invoke(client *sdk.Client) (resp *StopJumpserverResponse, err error) {
	req.InitWithApiInfo("HPC", "2016-12-13", "StopJumpserver", "hpc", "")
	resp = &StopJumpserverResponse{}
	err = client.DoAction(req, resp)
	return
}

type StopJumpserverResponse struct {
	responses.BaseResponse
	RequestId common.String
}
