package hpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyJumpserverPasswordRequest struct {
	requests.RpcRequest
	TOKEN       string `position:"Query" name:"TOKEN"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	NewPassword string `position:"Query" name:"NewPassword"`
}

func (req *ModifyJumpserverPasswordRequest) Invoke(client *sdk.Client) (resp *ModifyJumpserverPasswordResponse, err error) {
	req.InitWithApiInfo("HPC", "2016-12-13", "ModifyJumpserverPassword", "hpc", "")
	resp = &ModifyJumpserverPasswordResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyJumpserverPasswordResponse struct {
	responses.BaseResponse
	RequestId string
}
