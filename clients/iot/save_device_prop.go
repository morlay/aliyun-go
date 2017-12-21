package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveDevicePropRequest struct {
	requests.RpcRequest
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
	Props      string `position:"Query" name:"Props"`
}

func (req *SaveDevicePropRequest) Invoke(client *sdk.Client) (resp *SaveDevicePropResponse, err error) {
	req.InitWithApiInfo("Iot", "2017-04-20", "SaveDeviceProp", "", "")
	resp = &SaveDevicePropResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveDevicePropResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
}
