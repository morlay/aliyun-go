package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryDevicePropRequest struct {
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (r QueryDevicePropRequest) Invoke(client *sdk.Client) (response *QueryDevicePropResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryDevicePropRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Iot", "2017-04-20", "QueryDeviceProp", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryDevicePropResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryDevicePropResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryDevicePropResponse struct {
	RequestId    string
	Success      bool
	ErrorMessage string
	Props        string
}
