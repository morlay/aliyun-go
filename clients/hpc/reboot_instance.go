package hpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RebootInstanceRequest struct {
	requests.RpcRequest
	TOKEN      string `position:"Query" name:"TOKEN"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *RebootInstanceRequest) Invoke(client *sdk.Client) (resp *RebootInstanceResponse, err error) {
	req.InitWithApiInfo("HPC", "2016-12-13", "RebootInstance", "hpc", "")
	resp = &RebootInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RebootInstanceResponse struct {
	responses.BaseResponse
	RequestId string
}
