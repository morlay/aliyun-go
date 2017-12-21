package hpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RebootJumpserverRequest struct {
	TOKEN      string `position:"Query" name:"TOKEN"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Force      int    `position:"Query" name:"Force"`
}

func (r RebootJumpserverRequest) Invoke(client *sdk.Client) (response *RebootJumpserverResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RebootJumpserverRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("HPC", "2016-12-13", "RebootJumpserver", "hpc", "")

	resp := struct {
		*responses.BaseResponse
		RebootJumpserverResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RebootJumpserverResponse

	err = client.DoAction(&req, &resp)
	return
}

type RebootJumpserverResponse struct {
	RequestId string
}
