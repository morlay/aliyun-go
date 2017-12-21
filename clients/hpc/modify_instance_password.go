package hpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyInstancePasswordRequest struct {
	TOKEN       string `position:"Query" name:"TOKEN"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	NewPassword string `position:"Query" name:"NewPassword"`
}

func (r ModifyInstancePasswordRequest) Invoke(client *sdk.Client) (response *ModifyInstancePasswordResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyInstancePasswordRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("HPC", "2016-12-13", "ModifyInstancePassword", "hpc", "")

	resp := struct {
		*responses.BaseResponse
		ModifyInstancePasswordResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyInstancePasswordResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyInstancePasswordResponse struct {
	RequestId string
}
