package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryDeviceByNameRequest struct {
	requests.RpcRequest
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (req *QueryDeviceByNameRequest) Invoke(client *sdk.Client) (resp *QueryDeviceByNameResponse, err error) {
	req.InitWithApiInfo("Iot", "2017-04-20", "QueryDeviceByName", "", "")
	resp = &QueryDeviceByNameResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryDeviceByNameResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
	DeviceInfo   QueryDeviceByNameDeviceInfo
}

type QueryDeviceByNameDeviceInfo struct {
	DeviceId     string
	DeviceSecret string
	ProductKey   string
	DeviceStatus string
	DeviceName   string
	DeviceType   string
	GmtCreate    string
	GmtModified  string
}
