package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateVirtualMFADeviceRequest struct {
	requests.RpcRequest
	VirtualMFADeviceName string `position:"Query" name:"VirtualMFADeviceName"`
}

func (req *CreateVirtualMFADeviceRequest) Invoke(client *sdk.Client) (resp *CreateVirtualMFADeviceResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "CreateVirtualMFADevice", "", "")
	resp = &CreateVirtualMFADeviceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateVirtualMFADeviceResponse struct {
	responses.BaseResponse
	RequestId        string
	VirtualMFADevice CreateVirtualMFADeviceVirtualMFADevice
}

type CreateVirtualMFADeviceVirtualMFADevice struct {
	SerialNumber     string
	Base32StringSeed string
	QRCodePNG        string
}
