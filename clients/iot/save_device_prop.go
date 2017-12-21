package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveDevicePropRequest struct {
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
	Props      string `position:"Query" name:"Props"`
}

func (r SaveDevicePropRequest) Invoke(client *sdk.Client) (response *SaveDevicePropResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SaveDevicePropRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Iot", "2017-04-20", "SaveDeviceProp", "", "")

	resp := struct {
		*responses.BaseResponse
		SaveDevicePropResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SaveDevicePropResponse

	err = client.DoAction(&req, &resp)
	return
}

type SaveDevicePropResponse struct {
	RequestId    string
	Success      bool
	ErrorMessage string
}
