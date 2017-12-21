package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RegistDeviceRequest struct {
	requests.RpcRequest
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (req *RegistDeviceRequest) Invoke(client *sdk.Client) (resp *RegistDeviceResponse, err error) {
	req.InitWithApiInfo("Iot", "2017-04-20", "RegistDevice", "", "")
	resp = &RegistDeviceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RegistDeviceResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
	DeviceId     string
	DeviceSecret string
	DeviceStatus string
	DeviceName   string
}
