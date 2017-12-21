package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeviceUpdateRequest struct {
	requests.RpcRequest
	DevicePosition string `position:"Query" name:"DevicePosition"`
	DeviceName     string `position:"Query" name:"DeviceName"`
	Did            int64  `position:"Query" name:"Did"`
}

func (req *DeviceUpdateRequest) Invoke(client *sdk.Client) (resp *DeviceUpdateResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DeviceUpdate", "", "")
	resp = &DeviceUpdateResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeviceUpdateResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
