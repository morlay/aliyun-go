package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OpenCCProtectRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

func (r OpenCCProtectRequest) Invoke(client *sdk.Client) (response *OpenCCProtectResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OpenCCProtectRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "OpenCCProtect", "", "")

	resp := struct {
		*responses.BaseResponse
		OpenCCProtectResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OpenCCProtectResponse

	err = client.DoAction(&req, &resp)
	return
}

type OpenCCProtectResponse struct {
	RequestId string
}
