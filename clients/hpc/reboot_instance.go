package hpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RebootInstanceRequest struct {
	TOKEN      string `position:"Query" name:"TOKEN"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (r RebootInstanceRequest) Invoke(client *sdk.Client) (response *RebootInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RebootInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("HPC", "2016-12-13", "RebootInstance", "hpc", "")

	resp := struct {
		*responses.BaseResponse
		RebootInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RebootInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type RebootInstanceResponse struct {
	RequestId string
}
