package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OpenCCProtectRequest struct {
	requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

func (req *OpenCCProtectRequest) Invoke(client *sdk.Client) (resp *OpenCCProtectResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "OpenCCProtect", "", "")
	resp = &OpenCCProtectResponse{}
	err = client.DoAction(req, resp)
	return
}

type OpenCCProtectResponse struct {
	responses.BaseResponse
	RequestId string
}
