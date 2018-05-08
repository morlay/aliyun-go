package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CheckDeviceRequest struct {
	requests.RpcRequest
	AppKey   int64  `position:"Query" name:"AppKey"`
	DeviceId string `position:"Query" name:"DeviceId"`
}

func (req *CheckDeviceRequest) Invoke(client *sdk.Client) (resp *CheckDeviceResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "CheckDevice", "", "")
	resp = &CheckDeviceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CheckDeviceResponse struct {
	responses.BaseResponse
	RequestId common.String
	Available bool
}
