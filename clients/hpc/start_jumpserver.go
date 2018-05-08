package hpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type StartJumpserverRequest struct {
	requests.RpcRequest
	TOKEN      string `position:"Query" name:"TOKEN"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Force      int    `position:"Query" name:"Force"`
}

func (req *StartJumpserverRequest) Invoke(client *sdk.Client) (resp *StartJumpserverResponse, err error) {
	req.InitWithApiInfo("HPC", "2016-12-13", "StartJumpserver", "hpc", "")
	resp = &StartJumpserverResponse{}
	err = client.DoAction(req, resp)
	return
}

type StartJumpserverResponse struct {
	responses.BaseResponse
	RequestId common.String
}
