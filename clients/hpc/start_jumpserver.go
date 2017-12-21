package hpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StartJumpserverRequest struct {
	TOKEN      string `position:"Query" name:"TOKEN"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Force      int    `position:"Query" name:"Force"`
}

func (r StartJumpserverRequest) Invoke(client *sdk.Client) (response *StartJumpserverResponse, err error) {
	req := struct {
		*requests.RpcRequest
		StartJumpserverRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("HPC", "2016-12-13", "StartJumpserver", "hpc", "")

	resp := struct {
		*responses.BaseResponse
		StartJumpserverResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.StartJumpserverResponse

	err = client.DoAction(&req, &resp)
	return
}

type StartJumpserverResponse struct {
	RequestId string
}
