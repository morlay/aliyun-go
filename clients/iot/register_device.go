package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId    common.String
	Success      bool
	ErrorMessage common.String
	Data         RegisterDeviceData
}

type RegisterDeviceData struct {
	IotId        common.String
	ProductKey   common.String
	DeviceName   common.String
	DeviceSecret common.String
}
