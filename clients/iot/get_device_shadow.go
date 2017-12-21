package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetDeviceShadowRequest struct {
	ShadowMessage string `position:"Query" name:"ShadowMessage"`
	DeviceName    string `position:"Query" name:"DeviceName"`
	ProductKey    string `position:"Query" name:"ProductKey"`
}

func (r GetDeviceShadowRequest) Invoke(client *sdk.Client) (response *GetDeviceShadowResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetDeviceShadowRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Iot", "2017-04-20", "GetDeviceShadow", "", "")

	resp := struct {
		*responses.BaseResponse
		GetDeviceShadowResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetDeviceShadowResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetDeviceShadowResponse struct {
	RequestId     string
	Success       bool
	ErrorMessage  string
	ShadowMessage string
}
