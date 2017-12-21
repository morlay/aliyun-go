package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteVirtualMFADeviceRequest struct {
	SerialNumber string `position:"Query" name:"SerialNumber"`
}

func (r DeleteVirtualMFADeviceRequest) Invoke(client *sdk.Client) (response *DeleteVirtualMFADeviceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteVirtualMFADeviceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "DeleteVirtualMFADevice", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteVirtualMFADeviceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteVirtualMFADeviceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteVirtualMFADeviceResponse struct {
	RequestId string
}
