package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryDevicePropRequest struct {
	requests.RpcRequest
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (req *QueryDevicePropRequest) Invoke(client *sdk.Client) (resp *QueryDevicePropResponse, err error) {
	req.InitWithApiInfo("Iot", "2017-04-20", "QueryDeviceProp", "", "")
	resp = &QueryDevicePropResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryDevicePropResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
	Props        string
}
