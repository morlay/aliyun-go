package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteVirtualMFADeviceRequest struct {
	requests.RpcRequest
	SerialNumber string `position:"Query" name:"SerialNumber"`
}

func (req *DeleteVirtualMFADeviceRequest) Invoke(client *sdk.Client) (resp *DeleteVirtualMFADeviceResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "DeleteVirtualMFADevice", "", "")
	resp = &DeleteVirtualMFADeviceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteVirtualMFADeviceResponse struct {
	responses.BaseResponse
	RequestId string
}
