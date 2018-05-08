package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId    common.String
	Success      bool
	ErrorMessage common.String
	Data         GetDeviceStatusData
}

type GetDeviceStatusData struct {
	Status common.String
}
