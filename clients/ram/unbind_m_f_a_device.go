package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UnbindMFADeviceRequest struct {
	UserName string `position:"Query" name:"UserName"`
}

func (r UnbindMFADeviceRequest) Invoke(client *sdk.Client) (response *UnbindMFADeviceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UnbindMFADeviceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "UnbindMFADevice", "", "")

	resp := struct {
		*responses.BaseResponse
		UnbindMFADeviceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UnbindMFADeviceResponse

	err = client.DoAction(&req, &resp)
	return
}

type UnbindMFADeviceResponse struct {
	RequestId string
	MFADevice UnbindMFADeviceMFADevice
}

type UnbindMFADeviceMFADevice struct {
	SerialNumber string
}
