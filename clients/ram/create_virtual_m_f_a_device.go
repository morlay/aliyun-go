package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId        common.String
	VirtualMFADevice CreateVirtualMFADeviceVirtualMFADevice
}

type CreateVirtualMFADeviceVirtualMFADevice struct {
	SerialNumber     common.String
	Base32StringSeed common.String
	QRCodePNG        common.String
}
