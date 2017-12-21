package hpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyJumpserverPasswordRequest struct {
	TOKEN       string `position:"Query" name:"TOKEN"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	NewPassword string `position:"Query" name:"NewPassword"`
}

func (r ModifyJumpserverPasswordRequest) Invoke(client *sdk.Client) (response *ModifyJumpserverPasswordResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyJumpserverPasswordRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("HPC", "2016-12-13", "ModifyJumpserverPassword", "hpc", "")

	resp := struct {
		*responses.BaseResponse
		ModifyJumpserverPasswordResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyJumpserverPasswordResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyJumpserverPasswordResponse struct {
	RequestId string
}
