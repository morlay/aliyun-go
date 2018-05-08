package hpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyInstancePasswordRequest struct {
	requests.RpcRequest
	TOKEN       string `position:"Query" name:"TOKEN"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	NewPassword string `position:"Query" name:"NewPassword"`
}

func (req *ModifyInstancePasswordRequest) Invoke(client *sdk.Client) (resp *ModifyInstancePasswordResponse, err error) {
	req.InitWithApiInfo("HPC", "2016-12-13", "ModifyInstancePassword", "hpc", "")
	resp = &ModifyInstancePasswordResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyInstancePasswordResponse struct {
	responses.BaseResponse
	RequestId common.String
}
