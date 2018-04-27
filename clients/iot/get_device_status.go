package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetDeviceStatusRequest struct {
	requests.RpcRequest
	IotId      string `position:"Query" name:"IotId"`
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (req *GetDeviceStatusRequest) Invoke(client *sdk.Client) (resp *GetDeviceStatusResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "GetDeviceStatus", "", "")
	resp = &GetDeviceStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetDeviceStatusResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
	Data         GetDeviceStatusData
}

type GetDeviceStatusData struct {
	Status string
}
