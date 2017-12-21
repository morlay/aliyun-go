package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeviceCreateRequest struct {
	requests.RpcRequest
	DeviceNum      string `position:"Query" name:"DeviceNum"`
	DevicePosition string `position:"Query" name:"DevicePosition"`
	DeviceName     string `position:"Query" name:"DeviceName"`
	DeviceType     int    `position:"Query" name:"DeviceType"`
	Sid            int64  `position:"Query" name:"Sid"`
}

func (req *DeviceCreateRequest) Invoke(client *sdk.Client) (resp *DeviceCreateResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DeviceCreate", "", "")
	resp = &DeviceCreateResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeviceCreateResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
