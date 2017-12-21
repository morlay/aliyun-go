package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateDeviceShadowRequest struct {
	ShadowMessage string `position:"Query" name:"ShadowMessage"`
	DeviceName    string `position:"Query" name:"DeviceName"`
	ProductKey    string `position:"Query" name:"ProductKey"`
}

func (r UpdateDeviceShadowRequest) Invoke(client *sdk.Client) (response *UpdateDeviceShadowResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateDeviceShadowRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Iot", "2017-04-20", "UpdateDeviceShadow", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateDeviceShadowResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateDeviceShadowResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateDeviceShadowResponse struct {
	RequestId    string
	Success      bool
	ErrorMessage string
}
