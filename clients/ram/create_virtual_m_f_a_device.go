package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateVirtualMFADeviceRequest struct {
	VirtualMFADeviceName string `position:"Query" name:"VirtualMFADeviceName"`
}

func (r CreateVirtualMFADeviceRequest) Invoke(client *sdk.Client) (response *CreateVirtualMFADeviceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateVirtualMFADeviceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "CreateVirtualMFADevice", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateVirtualMFADeviceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateVirtualMFADeviceResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateVirtualMFADeviceResponse struct {
	RequestId        string
	VirtualMFADevice CreateVirtualMFADeviceVirtualMFADevice
}

type CreateVirtualMFADeviceVirtualMFADevice struct {
	SerialNumber     string
	Base32StringSeed string
	QRCodePNG        string
}
