package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeviceUpdateRequest struct {
	DevicePosition string `position:"Query" name:"DevicePosition"`
	DeviceName     string `position:"Query" name:"DeviceName"`
	Did            int64  `position:"Query" name:"Did"`
}

func (r DeviceUpdateRequest) Invoke(client *sdk.Client) (response *DeviceUpdateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeviceUpdateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DeviceUpdate", "", "")

	resp := struct {
		*responses.BaseResponse
		DeviceUpdateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeviceUpdateResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeviceUpdateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
