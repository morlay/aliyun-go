package hpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RebootJumpserverRequest struct {
	requests.RpcRequest
	TOKEN      string `position:"Query" name:"TOKEN"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Force      int    `position:"Query" name:"Force"`
}

func (req *RebootJumpserverRequest) Invoke(client *sdk.Client) (resp *RebootJumpserverResponse, err error) {
	req.InitWithApiInfo("HPC", "2016-12-13", "RebootJumpserver", "hpc", "")
	resp = &RebootJumpserverResponse{}
	err = client.DoAction(req, resp)
	return
}

type RebootJumpserverResponse struct {
	responses.BaseResponse
	RequestId common.String
}
