package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RegistDeviceRequest struct {
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (r RegistDeviceRequest) Invoke(client *sdk.Client) (response *RegistDeviceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RegistDeviceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Iot", "2017-04-20", "RegistDevice", "", "")

	resp := struct {
		*responses.BaseResponse
		RegistDeviceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RegistDeviceResponse

	err = client.DoAction(&req, &resp)
	return
}

type RegistDeviceResponse struct {
	RequestId    string
	Success      bool
	ErrorMessage string
	DeviceId     string
	DeviceSecret string
	DeviceStatus string
	DeviceName   string
}
