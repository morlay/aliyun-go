package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDevicePropRequest struct {
	requests.RpcRequest
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
	PropKey    string `position:"Query" name:"PropKey"`
}

func (req *DeleteDevicePropRequest) Invoke(client *sdk.Client) (resp *DeleteDevicePropResponse, err error) {
	req.InitWithApiInfo("Iot", "2017-04-20", "DeleteDeviceProp", "", "")
	resp = &DeleteDevicePropResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteDevicePropResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
}
