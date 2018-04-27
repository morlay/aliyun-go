package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RegisterDeviceRequest struct {
	requests.RpcRequest
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (req *RegisterDeviceRequest) Invoke(client *sdk.Client) (resp *RegisterDeviceResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "RegisterDevice", "", "")
	resp = &RegisterDeviceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RegisterDeviceResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
	Data         RegisterDeviceData
}

type RegisterDeviceData struct {
	IotId        string
	ProductKey   string
	DeviceName   string
	DeviceSecret string
}
