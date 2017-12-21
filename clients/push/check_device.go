package push

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckDeviceRequest struct {
	AppKey   int64  `position:"Query" name:"AppKey"`
	DeviceId string `position:"Query" name:"DeviceId"`
}

func (r CheckDeviceRequest) Invoke(client *sdk.Client) (response *CheckDeviceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CheckDeviceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Push", "2016-08-01", "CheckDevice", "", "")

	resp := struct {
		*responses.BaseResponse
		CheckDeviceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CheckDeviceResponse

	err = client.DoAction(&req, &resp)
	return
}

type CheckDeviceResponse struct {
	RequestId string
	Available bool
}
