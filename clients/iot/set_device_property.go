package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetDevicePropertyRequest struct {
	requests.RpcRequest
	IotId      string `position:"Query" name:"IotId"`
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
	Items      string `position:"Query" name:"Items"`
}

func (req *SetDevicePropertyRequest) Invoke(client *sdk.Client) (resp *SetDevicePropertyResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "SetDeviceProperty", "", "")
	resp = &SetDevicePropertyResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetDevicePropertyResponse struct {
	responses.BaseResponse
	RequestId    common.String
	Success      bool
	ErrorMessage common.String
}
