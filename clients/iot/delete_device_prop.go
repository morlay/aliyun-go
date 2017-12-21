package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDevicePropRequest struct {
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
	PropKey    string `position:"Query" name:"PropKey"`
}

func (r DeleteDevicePropRequest) Invoke(client *sdk.Client) (response *DeleteDevicePropResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteDevicePropRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Iot", "2017-04-20", "DeleteDeviceProp", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteDevicePropResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteDevicePropResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteDevicePropResponse struct {
	RequestId    string
	Success      bool
	ErrorMessage string
}
