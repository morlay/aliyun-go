package hpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StopJumpserverRequest struct {
	TOKEN      string `position:"Query" name:"TOKEN"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Force      int    `position:"Query" name:"Force"`
}

func (r StopJumpserverRequest) Invoke(client *sdk.Client) (response *StopJumpserverResponse, err error) {
	req := struct {
		*requests.RpcRequest
		StopJumpserverRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("HPC", "2016-12-13", "StopJumpserver", "hpc", "")

	resp := struct {
		*responses.BaseResponse
		StopJumpserverResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.StopJumpserverResponse

	err = client.DoAction(&req, &resp)
	return
}

type StopJumpserverResponse struct {
	RequestId string
}
