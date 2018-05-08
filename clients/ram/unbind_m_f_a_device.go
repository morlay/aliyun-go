package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UnbindMFADeviceRequest struct {
	requests.RpcRequest
	UserName string `position:"Query" name:"UserName"`
}

func (req *UnbindMFADeviceRequest) Invoke(client *sdk.Client) (resp *UnbindMFADeviceResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "UnbindMFADevice", "", "")
	resp = &UnbindMFADeviceResponse{}
	err = client.DoAction(req, resp)
	return
}

type UnbindMFADeviceResponse struct {
	responses.BaseResponse
	RequestId common.String
	MFADevice UnbindMFADeviceMFADevice
}

type UnbindMFADeviceMFADevice struct {
	SerialNumber common.String
}
